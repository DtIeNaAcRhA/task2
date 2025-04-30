package spentenergy

import (
	"testing"
	"time"

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

func (suite *SpentCaloriesTestSuite) TestDistance() {
	tests := []struct {
		name     string
		steps    int
		height   float64
		wantDist float64
	}{
		{
			name:     "нормальное количество шагов",
			steps:    1000,
			height:   1.75,
			wantDist: 0.7875,
		},
		{
			name:     "большое количество шагов",
			steps:    10000,
			height:   1.75,
			wantDist: 7.875,
		},
		{
			name:     "маленькое количество шагов",
			steps:    100,
			height:   1.75,
			wantDist: 0.07875,
		},
		{
			name:     "ноль шагов",
			steps:    0,
			height:   1.75,
			wantDist: 0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := Distance(tt.steps, tt.height)
			assert.Equal(suite.T(), tt.wantDist, got, "Для тестового случая %q (шаги: %d, рост: %.2f) ожидалось значение %.2f, но получено: %.2f",
				tt.name, tt.steps, tt.height, tt.wantDist, got)
		})
	}
}

func (suite *SpentCaloriesTestSuite) TestMeanSpeed() {
	tests := []struct {
		name      string
		steps     int
		height    float64
		duration  time.Duration
		wantSpeed float64
	}{
		{
			name:      "нормальная скорость - один час",
			steps:     6000,
			height:    1.75,
			duration:  1 * time.Hour,
			wantSpeed: 4.725,
		},
		{
			name:      "нормальная скорость - полчаса",
			steps:     3000,
			height:    1.75,
			duration:  30 * time.Minute,
			wantSpeed: 4.725,
		},
		{
			name:      "нормальная скорость - два часа",
			steps:     12000,
			height:    1.75,
			duration:  2 * time.Hour,
			wantSpeed: 4.725,
		},
		{
			name:      "маленькая скорость",
			steps:     1000,
			height:    1.75,
			duration:  2 * time.Hour,
			wantSpeed: 0.39375,
		},
		{
			name:      "большая скорость",
			steps:     20000,
			height:    1.75,
			duration:  1 * time.Hour,
			wantSpeed: 15.75,
		},
		{
			name:      "нулевая продолжительность",
			steps:     1000,
			height:    1.75,
			duration:  0,
			wantSpeed: 0,
		},
		{
			name:      "отрицательная продолжительность",
			steps:     1000,
			height:    1.75,
			duration:  -1 * time.Hour,
			wantSpeed: 0,
		},
		{
			name:      "ноль шагов",
			steps:     0,
			height:    1.75,
			duration:  1 * time.Hour,
			wantSpeed: 0,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			got := MeanSpeed(tt.steps, tt.height, tt.duration)
			assert.Equal(suite.T(), tt.wantSpeed, got, "Для тестового случая %q (шаги: %d, рост: %.2f, продолжительность: %v) ожидалось значение %.2f, но получено: %.2f",
				tt.name, tt.steps, tt.height, tt.duration, tt.wantSpeed, got)
		})
	}
}

