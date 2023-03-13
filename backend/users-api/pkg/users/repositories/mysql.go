package repositories

import (
    "context"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/users"
)

type usersMySQL struct {
    database *sql.DB
    logger   *logger.Logger
}

// NewUsersMySQL instances a new users' repository against MySQL
func NewUsersMySQL(ctx context.Context, host string, port int, database string, user string, password string, logger *logger.Logger) (usersMySQL, error) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database))
    if err != nil {
        logger.Errorf(ctx, "Error connecting to MySQL: %s", err.Error())
        return usersMySQL{}, err
    }

    return usersMySQL{
        database: db,
        logger:   logger,
    }, nil
}

// GetUser fetches an user from MySQL
func (repo usersMySQL) GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError) {

    return users.User{}, apierrors.NewNotImplementedError("not implemented yet")
}

// ListUsers fetches a list of users from MySQL
func (repo usersMySQL) ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError) {

    return users.UserList{}, apierrors.NewNotImplementedError("not implemented yet")
}

// SaveUser inserts a user into MySQL
func (repo usersMySQL) SaveUser(ctx context.Context, user users.User) apierrors.APIError {

    return apierrors.NewNotImplementedError("not implemented yet")
}

// UpdateUser modifies a user into MySQL
func (repo usersMySQL) UpdateUser(ctx context.Context, user users.User) apierrors.APIError {

    return apierrors.NewNotImplementedError("not implemented yet")
}

// DeleteUser removes an user from MySQL
func (repo usersMySQL) DeleteUser(ctx context.Context, id int64) apierrors.APIError {

    return apierrors.NewNotImplementedError("not implemented yet")
}
