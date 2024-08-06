package days

import (
	"time"
	"fmt"
    "strings"
    "strconv"
)




type Task struct {
    date        string  `json:"date"`
    title       string  `json:"title,omitempty"`
    comment     string  `json:"comment"`
    repeat		string  `json:"repeat"`
} 




func Validate(t Task) error {
  //проверка на пустой заголовок
    if t.title == "" {
        return fmt.Errorf("Необходимо заполнить описание задания")
    }
 //проверка на незаполненую дату
    if t.date == "" {
        t.date = time.Now().Format("20060102")
    }

//преобразование даты из строки в формат времени
    validDate, err := time.Parse("20060102", t.date)
    if err != nil {
        return fmt.Errorf("Приложение работает с датами формата YYYYMMDD")
    }

 //проверка на неправильное правило
    if t.repeat != "" && t.repeat[0] != 'd' &&  t.repeat[0] != 'y' {
        return fmt.Errorf("Неизвестное правило")
    }

    if t.repeat[0] == 'd' {
        NumberOfDays, err := strconv.Atoi(strings.Split(t.repeat, " ")[1])
        if err != nil {
            return err
        }
        if NumberOfDays > 400 {
        return fmt.Errorf("В этом правиле число дней должно быть 400 или менее")
        }
    }
now := time.Now()
nowString := now.Format("20060102")
now, err := time.Parse("20060102", nowString)
if err != nil {
    fmt.Println(err)
}
    if now.After(dateIn) {
		if t.Repeate == "" {
			t.Date = fmt.Sprint(now.Format("20060102"))
		} else {
			taskData.Date, err = NextDate(now, taskData.Date, taskData.Repeate)
			if err != nil {
				return err
    }
        }

}




