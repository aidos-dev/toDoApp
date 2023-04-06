package main

import (
	"log"

	todo "github.com/aidos-dev/toDoApp"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Println(err.Error())
		return
	}
}
