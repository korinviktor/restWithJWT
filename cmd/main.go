package main

import (
	todo "github.com/korinviktor/restWithJWT"
	"github.com/korinviktor/restWithJWT/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
