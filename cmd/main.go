package main

import (
	"github.com/MKKurbandibirov/todo-app"
	"github.com/MKKurbandibirov/todo-app/pkg/handler"
	"github.com/MKKurbandibirov/todo-app/pkg/repository"
	"github.com/MKKurbandibirov/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("couldn't run http server: %s", err.Error())
	}
}
