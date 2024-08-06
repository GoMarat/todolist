package days

import (
	"time"
	"strconv"
	"strings"
	"fmt"
	"errors"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	var (
		nextDateString string
		nextDateInt    int
		//nextDate	time.Time
	)
	repeatDetails := strings.Split(repeat, " ")
	if len(repeatDetails) == 1 && repeatDetails[1] != "y" {
		return "", fmt.Sprintf("Формат правила повторения не соблюден")
	}	if len(repeatDetails) == 1 && repeatDetails[1] == "y" {
		nextDateInt, err = strconv.Atoi(date) 
		if err != nil {
			panic(err)
		}
		if strings.Contains(date, "0229") {
			nextDateString, err = strconv.Itoa(nextDateInt + 10072)
			if err != nil {
				panic(err)
			} 
			} else {
			nextDateString, err = strconv.Itoa(nextDateInt + 10000)
		if err != nil {
			panic(err)
		}
	}
	return nextDateString, nil
}	if len(repeatDetails) == 2 && repeatDetails[1] == "d" {
	nextDateTime, err := time.Parse("20060102", date)
	daysCount, err := strconv.Atoi(repeatDetails[2])
	if err != nil {
		panic(err)
	}
	nextDateTime = time.AddDate(0, 0, daysCount)
	nextDateString = nextDateTime.Format("20060102")
}


}

