package personaldata

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name     string
		personal Personal
		want     string
	}{
		{
			name: "стандартные данные",
			personal: Personal{
				Name:   "Иван",
				Weight: 75.0,
				Height: 1.75,
			},
			want: "Имя: Иван\nВес: 75.00 кг.\nРост: 1.75 м.\n\n",
		},
		{
			name: "нулевые значения",
			personal: Personal{
				Name:   "",
				Weight: 0.0,
				Height: 0.0,
			},
			want: "Имя: \nВес: 0.00 кг.\nРост: 0.00 м.\n\n",
		},
		{
			name: "дробные значения",
			personal: Personal{
				Name:   "Петр",
				Weight: 75.5,
				Height: 1.85,
			},
			want: "Имя: Петр\nВес: 75.50 кг.\nРост: 1.85 м.\n\n",
		},
		{
			name: "большие значения",
			personal: Personal{
				Name:   "Алексей",
				Weight: 100.0,
				Height: 2.00,
			},
			want: "Имя: Алексей\nВес: 100.00 кг.\nРост: 2.00 м.\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, err := os.Pipe()
			require.NoError(t, err, "Не удалось создать pipe")
			os.Stdout = w

			done := make(chan bool)
			go func() {
				tt.personal.Print()
				w.Close()
				done <- true
			}()

			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			os.Stdout = old

			<-done

			assert.Equal(t, tt.want, got, "Вывод Print() не соответствует ожидаемому")
		})
	}
}
