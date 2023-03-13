package http

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "net/http"
    _ "users-api/docs/openapi"
    "users-api/internal/apierrors"
    "users-api/internal/logger"
    "users-api/pkg/users"
)

type UsersService interface {
    GetUser(ctx context.Context, id int64) (users.User, apierrors.APIError)
    ListUsers(ctx context.Context, limit int, offset int) (users.UserList, apierrors.APIError)
    SaveUser(ctx context.Context, user users.User) (users.User, apierrors.APIError)
    UpdateUser(ctx context.Context, user users.User) (users.User, apierrors.APIError)
    DeleteUser(ctx context.Context, id int64) apierrors.APIError
}

// DocsHandler sets up the Docs request handler
func DocsHandler(logger *logger.Logger) gin.HandlerFunc {
    handler := ginSwagger.WrapHandler(swaggerFiles.Handler)
    return handler
}

// MetricsHandler sets up the Metrics request handler
func MetricsHandler(logger *logger.Logger) gin.HandlerFunc {
    handler := promhttp.Handler()
    return func(ctx *gin.Context) {
        handler.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

// GetUserHandler sets up the GetUser request handler
// GetUser godoc
//	@Summary		Return the user information.
//	@Description	Return the user information fetching information from the database.
//	@Tags			Users
//	@Param			userID	path	int	true	"ID of the user to get"
//	@Produce		json
//	@Success		200	{object}	GetUserResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		404	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/users/{userID} [get]
func GetUserHandler(ctx context.Context, usersService UsersService, logger *logger.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        request, err := HTTPToGetUserRequest(ctx)
        if err != nil {
            apiErr := apierrors.NewBadRequestError(err.Error())
            logger.Errorf(ctx, "Error generating GetUserRequest: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        user, apiErr := usersService.GetUser(ctx, request.ID)
        if apiErr != nil {
            logger.Errorf(ctx, "Error getting user: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        response := users.GetUserResponse{
            User: user,
        }
        httpResponse := GetUserResponseToHTTP(response)
        ctx.JSON(http.StatusOK, httpResponse)
    }
}

// ListUsersHandler sets up the ListUsers request handler
// ListUsers godoc
//	@Summary		Return a list of users.
//	@Description	Return the users information fetching information from the database.
//	@Tags			Users
//	@Param			limit	query	int	true	"List limit"
//	@Param			offset	query	int	true	"List offset"
//	@Produce		json
//	@Success		200	{object}	ListUsersResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/users [get]
func ListUsersHandler(ctx context.Context, usersService UsersService, logger *logger.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        request, err := HTTPToListUsersRequest(ctx)
        if err != nil {
            apiErr := apierrors.NewBadRequestError(err.Error())
            logger.Errorf(ctx, "Error generating ListUsersRequest: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        list, apiErr := usersService.ListUsers(ctx, request.Limit, request.Offset)
        if apiErr != nil {
            logger.Errorf(ctx, "Error listing users: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        response := users.ListUsersResponse{
            Paging: list.Paging,
            Users:  list.Users,
        }
        httpResponse := ListUsersResponseToHTTP(response)
        ctx.JSON(http.StatusOK, httpResponse)
    }
}

// SaveUserHandler sets up the SaveUser request handler
// SaveUser godoc
//	@Summary		Store the user information.
//	@Description	Store the user information against the database.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		SaveUserRequestHTTP	true	"User to save"
//	@Success		201		{object}	SaveUserResponseHTTP
//	@Failure		400		{object}	APIErrorHTTP
//	@Failure		500		{object}	APIErrorHTTP
//	@Router			/users [post]
func SaveUserHandler(ctx context.Context, usersService UsersService, logger *logger.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        request, err := HTTPToSaveUserRequest(ctx)
        if err != nil {
            apiErr := apierrors.NewBadRequestError(err.Error())
            logger.Errorf(ctx, "Error generating SaveUserRequest: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        user, apiErr := usersService.SaveUser(ctx, request.User)
        if apiErr != nil {
            logger.Errorf(ctx, "Error saving user: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        response := users.SaveUserResponse{
            User: user,
        }
        httpResponse := SaveUserResponseToHTTP(response)
        ctx.JSON(http.StatusCreated, httpResponse)
    }
}

// UpdateUserHandler sets up the UpdateUser request handler
// UpdateUser godoc
//	@Summary		Updates the user information.
//	@Description	Updates the user information against the database.
//	@Tags			Users
//	@Param			userID	path	int	true	"ID of the user to get"
//	@Accept			json
//	@Produce		json
//	@Param			request	body		UpdateUserRequestHTTP	true	"User fields to update"
//	@Success		200		{object}	UpdateUserResponseHTTP
//	@Failure		400		{object}	APIErrorHTTP
//	@Failure		500		{object}	APIErrorHTTP
//	@Router			/users/{userID} [put]
func UpdateUserHandler(ctx context.Context, usersService UsersService, logger *logger.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        request, err := HTTPToUpdateUserRequest(ctx)
        if err != nil {
            apiErr := apierrors.NewBadRequestError(err.Error())
            logger.Errorf(ctx, "Error generating UpdateUserRequest: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        user, apiErr := usersService.UpdateUser(ctx, request.User)
        if apiErr != nil {
            logger.Errorf(ctx, "Error updating user: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        response := users.UpdateUserResponse{
            User: user,
        }
        httpResponse := UpdateUserResponseToHTTP(response)
        ctx.JSON(http.StatusOK, httpResponse)
    }
}

// DeleteUserHandler sets up the DeleteUser request handler
// DeleteUser godoc
//	@Summary		Delete the user information.
//	@Description	Delete the user information against the database.
//	@Tags			Users
//	@Param			userID	path	int	true	"ID of the user to delete"
//	@Produce		json
//	@Success		200	{object}	DeleteUserResponseHTTP
//	@Failure		400	{object}	APIErrorHTTP
//	@Failure		404	{object}	APIErrorHTTP
//	@Failure		500	{object}	APIErrorHTTP
//	@Router			/users/{userID} [delete]
func DeleteUserHandler(ctx context.Context, usersService UsersService, logger *logger.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        request, err := HTTPToDeleteUserRequest(ctx)
        if err != nil {
            apiErr := apierrors.NewBadRequestError(err.Error())
            logger.Errorf(ctx, "Error generating DeleteUserRequest: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        if apiErr := usersService.DeleteUser(ctx, request.ID); apiErr != nil {
            logger.Errorf(ctx, "Error deleting user: %s", apiErr.Error())
            httpResponse := APIErrorToHTTP(apiErr)
            ctx.JSON(apiErr.Status(), httpResponse)
            return
        }
        response := users.DeleteUserResponse{
            ID: request.ID,
        }
        httpResponse := DeleteUserResponseToHTTP(response)
        ctx.JSON(http.StatusOK, httpResponse)
    }
}
