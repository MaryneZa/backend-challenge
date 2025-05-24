package main

import (
	"log"
	"fmt"
	"context"
	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
	"github.com/MaryneZa/backend-challenge/internal/adapter/handler"
	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo"
	"net/http"
)

func main() {

	config, err := config.New()
	if err != nil {
		fmt.Errorf("Error loading environment variables %s", err)
		panic(err)
	}

	client, db, err := mongo.ConnectMongoDB(config.MongoDB)
	if err != nil {
		fmt.Println("MongoDB connection error:", err)
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println("MongoDB disconnect error:", err)
		}
	}()

	router := handler.InitRoutes(db, config)

	log.Println("Server is listening on port 8090 ...")
	if err := http.ListenAndServe(":8090", router); err != nil {
		fmt.Println("Server error:", err)
	}
	

}
