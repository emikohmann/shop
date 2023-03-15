package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"
	"items-api/internal/apierrors"
	"items-api/internal/logger"
	"items-api/pkg/items"
	"time"
)

type itemsMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection string
	logger     *logger.Logger
}

type mongoDBItem struct {
	ID           int64     `bson:"id"`
	Name         string    `bson:"name"`
	Description  string    `bson:"description"`
	Thumbnail    string    `bson:"thumbnail"`
	Images       []string  `bson:"images"`
	IsActive     bool      `bson:"is_active"`
	Restrictions []string  `bson:"restrictions"`
	Price        float64   `bson:"price"`
	Stock        int       `bson:"stock"`
	DateCreated  time.Time `bson:"date_created"`
	LastUpdated  time.Time `bson:"last_updated"`
}

func (item mongoDBItem) toItem() items.Item {
	return items.Item{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Thumbnail:    item.Thumbnail,
		Images:       item.Images,
		IsActive:     item.IsActive,
		Restrictions: item.Restrictions,
		Price:        item.Price,
		Stock:        item.Stock,
		DateCreated:  item.DateCreated,
		LastUpdated:  item.LastUpdated,
	}
}

func itemToMongoDBItem(item items.Item) mongoDBItem {
	return mongoDBItem{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Thumbnail:    item.Thumbnail,
		Images:       item.Images,
		IsActive:     item.IsActive,
		Restrictions: item.Restrictions,
		Price:        item.Price,
		Stock:        item.Stock,
		DateCreated:  item.DateCreated,
		LastUpdated:  item.LastUpdated,
	}
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
		return items.Item{}, apierrors.NewNotFoundError(fmt.Sprintf("not found mongoDBItem %d in MongoDB", id))
	}
	var mongoDBItem mongoDBItem
	if err := result.Decode(&mongoDBItem); err != nil {
		return items.Item{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing mongoDBItem %d from MongoDB: %s", id, err.Error()))
	}
	return mongoDBItem.toItem(), nil
}

// ListItems fetches a list of items from MongoDB
func (repo itemsMongoDB) ListItems(ctx context.Context, limit int, offset int) (items.ItemList, apierrors.APIError) {
	options := options.Find()
	options.SetSort(bson.M{"date_created": 1})
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
	itemList := make([]items.Item, 0)
	for result.Next(ctx) {
		var mongoDBItem mongoDBItem
		if err := result.Decode(&mongoDBItem); err != nil {
			return items.ItemList{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing mongoDBItem from MongoDB: %s", err.Error()))
		}
		itemList = append(itemList, mongoDBItem.toItem())
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
		Items: itemList,
	}, nil
}

// SaveItem inserts an item into MongoDB
func (repo itemsMongoDB) SaveItem(ctx context.Context, item items.Item) apierrors.APIError {
	if _, err := repo.database.Collection(repo.collection).InsertOne(ctx, itemToMongoDBItem(item)); err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error saving item in MongoDB: %s", err.Error()))
	}
	return nil
}

// UpdateItem modifies an item into MongoDB
func (repo itemsMongoDB) UpdateItem(ctx context.Context, item items.Item) apierrors.APIError {
	result, err := repo.database.Collection(repo.collection).UpdateOne(ctx, bson.M{"id": item.ID}, bson.D{{"$set", itemToMongoDBItem(item)}})
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
