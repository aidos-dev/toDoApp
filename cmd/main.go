package main

import (
	"os"

	todo "github.com/aidos-dev/toDoApp"
	"github.com/aidos-dev/toDoApp/pkg/handler"
	"github.com/aidos-dev/toDoApp/pkg/repository"
	"github.com/aidos-dev/toDoApp/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Printf("error occured while running initConfig: %s", err.Error())
		return
	}

	if err := godotenv.Load(); err != nil {
		logrus.Printf("error loading env variables: %s", err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Printf("failed to initialize db: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Printf("error occured while running http server: %s", err.Error())
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
