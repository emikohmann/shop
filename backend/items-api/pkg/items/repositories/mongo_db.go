package repositories

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type itemsMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection string
	logger     *logrus.Logger
}

// NewItemsMongoDB instances a new items' repository against MongoDB
func NewItemsMongoDB(host string, port int, database string, collection string, logger *logrus.Logger) (itemsMongoDB, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		return itemsMongoDB{}, nil
	}

	names, err := client.ListDatabaseNames(ctx, bson.M{})
	logger.Infof("Current MongoDB available databases: %v", names)
	if err != nil {
		return itemsMongoDB{}, nil
	}

	return itemsMongoDB{
		client:     client,
		database:   client.Database(database),
		collection: collection,
	}, nil
}

// GetItem fetches an item from MongoDB
func (repo itemsMongoDB) GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError) {
	result := repo.database.Collection(repo.collection).FindOne(ctx, bson.M{"id": id})
	if result.Err() == mongo.ErrNoDocuments {
		return items.Item{}, apierrors.NewNotFoundError(fmt.Sprintf("not found item %d in MongoDB", id))
	}
	var item items.Item
	if err := result.Decode(&item); err != nil {
		return items.Item{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing item %d from MongoDB: %s", id, err.Error()))
	}
	return item, nil
}

// SaveItem inserts an item into MongoDB
func (repo itemsMongoDB) SaveItem(ctx context.Context, item items.Item) apierrors.APIError {
	_, err := repo.database.Collection(repo.collection).InsertOne(ctx, item)
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error saving item in MongoDB: %s", err.Error()))
	}
	return nil
}

// UpdateItem modifies an item into MongoDB
func (repo itemsMongoDB) UpdateItem(ctx context.Context, item items.Item) apierrors.APIError {
	result, err := repo.database.Collection(repo.collection).UpdateOne(ctx, bson.M{"id": item.ID}, bson.D{{"$set", item}})
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error updating item in MongoDB: %s", err.Error()))
	}
	if result.MatchedCount == 0 {
		return apierrors.NewNotFoundError(fmt.Sprintf("not found item %d in MongoDB", item.ID))
	}
	return nil
}

// DeleteItem removes an item from MongoDB
func (repo itemsMongoDB) DeleteItem(ctx context.Context, id int64) apierrors.APIError {
	result, err := repo.database.Collection(repo.collection).DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error deleting item %d in MongoDB: %s", id, err.Error()))
	}
	if result.DeletedCount == 0 {
		return apierrors.NewNotFoundError(fmt.Sprintf("not found item %d in MongoDB", id))
	}
	return nil
}
