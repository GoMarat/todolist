package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)


// CreateDB - создаёт БД в каталоге запуска приложения
func CreateDB(dbFile string) error {

	//Создаём БД
	_, err := os.Create(dbFile)
	if err != nil {
		fmt.Println("БД не создана ", err)
		return err
	}

	//Открываем файл БД
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		fmt.Println("БД не доступна ", err)
		return err
	}
	defer db.Close()

	//Считываем файл SQL для создания БД.
	textSQL, err := os.ReadFile("sqlite/scheduler.sql")
	if err != nil {
		fmt.Println("Не удалось открыть sqlite/scheduler.sql для создания БД ", err)
		return err
	}
	stringSQL := string(textSQL)

	//Создаём таблицу в БД.
	_, err = db.Exec(stringSQL)
	if err != nil {
		fmt.Println("Ошибка при создании БД", err)
		return err
	}
	defer db.Close()

	return nil
}