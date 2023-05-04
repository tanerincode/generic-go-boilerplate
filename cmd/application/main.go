package main

import (
	"github.com/joho/godotenv"
	"log"
	"tanerincode/generic-go-boilerplate/interval/Server"
	"tanerincode/generic-go-boilerplate/interval/storage/mongo"
)

func main() {
	if dotEnvErr := godotenv.Load(); dotEnvErr != nil {
		log.Fatal("environment file cannot load!")
	}

	storage, err := mongo.New()
	if err != nil {
		log.Printf("could not start app, %v", err)
	}

	defer func() {
		if err = storage.Disconnect(); err != nil {
			log.Fatalln("an error occurred :", err.Error())
		}
	}()

	handlers := Server.NewHandler(storage)
	server := Server.New(":8080", handlers)

	if err = server.Start(); err != nil {
		log.Printf("could not start app, %v", err)
	}
}
