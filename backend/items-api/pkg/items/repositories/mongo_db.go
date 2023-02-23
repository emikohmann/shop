package repositories

import (
	"context"
	"fmt"
	"github.com/emikohmann/shop/backend/items-api/internal/apierrors"
	"github.com/emikohmann/shop/backend/items-api/pkg/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemsMongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection string
}

// NewItemsMongoDB TODO: Generate connection
func NewItemsMongoDB() itemsMongoDB {
	return itemsMongoDB{}
}

// GetItem fetches an item from MongoDB
func (repo itemsMongoDB) GetItem(ctx context.Context, id int64) (items.Item, apierrors.APIError) {
	objectID, err := primitive.ObjectIDFromHex(fmt.Sprintf("%d", id))
	if err != nil {
		return items.Item{}, apierrors.NewInternalServerError(fmt.Sprintf("error generating object ID for item %d: %s", id, err.Error()))
	}
	result := repo.database.Collection(repo.collection).FindOne(ctx, bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return items.Item{}, apierrors.NewNotFoundError(fmt.Sprintf("not found item %d in MongoDB", id))
	}
	var item items.Item
	if err := result.Decode(&item); err != nil {
		return items.Item{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing item %d from MongoDB: %s", id, err.Error()))
	}
	return item, nil
}
