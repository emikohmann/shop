package repositories

import (
    "context"
    "errors"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/users"
)

type usersMySQL struct {
    database *gorm.DB
    table    string
    logger   *logger.Logger
}

type mySQLUser struct {
    ID             int64     `gorm:"column:id"`
    Email          string    `gorm:"column:email"`
    Username       string    `gorm:"column:username"`
    Password       string    `gorm:"column:password"`
    ProfilePicture string    `gorm:"column:profile_picture"`
    IsActive       bool      `gorm:"column:is_active"`
    DateCreated    time.Time `gorm:"column:date_created"`
    LastUpdated    time.Time `gorm:"column:last_updated"`
}

func (user mySQLUser) toUser() users.User {
    return users.User{
        ID:             user.ID,
        Email:          user.Email,
        Username:       user.Username,
        Password:       user.Password,
        ProfilePicture: user.ProfilePicture,
        IsActive:       user.IsActive,
        DateCreated:    user.DateCreated,
        LastUpdated:    user.LastUpdated,
    }
}

func userToMySQLUser(user users.User) mySQLUser {
    return mySQLUser{
        ID:             user.ID,
        Email:          user.Email,
        Username:       user.Username,
        Password:       user.Password,
        ProfilePicture: user.ProfilePicture,
        IsActive:       user.IsActive,
        DateCreated:    user.DateCreated,
        LastUpdated:    user.LastUpdated,
    }
}

// NewUsersMySQL instances a new users' repository against MySQL
func NewUsersMySQL(ctx context.Context, host string, port int, database string, user string, password string, table string, logger *logger.Logger) (usersMySQL, error) {
    datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
    db, err := gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
    if err != nil {
        logger.Errorf(ctx, "Error connecting to MySQL: %s", err.Error())
        return usersMySQL{}, err
    }

    return usersMySQL{
        database: db,
        table:    table,
        logger:   logger,
    }, nil
}

// GetUser fetches a user from MySQL
func (repo usersMySQL) GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError) {
    var mySQLUser mySQLUser
    txn := repo.database.Table(repo.table).First(&mySQLUser, id)
    if txn.Error != nil {
        if errors.Is(txn.Error, gorm.ErrRecordNotFound) {
            return users.User{}, apierrors.NewNotFoundError(fmt.Sprintf("not found user with id %d in MySQL", id))
        }
        return users.User{}, apierrors.NewInternalServerError(fmt.Sprintf("error getting user %d from MySQL: %s", id, txn.Error.Error()))
    }
    return mySQLUser.toUser(), nil
}

// ListUsers fetches a list of users from MySQL
func (repo usersMySQL) ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError) {
    var count int64
    txnCount := repo.database.Table(repo.table).Count(&count)
    if txnCount.Error != nil {
        return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error counting users in list from MySQL: %s", txnCount.Error.Error()))
    }
    mySQLUsers := make([]mySQLUser, 0)
    txn := repo.database.Table(repo.table).Limit(limit).Offset(offset).Find(&mySQLUsers)
    if txn.Error != nil {
        return users.UserList{}, apierrors.NewInternalServerError(fmt.Sprintf("error listing users from MySQL: %s", txn.Error.Error()))
    }

    userList := make([]users.User, 0)
    for _, mySQLUser := range mySQLUsers {
        userList = append(userList, mySQLUser.toUser())
    }
    return users.UserList{
        Paging: users.Paging{
            Total:  int(count),
            Limit:  limit,
            Offset: offset,
        },
        Users: userList,
    }, nil
}

// SaveUser inserts a user into MySQL
func (repo usersMySQL) SaveUser(ctx context.Context, user users.User) apierrors.APIError {
    txn := repo.database.Table(repo.table).Save(userToMySQLUser(user))
    if txn.Error != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error saving user %d in MySQL: %s", user.ID, txn.Error.Error()))
    }
    return nil
}

// UpdateUser modifies a user into MySQL
func (repo usersMySQL) UpdateUser(ctx context.Context, user users.User) apierrors.APIError {
    // save with ID is equivalent to update for GORM
    txn := repo.database.Table(repo.table).Save(userToMySQLUser(user))
    if txn.Error != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error updating user %d in MySQL: %s", user.ID, txn.Error.Error()))
    }
    return nil
}

// DeleteUser removes an user from MySQL
func (repo usersMySQL) DeleteUser(ctx context.Context, id int64) apierrors.APIError {
    txn := repo.database.Table(repo.table).Delete(&mySQLUser{}, id)
    if txn.Error != nil {
        if errors.Is(txn.Error, gorm.ErrRecordNotFound) {
            return apierrors.NewNotFoundError(fmt.Sprintf("not found user with id %d in MySQL", id))
        }
        return apierrors.NewInternalServerError(fmt.Sprintf("error deleting user %d in MySQL: %s", id, txn.Error.Error()))
    }
    return nil
}
