package storage

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Подключаемся к БД
func InitDB(dbURL string) (*sql.DB, error) {

	db, err := sql.Open("sqlite", dbURL)

	if err != nil {
		fmt.Println("Отсутствует подключение к базе данных", err)
	}
	
	return db, nil
}