package repositories

import (
    "context"
    "database/sql"
    "errors"
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
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, database))
    if err != nil {
        logger.Errorf(ctx, "Error connecting to MySQL: %s", err.Error())
        return usersMySQL{}, err
    }

    return usersMySQL{
        database: db,
        logger:   logger,
    }, nil
}

// GetUser fetches a user from MySQL
func (repo usersMySQL) GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError) {
    statement, err := repo.database.PrepareContext(ctx, "SELECT id, email, username, password, profile_picture, is_active, date_created, last_updated FROM users WHERE id=?")
    if err != nil {
        return users.User{}, apierrors.NewInternalServerError(fmt.Sprintf("error preparing statement to get user %d from MySQL: %s", id, err.Error()))
    }
    row := statement.QueryRowContext(ctx, id)
    if row.Err() != nil {
        return users.User{}, apierrors.NewInternalServerError(fmt.Sprintf("error getting user %d from MySQL: %s", id, err.Error()))
    }
    var user users.User
    if err := row.Scan(
        &user.ID,
        &user.Email,
        &user.Username,
        &user.Password,
        &user.ProfilePicture,
        &user.IsActive,
        &user.DateCreated,
        &user.LastUpdated); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return users.User{}, apierrors.NewNotFoundError(fmt.Sprintf("not found user with id %d in MySQL", id))
        }
        return users.User{}, apierrors.NewInternalServerError(fmt.Sprintf("error scanning result of user %d from MySQL: %s", id, err.Error()))
    }
    return user, nil
}

// ListUsers fetches a list of users from MySQL
func (repo usersMySQL) ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError) {

    return users.UserList{}, apierrors.NewNotImplementedError("not implemented yet")
}

// SaveUser inserts a user into MySQL
func (repo usersMySQL) SaveUser(ctx context.Context, user users.User) apierrors.APIError {
    statement, err := repo.database.PrepareContext(ctx, "INSERT INTO users(id, email, username, password, profile_picture, is_active, date_created, last_updated) values (?,?,?,?,?,?,?,?)")
    if err != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error preparing statement to save user %d from MySQL: %s", user.ID, err.Error()))
    }
    if _, err := statement.ExecContext(ctx, user.ID, user.Email, user.Username, user.Password, user.ProfilePicture, user.IsActive, user.DateCreated, user.LastUpdated); err != nil {
        return apierrors.NewInternalServerError(fmt.Sprintf("error saving user %d in MySQL: %s", user.ID, err.Error()))
    }
    return nil
}

// UpdateUser modifies a user into MySQL
func (repo usersMySQL) UpdateUser(ctx context.Context, user users.User) apierrors.APIError {

    return apierrors.NewNotImplementedError("not implemented yet")
}

// DeleteUser removes an user from MySQL
func (repo usersMySQL) DeleteUser(ctx context.Context, id int64) apierrors.APIError {

    return apierrors.NewNotImplementedError("not implemented yet")
}
