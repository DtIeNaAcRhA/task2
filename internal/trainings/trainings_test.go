package trainings

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SpentCaloriesTestSuite struct {
	suite.Suite
}

func TestSpentCaloriesSuite(t *testing.T) {
	suite.Run(t, new(SpentCaloriesTestSuite))
}

func (suite *SpentCaloriesTestSuite) TestParseTraining() {
	tests := []struct {
		name         string
		input        string
		wantSteps    int
		wantType     string
		wantDuration time.Duration
		wantErr      bool
	}{
		{
			name:         "корректный ввод с часами и минутами",
			input:        "3456,Ходьба,3h00m",
			wantSteps:    3456,
			wantType:     "Ходьба",
			wantDuration: 3 * time.Hour,
			wantErr:      false,
		},
		{
			name:         "корректный ввод с минутами",
			input:        "678,Бег,5m",
			wantSteps:    678,
			wantType:     "Бег",
			wantDuration: 5 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "положительное число с плюсом",
			input:        "+12345,Ходьба,1h30m",
			wantSteps:    12345,
			wantType:     "Ходьба",
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "продолжительность - только минуты",
			input:        "1000,Бег,30m",
			wantSteps:    1000,
			wantType:     "Бег",
			wantDuration: 30 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "продолжительность - только часы",
			input:        "1000,Ходьба,2h",
			wantSteps:    1000,
			wantType:     "Ходьба",
			wantDuration: 2 * time.Hour,
			wantErr:      false,
		},
		{
			name:         "продолжительность - дробные часы",
			input:        "1000,Бег,1.5h",
			wantSteps:    1000,
			wantType:     "Бег",
			wantDuration: 90 * time.Minute,
			wantErr:      false,
		},
		{
			name:         "продолжительность - дробные минуты",
			input:        "1000,Ходьба,30.5m",
			wantSteps:    1000,
			wantType:     "Ходьба",
			wantDuration: 30*time.Minute + 30*time.Second,
			wantErr:      false,
		},
		{
			name:         "неверный формат - неправильное количество параметров",
			input:        "678,Ходьба",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверный формат - четыре параметра",
			input:        "678,Ходьба,1h30m,extra",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "пустой ввод",
			input:        "",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - не числовое значение",
			input:        "abc,Ходьба,1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - ноль",
			input:        "0,Ходьба,1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - отрицательное значение",
			input:        "-100,Ходьба,1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - только знак минус",
			input:        "-,Ходьба,1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверные шаги - только знак плюс",
			input:        "+,Ходьба,1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверный формат продолжительности",
			input:        "678,Ходьба,invalid",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - ноль",
			input:        "678,Бег,0h0m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - отрицательное значение",
			input:        "678,Ходьба,-1h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - отрицательные минуты",
			input:        "678,Бег,1h-30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - неверная единица измерения",
			input:        "678,Ходьба,1.5d",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - пробел между числом и единицей",
			input:        "678,Бег,1 h30m",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
		{
			name:         "неверная продолжительность - пропущена единица измерения",
			input:        "678,Ходьба,30",
			wantSteps:    0,
			wantType:     "",
			wantDuration: 0,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			training := &Training{}
			err := training.Parse(tt.input)

			if tt.wantErr {
				require.Error(suite.T(), err)
				return
			}

			require.NoError(suite.T(), err)
			assert.Equal(suite.T(), tt.wantSteps, training.Steps)
			assert.Equal(suite.T(), tt.wantType, training.TrainingType)
			assert.Equal(suite.T(), tt.wantDuration, training.Duration)
		})
	}
}

