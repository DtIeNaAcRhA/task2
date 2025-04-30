package daysteps

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DayStepsTestSuite struct {
	suite.Suite
}

func TestDayStepsSuite(t *testing.T) {
	suite.Run(t, new(DayStepsTestSuite))
}

func (suite *DayStepsTestSuite) TestParsePackage() {
	tests := []struct {
		name         string
		input        string
		wantSteps    int
		wantDuration time.Duration
		wantErr      bool
	}{
		// Корректные значения
		{
			name:         "корректный ввод",
			input:        "678,0h50m",
			wantSteps:    678,
			wantDuration: 50 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "корректный ввод с часами и минутами",
			input:        "1000,1h30m",
			wantSteps:    1000,
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "положительное число с плюсом",
			input:        "+12345,1h30m",
			wantSteps:    12345,
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		// Корректные значения продолжительности
		{
			name:         "продолжительность - только минуты",
			input:        "1000,30m",
			wantSteps:    1000,
			wantDuration: 30 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "продолжительность - только часы",
			input:        "1000,2h",
			wantSteps:    1000,
			wantDuration: 2 * time.Hour,
			wantErr:      false,
		},
		{
			name:         "продолжительность - дробные часы",
			input:        "1000,1.5h",
			wantSteps:    1000,
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "продолжительность - дробные минуты",
			input:        "1000,30.5m",
			wantSteps:    1000,
			wantDuration: 30*time.Minute + 30*time.Second,
			wantErr:      false,
		},
		// Ошибки формата
		{
			name:         "неверный формат - неправильное количество параметров",
			input:        "678",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверный формат - три параметра",
			input:        "678,1h30m,extra",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "пустой ввод",
			input:        "",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		// Ошибки в количестве шагов
		{
			name:         "неверные шаги - не числовое значение",
			input:        "abc,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - ноль",
			input:        "0,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - отрицательное значение",
			input:        "-100,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - только знак минус",
			input:        "-,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - только знак плюс",
			input:        "+,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - пробелы в начале",
			input:        " 12345,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - пробелы в конце",
			input:        "12345 ,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - некорректные символы",
			input:        "123abc,1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		// Ошибки в продолжительности
		{
			name:         "неверный формат продолжительности",
			input:        "678,invalid",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - ноль",
			input:        "678,0h0m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - отрицательное значение",
			input:        "678,-1h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - отрицательные минуты",
			input:        "678,1h-30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - неверная единица измерения",
			input:        "678,1.5d",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - пробел между числом и единицей",
			input:        "678,1 h30m",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - пропущена единица измерения",
			input:        "678,30",
			wantSteps:    0,
			wantDuration: 0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ds := &DaySteps{}
			err := ds.Parse(tt.input)

			if tt.wantErr {
				require.Error(suite.T(), err, "Parse() для строки данных %q ожидалась ошибка, но её нет", tt.input)
				return
			}
			require.NoError(suite.T(), err, "Parse() неожиданная ошибка для строки данных %q: %v", tt.input, err)
			assert.Equal(suite.T(), tt.wantSteps, ds.Steps, "Parse() полученное количество шагов: %v, ожидается %v", ds.Steps, tt.wantSteps)
			assert.Equal(suite.T(), tt.wantDuration, ds.Duration, "Parse() полученная продолжительность прогулки: %v, ожидается %v", ds.Duration, tt.wantDuration)
		})
	}
}

func (suite *DayStepsTestSuite) TestDayActionInfo() {
	tests := []struct {
		name    string
		ds      DaySteps
		want    string
		wantErr bool
	}{
		{
			name: "нормальная нагрузка - один час",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "Количество шагов: 6000.\nДистанция составила 4.72 км.\nВы сожгли 177.19 ккал.\n",
			wantErr: false,
		},
		{
			name: "нормальная нагрузка - полчаса",
			ds: DaySteps{
				Steps:    3000,
				Duration: 30 * time.Minute,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "Количество шагов: 3000.\nДистанция составила 2.36 км.\nВы сожгли 88.59 ккал.\n",
			wantErr: false,
		},
		{
			name: "высокая нагрузка",
			ds: DaySteps{
				Steps:    20000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "Количество шагов: 20000.\nДистанция составила 15.75 км.\nВы сожгли 590.62 ккал.\n",
			wantErr: false,
		},
		{
			name: "низкая нагрузка",
			ds: DaySteps{
				Steps:    1000,
				Duration: 2 * time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "Количество шагов: 1000.\nДистанция составила 0.79 км.\nВы сожгли 29.53 ккал.\n",
			wantErr: false,
		},
		{
			name: "другой вес и рост",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 60.0,
					Height: 1.85,
				},
			},
			want:    "Количество шагов: 6000.\nДистанция составила 5.00 км.\nВы сожгли 149.85 ккал.\n",
			wantErr: false,
		},
		{
			name: "нулевые шаги",
			ds: DaySteps{
				Steps:    0,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "отрицательные шаги",
			ds: DaySteps{
				Steps:    -1000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "нулевая продолжительность",
			ds: DaySteps{
				Steps:    1000,
				Duration: 0,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "отрицательная продолжительность",
			ds: DaySteps{
				Steps:    1000,
				Duration: -time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "нулевой вес",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "отрицательный вес",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: -75.0,
					Height: 1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "нулевой рост",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: 0,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "отрицательный рост",
			ds: DaySteps{
				Steps:    6000,
				Duration: time.Hour,
				Personal: personaldata.Personal{
					Weight: 75.0,
					Height: -1.75,
				},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got, err := tt.ds.ActionInfo()
			if tt.wantErr {
				require.Error(suite.T(), err, "Для тестового случая %q (шаги: %d, продолжительность: %v, вес: %.1f, рост: %.2f) ожидалась ошибка, но её нет",
					tt.name, tt.ds.Steps, tt.ds.Duration, tt.ds.Personal.Weight, tt.ds.Personal.Height)
				require.Empty(suite.T(), got, "Для тестового случая %q (шаги: %d, продолжительность: %v, вес: %.1f, рост: %.2f) ожидалась пустая строка, но получено: %q",
					tt.name, tt.ds.Steps, tt.ds.Duration, tt.ds.Personal.Weight, tt.ds.Personal.Height, got)
				return
			}
			require.NoError(suite.T(), err)
			require.Equal(suite.T(), tt.want, got, "\nActionInfo() получено:\n%v\nожидается:\n%v\n(шаги: %d, продолжительность: %v, вес: %.1f, рост: %.2f)",
				got, tt.want, tt.ds.Steps, tt.ds.Duration, tt.ds.Personal.Weight, tt.ds.Personal.Height)
		})
	}
}
