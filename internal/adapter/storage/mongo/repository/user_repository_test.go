package repository

import (
	"context"
	"github.com/MaryneZa/backend-challenge/internal/adapter/config"
	mongo_ "github.com/MaryneZa/backend-challenge/internal/adapter/storage/mongo"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"testing"
)

func SetUpTestDB() (*mongo.Client, *mongo.Database, error) {
	config, err := config.New()
	if err != nil {
		log.Fatalln("Error loading environment variables", err)
		panic(err)
	}
	client, db, err := mongo_.ConnectMongoDB(config.TestMongoDB)
	if err != nil {
		log.Fatalln("MongoDB connection error:", err)
		panic(err)
	}
	return client, db, err
}

func TestUserReopo(t *testing.T) {
	client, db, err := SetUpTestDB()
	if err != nil {
		log.Fatalln("MongoDBForTest connection error:", err)
		panic(err)
	}
	ctx := context.TODO()
	defer func() {
		// drop users table !
		if err := db.Collection("users").Drop(ctx); err != nil {
			log.Panicln(err)
		}

		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("MongoDB disconnect error:", err)
		}
	}()

	userRepo := NewUserRepository(db)

	_ = userRepo.SetUpIndexes(ctx)

	user := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	t.Run("create user", func(t *testing.T) {
		err := userRepo.Create(ctx, &user)
		assert.NoError(t, err)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("find user by email (success)", func(t *testing.T) {
		user_, err := userRepo.FindByEmail(ctx, user.Email)
		assert.NoError(t, err)
		assert.Equal(t, user_.Email, user.Email)
		assert.NotNil(t, user_)
		user.ID = user_.ID
		t.Logf("testing find by ID = %v", user.ID.Hex())

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})

	t.Run("find user by email (not success)", func(t *testing.T) {
		user_, err := userRepo.FindByEmail(ctx, "neverbeuser@example.com")
		assert.NoError(t, err)
		assert.Nil(t, user_)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("find user by id (success)", func(t *testing.T) {
		user_, err := userRepo.FindByID(ctx, user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, user_.Email)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})

	t.Run("find user by id (not success)", func(t *testing.T) {
		id := "6831b53acf66afd6d203efbe"
		objectID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			t.Error("cannot parse ObjectID")
		}
		user_, err := userRepo.FindByID(ctx, objectID)
		assert.NoError(t, err)
		assert.Nil(t, user_)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})

	t.Run("update username (success)", func(t *testing.T) {
		name := "testtest"
		err := userRepo.UpdateName(ctx, user.ID, name)
		assert.NoError(t, err)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		user_, err := userRepo.FindByID(ctx, user.ID)
		assert.Equal(t, name, user_.Name)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

	})

	t.Run("update user email (success)", func(t *testing.T) {
		email := "change@example.com"
		err := userRepo.UpdateEmail(ctx, user.ID, email)
		assert.NoError(t, err)

		user_, err := userRepo.FindByID(ctx, user.ID)
		assert.Equal(t, user_.Email, email)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("delete user (success)", func(t *testing.T) {
		email := "change@example.com"
		err := userRepo.Delete(ctx, email)
		assert.NoError(t, err)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		user_, err := userRepo.FindByEmail(ctx, email)
		assert.Nil(t, user_)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})

	t.Run("delete user (not success)", func(t *testing.T) {
		email := "change@example.com"
		err := userRepo.Delete(ctx, email)
		assert.NoError(t, err)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
	})
}
