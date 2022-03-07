package app

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	. "search-users/internal/usecase/users"
)

type ApplicationContext struct {
	UserHandler   UserHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
	if err != nil {
		return nil, err
	}

	mongoDb := client.Database(root.Mongo.Database)
	userCollection := "user"
	userService := NewUserService(mongoDb, userCollection)
	userHandler := NewUserHandler(userService)

	return &ApplicationContext{
		UserHandler:   userHandler,
	}, nil
}