package actioninfo

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDataParser is a mock implementation of the DataParser interface
type MockDataParser struct {
	mock.Mock
}

func (m *MockDataParser) Parse(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDataParser) ActionInfo() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestInfo(t *testing.T) {
	// Set up log output capture
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	// Set up stdout capture
	var stdoutBuf bytes.Buffer
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	defer func() { os.Stdout = old }()

	tests := []struct {
		name           string
		dataset        []string
		setup          func(*MockDataParser)
		expectedOutput string
	}{
		{
			name:    "happy path - single item",
			dataset: []string{"test data"},
			setup: func(m *MockDataParser) {
				m.On("Parse", "test data").Return(nil)
				m.On("ActionInfo").Return("processed test data", nil)
			},
			expectedOutput: "processed test data\n",
		},
		{
			name:           "empty dataset",
			dataset:        []string{},
			setup:          func(m *MockDataParser) {},
			expectedOutput: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logBuf.Reset()
			stdoutBuf.Reset()

			mockParser := new(MockDataParser)
			tt.setup(mockParser)

			Info(tt.dataset, mockParser)

			// Close the write end of the pipe and read the output
			w.Close()
			stdoutBuf.ReadFrom(r)

			mockParser.AssertExpectations(t)
			assert.Contains(t, tt.expectedOutput, stdoutBuf.String(), "вывод в stdout должен соответствовать ожидаемому результату")
		})
	}
}
