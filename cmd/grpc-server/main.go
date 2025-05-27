package main

import (
	"context"
	"log"
	"net"

	pb "github.com/MaryneZa/backend-challenge/internal/adapter/grpc/stub"

	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo/repository"
	"github.com/MaryneZa/backend-challenge/internal/core/service"

	port_ "github.com/MaryneZa/backend-challenge/internal/core/port"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	config_ "github.com/MaryneZa/backend-challenge/internal/adapter/config"
	"github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo"

)

const port = ":50051"

type server struct{
	userService port_.UserService
	pb.UnimplementedUserServiceServer
}



func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := s.userService.Register(ctx, in.Email, in.Password); err != nil {
		return &pb.CreateUserResponse{}, status.Errorf(codes.NotFound, err.Error())
	}

	return &pb.CreateUserResponse{Message: "successful"}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := s.userService.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	resp := pb.User{
		Id: user.ID.Hex(),
		Email: user.Email,
		Name: user.Name,
		CreatedAt: user.CreatedAt.String(),
	}

	return &pb.GetUserResponse{User: &resp}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed t listen: %v", err)
	}

	config, err := config_.New()
	if err != nil {
		log.Fatalln("Error loading environment variables", err)
		panic(err)
	}

	client, db, err := mongo.ConnectMongoDB(config.MongoDB)	
	if err != nil {
		log.Fatalln("MongoDB connection error:", err)
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("MongoDB disconnect error:", err)
		}
	}()

	grpcServer := grpc.NewServer()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	pb.RegisterUserServiceServer(grpcServer, &server{userService: userService})

	log.Printf("Server is listening on port %v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}