package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

var inputDataError = errors.New("Входные параметры некорректны")

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	person       personaldata.Personal
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
	t.Steps = steps
	t.TrainingType = dataArray[1]
	duration, err := time.ParseDuration(dataArray[2])
	if err != nil {
		return err
	}
	t.Duration = duration
	return nil

}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
