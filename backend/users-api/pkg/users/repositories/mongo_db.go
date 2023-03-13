package repositories

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "golang.org/x/exp/slices"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/users"
)

type usersMongoDB struct {
    client     *mongo.Client
    database   *mongo.Database
    collection string
    logger     *logger.Logger
}

// NewUsersMongoDB instances a new users' repository against MongoDB
func NewUsersMongoDB(ctx context.Context, host string, port int, database string, collection string, logger *logger.Logger) (usersMongoDB, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
    if err != nil {
        logger.Errorf(ctx, "Error connecting to MongoDB: %s", err.Error())
        return usersMongoDB{}, err
    }

    names, err := client.ListDatabaseNames(ctx, bson.M{})
    if err != nil {
        logger.Errorf(ctx, "Error listing database names: %s", err.Error())
        return usersMongoDB{}, err
    }

    if !slices.Contains(names, database) {
        err := fmt.Errorf("%s is not available as MongoDB database, please check the name or create it", database)
        logger.Errorf(ctx, "Error validating MongoDB database: %s", err.Error())
        return usersMongoDB{}, err
    }

    return usersMongoDB{
        client:     client,
        database:   client.Database(database),
        collection: collection,
        logger:     logger,
    }, nil
}

// GetUser fetches an user from MongoDB
func (repo usersMongoDB) GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError) {
    result := repo.database.Collection(repo.collection).FindOne(ctx, bson.M{"id": id})
    if result.Err() == mongo.ErrNoDocuments {
        return users.User{}, apierrors.NewNotFoundError(fmt.Sprintf("not found user %d in MongoDB", id))
    }
    var user users.User
    if err := result.Decode(&user); err != nil {
        return users.User{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing user %d from MongoDB: %s", id, err.Error()))
    }
    return user, nil
}

// ListUsers fetches a list of users from MongoDB
func (repo usersMongoDB) ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError) {
    options := options.Find()
    options.SetSort(bson.M{"dateCreated": 1})
    options.SetLimit(int64(limit))
    options.SetSkip(int64(offset))
    count, err := repo.database.Collection(repo.collection).CountDocuments(ctx, bson.M{})
    if err != nil {
        return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error counting users in MongoDB: %s", err.Error()))
    }
    result, err := repo.database.Collection(repo.collection).Find(ctx, bson.M{}, options)
    if err != nil {
        return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error listing users in MongoDB: %s", err.Error()))
    }
    list := make([]users.User, 0)
    for result.Next(ctx) {
        var user users.User
        if err := result.Decode(&user); err != nil {
            return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error parsing user from MongoDB: %s", err.Error()))
        }
        list = append(list, user)
    }
    if err := result.Err(); err != nil {
        return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error iterating user list from MongoDB: %s", err.Error()))
    }
    return users.UserList{
        Paging: users.Paging{
            Total:  int(count),
            Limit:  limit,
            Offset: offset,
        },
        Users: list,
    }, nil
}

// SaveUser inserts an user into MongoDB
func (repo usersMongoDB) SaveUser(ctx context.Context, user users.User) apierrors.APIError {
    if _, err := repo.database.Collection(repo.collection).InsertOne(ctx, user); err != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error saving user in MongoDB: %s", err.Error()))
    }
    return nil
}

// UpdateUser modifies an user into MongoDB
func (repo usersMongoDB) UpdateUser(ctx context.Context, user users.User) apierrors.APIError {
    result, err := repo.database.Collection(repo.collection).UpdateOne(ctx, bson.M{"id": user.ID}, bson.D{{"$set", user}})
    if err != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error updating user in MongoDB: %s", err.Error()))
    }
    if result.MatchedCount == 0 {
        return apierrors.NewNotFoundError(fmt.Sprintf("not found user %d in MongoDB", user.ID))
    }
    return nil
}

// DeleteUser removes an user from MongoDB
func (repo usersMongoDB) DeleteUser(ctx context.Context, id int64) apierrors.APIError {
    result, err := repo.database.Collection(repo.collection).DeleteOne(ctx, bson.M{"id": id})
    if err != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error deleting user %d in MongoDB: %s", id, err.Error()))
    }
    if result.DeletedCount == 0 {
        return apierrors.NewNotFoundError(fmt.Sprintf("not found user %d in MongoDB", id))
    }
    return nil
}
