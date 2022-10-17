package main

import (
	"github.com/MKKurbandibirov/todo-app"
	"github.com/MKKurbandibirov/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("couldn't run http server: %s", err.Error())
	}
}
