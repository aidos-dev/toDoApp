package main

import (
	"log"

	todo "github.com/aidos-dev/toDoApp"
	"github.com/aidos-dev/toDoApp/pkg/handler"
	"github.com/aidos-dev/toDoApp/pkg/repository"
	"github.com/aidos-dev/toDoApp/pkg/service"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Printf("error occured while running initConfig: %s", err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Printf("failed to initialize db: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Printf("error occured while running http server: %s", err.Error())
		return
	}
}

func initConfig() error {
	// AddConfigPath receives a derectory name
	viper.AddConfigPath("configs")
	// SetConfig receives a file name (from the directory above)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
