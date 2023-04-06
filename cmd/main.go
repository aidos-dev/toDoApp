package main

import (
	"log"

	todo "github.com/aidos-dev/toDoApp"
	"github.com/aidos-dev/toDoApp/pkg/handler"
	"github.com/aidos-dev/toDoApp/pkg/repository"
	"github.com/aidos-dev/toDoApp/pkg/service"
)

func main() {
	// handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Println(err.Error())
		return
	}
}
