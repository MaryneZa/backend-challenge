package main

import (
	"fmt"
	"os"

	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo"

)

func main() {

	config, err := config.New()
	if err != nil {
		fmt.Errorf("Error loading environment variables %s", err)
		os.Exit(1)
	}

	db, err := mongo.ConnectMongoDB(config.MongoDB)


}