func (suite *SpentCaloriesTestSuite) TestActionInfo() {
	tests := []struct {
		name    string
		input   string
		weight  float64
		height  float64
		want    string
		wantErr bool
	}{
		{
			name:    "ходьба - нормальная нагрузка",
			input:   "6000,Ходьба,1h00m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Ходьба\nДлительность: 1.00 ч.\nДистанция: 4.72 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 177.19\n",
			wantErr: false,
		},
		{
			name:    "бег - нормальная нагрузка",
			input:   "6000,Бег,1h00m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Бег\nДлительность: 1.00 ч.\nДистанция: 4.72 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 354.38\n",
			wantErr: false,
		},
		{
			name:    "ходьба - высокая скорость",
			input:   "20000,Ходьба,1h00m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Ходьба\nДлительность: 1.00 ч.\nДистанция: 15.75 км.\nСкорость: 15.75 км/ч\nСожгли калорий: 590.62\n",
			wantErr: false,
		},
		{
			name:    "бег - высокая скорость",
			input:   "20000,Бег,1h00m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Бег\nДлительность: 1.00 ч.\nДистанция: 15.75 км.\nСкорость: 15.75 км/ч\nСожгли калорий: 1181.25\n",
			wantErr: false,
		},
		{
			name:    "ходьба - другой вес и рост",
			input:   "6000,Ходьба,1h00m",
			weight:  60.0,
			height:  1.85,
			want:    "Тип тренировки: Ходьба\nДлительность: 1.00 ч.\nДистанция: 5.00 км.\nСкорость: 5.00 км/ч\nСожгли калорий: 149.85\n",
			wantErr: false,
		},
		{
			name:    "бег - другой вес",
			input:   "6000,Бег,1h00m",
			weight:  60.0,
			height:  1.75,
			want:    "Тип тренировки: Бег\nДлительность: 1.00 ч.\nДистанция: 4.72 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 283.50\n",
			wantErr: false,
		},
		{
			name:    "ходьба - полчаса",
			input:   "3000,Ходьба,30m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Ходьба\nДлительность: 0.50 ч.\nДистанция: 2.36 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 88.59\n",
			wantErr: false,
		},
		{
			name:    "бег - полчаса",
			input:   "3000,Бег,30m",
			weight:  75.0,
			height:  1.75,
			want:    "Тип тренировки: Бег\nДлительность: 0.50 ч.\nДистанция: 2.36 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 177.19\n",
			wantErr: false,
		},
		{
			name:    "неизвестный тип тренировки",
			input:   "6000,Плавание,1h00m",
			weight:  75.0,
			height:  1.75,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			training := &Training{
				Personal: personaldata.Personal{
					Weight: tt.weight,
					Height: tt.height,
				},
			}

			err := training.Parse(tt.input)
			require.NoError(suite.T(), err, "Для тестового случая %q (ввод: %q, вес: %.1f, рост: %.2f) ожидалась, что парсер сможет распознать ввод, но он не смог",
				tt.name, tt.input, tt.weight, tt.height)

			got, err := training.ActionInfo()
			if tt.wantErr {
				require.Error(suite.T(), err, "Для тестового случая %q (ввод: %q, вес: %.1f, рост: %.2f) ожидалась, что ActionInfo() вернет ошибку, но она не вернулась",
					tt.name, tt.input, tt.weight, tt.height)
				assert.Empty(suite.T(), got, "Для тестового случая %q (ввод: %q, вес: %.1f, рост: %.2f) ожидалась, что ActionInfo() вернет пустую строку, но вернулась: %q",
					tt.name, tt.input, tt.weight, tt.height, got)
				return
			}

			require.NoError(suite.T(), err, "Для тестового случая %q (ввод: %q, вес: %.1f, рост: %.2f) ожидалась, что ActionInfo() вернет ошибку, но она не вернулась",
				tt.name, tt.input, tt.weight, tt.height)
			assert.Equal(suite.T(), tt.want, got, "Для тестового случая %q (ввод: %q, вес: %.1f, рост: %.2f) ожидалось: %q, но вернулось: %q",
				tt.name, tt.input, tt.weight, tt.height, tt.want, got)
		})
	}
}
