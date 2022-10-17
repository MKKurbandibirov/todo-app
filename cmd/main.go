package main

import (
	"github.com/MKKurbandibirov/todo-app"
	"github.com/MKKurbandibirov/todo-app/pkg/handler"
	"github.com/MKKurbandibirov/todo-app/pkg/repository"
	"github.com/MKKurbandibirov/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("couldn't initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("couldn't run http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
