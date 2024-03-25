package main

import (
	"HttpServer"
	"HttpServer/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(HttpServer.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server:%s", err.Error())
	}
}
