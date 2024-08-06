package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FindDB - ищет .db в папке приложения
func FindDB(todoDB string) (string, error) {

	var dbURL string

	//В случае отсутствия переменной окружения присваиваем адрес текущего каталога
	if todoDB == "" {
		appPath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}

		dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
		_, err = os.Stat(dbFile)

		fmt.Println("База в ", dbFile)

		if err = CreateDB(dbFile); err != nil {
			fmt.Println(err)
			return dbURL, err
		}

		dbURL = dbFile

	} else {
		dbURL = todoDB
	}

	return dbURL, nil
}