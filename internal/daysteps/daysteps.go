package daysteps

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

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataArray := strings.Split(datastring, ",")
	if len(dataArray) != 2 {
		return inputDataError
	}
	steps, err := strconv.Atoi(dataArray[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(dataArray[1])
	if err != nil {
		return err
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories), nil
}
