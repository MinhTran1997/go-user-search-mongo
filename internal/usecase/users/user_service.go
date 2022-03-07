package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type UserService interface {
	All (ctx context.Context) ([]User, error)
	Search (ctx context.Context, filter *UserFilter) (*SearchResult, error)
}

func NewUserService(db *mongo.Database, userCollectionName string) UserService {
	return &userService{
		UserCollection: db.Collection(userCollectionName),
		db: db,
	}
}

type userService struct {
	UserCollection       *mongo.Collection
	db	*mongo.Database
}

func (s *userService)  All(ctx context.Context) ([]User, error){
	result, err := s.UserCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var userRes []User
	for result.Next(ctx) {
		var user User
		err := result.Decode(&user)
		if err != nil {
			return nil, err
		}
		userRes = append(userRes, user)
	}
	return userRes, nil
}

func (s *userService)  Search(ctx context.Context, filter *UserFilter) (*SearchResult, error){
	query := BuildSearchQuery(filter)
	totalSearch := getTotalSearch(s, ctx, query)
	optionsFind := options.Find()
	optionsFind.SetSort(bson.D{{"username", 1}})
	var pageNumber int
	if filter.PageIndex > 0 {
		pageNumber = (filter.PageIndex - 1) * filter.PageSize
		optionsFind.SetSkip(int64(pageNumber))
	}
	if filter.PageSize > 0 {
		optionsFind.SetLimit(int64(filter.PageSize))
	}

	res, err := s.UserCollection.Find(ctx, query, optionsFind)
	if err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			return nil, nil
		}
		return nil, err
	}

	result := SearchResult{}
	for res.Next(ctx) {
		var user User
		if er1 := res.Decode(&user); er1 != nil {
			return nil, er1
		}
		result.List = append(result.List, user)
	}
	result.Total = totalSearch
	result.PageSize = filter.PageSize
	result.PageIndex = filter.PageIndex
	return &result, nil
}

func BuildSearchQuery(filter *UserFilter) bson.D {
	query := bson.D{}
	if filter.Username != "" {
		query = append(query, bson.E{Key: "username", Value: primitive.Regex{Pattern: filter.Username, Options: "i"}})
	}
	if filter.Email != "" {
		query = append(query, bson.E{Key: "email", Value: primitive.Regex{Pattern: filter.Email, Options: "i"}})
	}
	if filter.Phone != "" {
		query = append(query, bson.E{Key: "phone", Value: primitive.Regex{Pattern: filter.Phone, Options: "i"}})
	}
	return query
}

func getTotalSearch(s *userService, ctx context.Context, query bson.D) int {
	res, err := s.UserCollection.Find(ctx, query)
	if err != nil {
		return -1
	}

	result := SearchResult{}
	for res.Next(ctx) {
		var user User
		if er1 := res.Decode(&user); er1 != nil {
			return -1
		}
		result.List = append(result.List, user)
	}

	return len(result.List)
}