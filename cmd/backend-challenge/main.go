package main

import (
	"log"
	"context"
	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
	"github.com/MaryneZa/backend-challenge/internal/adapter/handler"
	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo"
	"net/http"
)

func main() {

	config, err := config.New()
	if err != nil {
		log.Fatalln("Error loading environment variables", err)
		panic(err)
	}

	client, db, err := mongo.ConnectMongoDB(config.MongoDB)	
	if err != nil {
		log.Fatalln("MongoDB connection error:", err)
		panic(err)
	}
	
	mongo.UserCountLogging(db)

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("MongoDB disconnect error:", err)
		}
	}()

	router := handler.InitRoutes(db, config)

	log.Println("Server is listening on port 8090 ...")
	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatalln("Server error:", err)
	}

}
