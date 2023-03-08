package repositories

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/internal/logger"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"
)

type itemsMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection string
	logger     *logger.Logger
}

// NewItemsMongoDB instances a new items' repository against MongoDB
func NewItemsMongoDB(ctx context.Context, host string, port int, database string, collection string, logger *logger.Logger) (itemsMongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		logger.Errorf(ctx, "Error connecting to MongoDB: %s", err.Error())
		return itemsMongoDB{}, err
	}

	names, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		logger.Errorf(ctx, "Error listing database names: %s", err.Error())
		return itemsMongoDB{}, err
	}

	if !slices.Contains(names, database) {
		err := fmt.Errorf("%s is not available as MongoDB database, please check the name or create it", database)
		logger.Errorf(ctx, "Error validating MongoDB database: %s", err.Error())
		return itemsMongoDB{}, err
	}

	return itemsMongoDB{
		client:     client,
		database:   client.Database(database),
		collection: collection,
		logger:     logger,
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

// ListItems fetches a list of items from MongoDB
func (repo itemsMongoDB) ListItems(ctx context.Context, limit int, offset int) (items.ItemList, apierrors.APIError) {
	options := options.Find()
	options.SetSort(bson.M{"dateCreated": 1})
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))
	count, err := repo.database.Collection(repo.collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		return items.ItemList{}, apierrors.NewInternalServerError(fmt.Sprintf("error counting items in MongoDB: %s", err.Error()))
	}
	result, err := repo.database.Collection(repo.collection).Find(ctx, bson.M{}, options)
	if err != nil {
		return items.ItemList{}, apierrors.NewInternalServerError(fmt.Sprintf("error listing items in MongoDB: %s", err.Error()))
	}
	list := make([]items.Item, 0)
	for result.Next(ctx) {
		var item items.Item
		if err := result.Decode(&item); err != nil {
			return items.ItemList{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing item from MongoDB: %s", err.Error()))
		}
		list = append(list, item)
	}
	if err := result.Err(); err != nil {
		return items.ItemList{}, apierrors.NewInternalServerError(fmt.Sprintf("error iterating item list from MongoDB: %s", err.Error()))
	}
	return items.ItemList{
		Paging: items.Paging{
			Total:  int(count),
			Limit:  limit,
			Offset: offset,
		},
		Items: list,
	}, nil
}

// SaveItem inserts an item into MongoDB
func (repo itemsMongoDB) SaveItem(ctx context.Context, item items.Item) apierrors.APIError {
	if _, err := repo.database.Collection(repo.collection).InsertOne(ctx, item); err != nil {
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
