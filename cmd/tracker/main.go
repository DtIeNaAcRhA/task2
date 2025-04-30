package main

import (
	"fmt"

	"github.com/Yandex-Practicum/tracker/internal/actioninfo"
	"github.com/Yandex-Practicum/tracker/internal/daysteps"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
)

func main() {
	person := personaldata.Personal{
		Name:   "Витя",
		Weight: 84.6,
		Height: 1.87,
	}

	// дневная активность
	input := []string{
		"678,0h50m",
		"792,1h14m",
		"1078,1h30m",
		"7830,2h40m",
		",3456",
		"12:40:00, 3456",
		"something is wrong",
	}

	fmt.Println("Активность в течение дня")

	daySteps := daysteps.DaySteps{
		Personal: person,
	}

	daySteps.Personal.Print()

	actioninfo.Info(input, &daySteps)

	// // тренировки
	actions := []string{
		"3456,Ходьба,3h00m",
		"something is wrong",
		"678,Бег,0h5m",
		"1078,Бег,0h10m",
		",3456 Ходьба",
		"7892,Ходьба,3h10m",
		"15392,Бег,0h45m",
	}

	trains := trainings.Training{
		Personal: person,
	}

	fmt.Println("Журнал тренировок")

	trains.Personal.Print()

	actioninfo.Info(actions, &trains)
}
