package spentenergy

import (
	"errors"
	"time"
)

var ErrInputData = errors.New("input data has a value less than or equal to zero")

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrInputData
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, ErrInputData
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * meanSpeed * duration.Minutes()) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps < 0 {
		return 0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	distance := (float64(steps) * stepLen) / mInKm
	return distance
}
