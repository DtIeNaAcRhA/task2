package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, i := range dataset {
		err := dp.Parse(i)
		if err != nil {
			log.Print(err)
			continue
		}
		mess, err := dp.ActionInfo()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(mess)
	}
}
