package main

import (
	"log"

	todo "github.com/aidos-dev/toDoApp"
	"github.com/aidos-dev/toDoApp/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Println(err.Error())
		return
	}
}