func (suite *SpentCaloriesTestSuite) TestRunningSpentCalories() {
	tests := []struct {
		name     string
		steps    int
		weight   float64
		height   float64
		duration time.Duration
		wantCal  float64
		wantErr  bool
	}{
		{
			name:     "нормальная нагрузка - один час",
			steps:    6000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  354.375,
			wantErr:  false,
		},
		{
			name:     "нормальная нагрузка - полчаса",
			steps:    3000,
			weight:   75.0,
			height:   1.75,
			duration: 30 * time.Minute,
			wantCal:  177.1875,
			wantErr:  false,
		},
		{
			name:     "высокая скорость",
			steps:    20000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  1181.25,
			wantErr:  false,
		},
		{
			name:     "низкая скорость",
			steps:    1000,
			weight:   75.0,
			height:   1.75,
			duration: 2 * time.Hour,
			wantCal:  59.0625,
			wantErr:  false,
		},
		{
			name:     "другой вес",
			steps:    6000,
			weight:   60.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  283.5,
			wantErr:  false,
		},
		{
			name:     "нулевая продолжительность",
			steps:    1000,
			weight:   75.0,
			height:   1.75,
			duration: 0,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательная продолжительность",
			steps:    1000,
			weight:   75.0,
			height:   1.75,
			duration: -1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "ноль шагов",
			steps:    0,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательные шаги",
			steps:    -1000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "нулевой вес",
			steps:    1000,
			weight:   0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательный вес",
			steps:    1000,
			weight:   -75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательный рост",
			steps:    1000,
			weight:   75.0,
			height:   -1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			gotCal, gotErr := RunningSpentCalories(tt.steps, tt.weight, tt.height, tt.duration)

			if tt.wantErr {
				require.Error(suite.T(), gotErr, "Для тестового случая %q (шаги: %d, вес: %.1f, рост: %.2f, продолжительность: %v) ожидалась ошибка, но её нет",
					tt.name, tt.steps, tt.weight, tt.height, tt.duration)
				assert.Equal(suite.T(), 0.0, gotCal, "Для тестового случая %q (шаги: %d, вес: %.1f, рост: %.2f, продолжительность: %v) ожидалось значение 0.0, но получено: %.2f",
					tt.name, tt.steps, tt.weight, tt.height, tt.duration, gotCal)
				return
			}

			require.NoError(suite.T(), gotErr)
			assert.InDelta(suite.T(), tt.wantCal, gotCal, 0.1, "Для тестового случая %q (шаги: %d, вес: %.1f, рост: %.2f, продолжительность: %v) ожидалось значение %.2f, но получено: %.2f",
				tt.name, tt.steps, tt.weight, tt.height, tt.duration, tt.wantCal, gotCal)
		})
	}
}

func (suite *SpentCaloriesTestSuite) TestWalkingSpentCalories() {
	tests := []struct {
		name     string
		steps    int
		weight   float64
		height   float64
		duration time.Duration
		wantCal  float64
		wantErr  bool
	}{
		{
			name:     "нормальная нагрузка",
			steps:    6000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  177.19,
			wantErr:  false,
		},
		{
			name:     "меньше шагов",
			steps:    3000,
			weight:   75.0,
			height:   1.75,
			duration: 30 * time.Minute,
			wantCal:  88.594,
			wantErr:  false,
		},
		{
			name:     "больше шагов",
			steps:    20000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  590.62,
			wantErr:  false,
		},
		{
			name:     "другой вес",
			steps:    6000,
			weight:   60.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  141.75,
			wantErr:  false,
		},
		{
			name:     "другой рост",
			steps:    6000,
			weight:   75.0,
			height:   1.85,
			duration: 1 * time.Hour,
			wantCal:  187.313,
			wantErr:  false,
		},
		{
			name:     "нулевые шаги",
			steps:    0,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательные шаги",
			steps:    -1000,
			weight:   75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "нулевой вес",
			steps:    6000,
			weight:   0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательный вес",
			steps:    6000,
			weight:   -75.0,
			height:   1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "нулевой рост",
			steps:    6000,
			weight:   75.0,
			height:   0,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
		{
			name:     "отрицательный рост",
			steps:    6000,
			weight:   75.0,
			height:   -1.75,
			duration: 1 * time.Hour,
			wantCal:  0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			gotCal, gotErr := WalkingSpentCalories(tt.steps, tt.weight, tt.height, tt.duration)

			if tt.wantErr {
				assert.Error(suite.T(), gotErr)
				assert.Equal(suite.T(), 0.0, gotCal)
				return
			}

			assert.NoError(suite.T(), gotErr)
			assert.InDelta(suite.T(), tt.wantCal, gotCal, 0.1)
		})
	}
}
