package repository

import (
	"time"
	"context"
	"errors"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"github.com/MaryneZa/backend-challenge/internal/core/port"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) port.UserRepository {
	return &UserRepository{collection: db.Collection("users")}
}

func (ur *UserRepository) SetUpIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true), 
	}
	_, err := ur.collection.Indexes().CreateOne(ctx, indexModel)
	return err
}


func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	user.ID = bson.NewObjectID()
	user.CreatedAt = time.Now()

	data := bson.D{
		{Key: "_id", Value: user.ID},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "name", Value: user.Name},
		{Key: "created_at", Value: user.CreatedAt},
	}

	_, err := ur.collection.InsertOne(ctx, data)
	return err
}

func (ur *UserRepository) FindByID(ctx context.Context, id bson.ObjectID) (*domain.User, error) {
	var user domain.User
	opts := options.FindOne().SetProjection(bson.D{{"password", 0}})
	filter := bson.D{{Key: "_id", Value: id}}
	if err := ur.collection.FindOne(ctx, filter, opts).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	opts := options.FindOne().SetProjection(bson.D{{"password", 0}})

	filter := bson.D{{Key: "email", Value: email}}
	if err := ur.collection.FindOne(ctx, filter, opts).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err 
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUser(ctx context.Context) ([]*domain.User, error) {
	opts := options.Find().SetProjection(bson.D{{"password", 0}})
	cursor, err := ur.collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*domain.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	if users == nil {
		users = []*domain.User{}
	}

	return users, nil
}

func (ur *UserRepository) UpdateEmail(ctx context.Context, id bson.ObjectID, email string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{{Key: "email", Value: email}}}}
	_, err := ur.collection.UpdateOne(ctx, filter, update)
	return err
}

func (ur *UserRepository) UpdateName(ctx context.Context, id bson.ObjectID, name string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{{Key: "name", Value: name}}}}
	_, err := ur.collection.UpdateOne(ctx, filter, update)
	return err
}

func (ur *UserRepository) Delete(ctx context.Context, email string) error {
	filter := bson.D{{Key: "email", Value: email}}
	_, err := ur.collection.DeleteOne(ctx, filter)
	return err
}
