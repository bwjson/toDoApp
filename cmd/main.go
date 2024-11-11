package main

import (
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(toDoApp.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running the server %s", err.Error())
	}
}
