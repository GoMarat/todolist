package main

import (
	"fmt"
	"net/http"
	"os"
	
	"todolist/storage"
)

func main () {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "localhost:7540"
	}

webDir := "./web"
http.Handle("/", http.FileServer(http.Dir(webDir)))

fmt.Println("Server is running on port " + port)

err := http.ListenAndServe(port, nil)
if err != nil {
	panic(err)

	r.Handle("/*", handler.GetFront())
	r.Get("/api/nextdate", handler.NextDate)
	r.Post("/api/task", handler.AddTask(db))
	r.Get("/api/tasks", handler.GetTasks(db))
	r.Get("/api/task", handler.GetTask(db))
	r.Put("/api/task", handler.ChangeTask(db))
	r.Post("/api/task/done", handler.DoneTask(db))
	r.Delete("/api/task", handler.DeleteTask(db))
}
}