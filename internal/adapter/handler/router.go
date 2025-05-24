package handler

import (
	"net/http"
	"context"
	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo/repository"
	"github.com/MaryneZa/backend-challenge/internal/core/service"
	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
	"github.com/MaryneZa/backend-challenge/internal/adapter/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitRoutes(db *mongo.Database, cfg *config.Container) http.Handler {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	_ = userRepo.SetUpIndexes(context.TODO())

	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := NewAuthHandler(authService)

	mux := http.NewServeMux()

	AuthMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret) 

	mux.HandleFunc("/register", userHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)

	mux.Handle("/users", AuthMiddleware(http.HandlerFunc(userHandler.GetAllUsers)))

	mux.Handle("/users/user/email", AuthMiddleware(http.HandlerFunc(userHandler.GetUserByEmail)))
	mux.Handle("/users/user/id", AuthMiddleware(http.HandlerFunc(userHandler.GetUserByID)))

	mux.Handle("/users/user/update-email", AuthMiddleware(http.HandlerFunc(userHandler.UpdateEmail)))
	mux.Handle("/users/user/update-name", AuthMiddleware(http.HandlerFunc(userHandler.UpdateName)))

	mux.Handle("/users/user/delete", AuthMiddleware(http.HandlerFunc(userHandler.DeleteByEmail)))

	return middleware.LoggingMiddleware(mux)
}
