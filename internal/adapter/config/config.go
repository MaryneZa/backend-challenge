package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Container struct {
	MongoDB *DB
	TestMongoDB *DB
	JWTSecret string
}

type DB struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}



func New() (*Container, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	mongoDB := &DB{
		Host: os.Getenv("MONGO_HOST"),
		User: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Name: os.Getenv("MONGO_DATABASE"),
		Port: os.Getenv("MONGO_PORT"),
	}

	testMongoDB := &DB{
		Host: os.Getenv("TEST_MONGO_HOST"),
		User: os.Getenv("TEST_MONGO_USERNAME"),
		Password: os.Getenv("TEST_MONGO_PASSWORD"),
		Name: os.Getenv("TEST_MONGO_DATABASE"),
		Port: os.Getenv("TEST_MONGO_PORT"),
	}

	return &Container{
		MongoDB: mongoDB,
		TestMongoDB: testMongoDB,
		JWTSecret: os.Getenv("JWT_SECRETKEY"),
	}, nil

}