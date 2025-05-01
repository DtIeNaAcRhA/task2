package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var inputDataError = errors.New("Входные параметры некорректны")

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataArray := strings.Split(datastring, ",")
	if len(dataArray) != 3 {
		return inputDataError
	}
	steps, err := strconv.Atoi(dataArray[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("Количество шагов должно быть больше нуля")
	}
	t.Steps = steps
	t.TrainingType = dataArray[1]
	duration, err := time.ParseDuration(dataArray[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("Продолжительность должна быть больше нуля")
	}
	t.Duration = duration
	return nil

}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil

}
