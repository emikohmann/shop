package http

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "strconv"
    "users-api/pkg/users"
)

// HTTPToGetUserRequest turns the HTTP request into a GetUserRequest
func HTTPToGetUserRequest(ctx *gin.Context) (users.GetUserRequest, error) {
    userIDStr := ctx.Param(paramUserID)
    userID, err := strconv.ParseInt(userIDStr, 10, 64)
    if err != nil {
        return users.GetUserRequest{}, fmt.Errorf("invalid user ID: %w", err)
    }
    return users.GetUserRequest{
        ID: userID,
    }, nil
}

// HTTPToListUsersRequest turns the HTTP request into a ListUsersRequest
func HTTPToListUsersRequest(ctx *gin.Context) (users.ListUsersRequest, error) {
    limitStr := ctx.Query(queryLimit)
    limit, err := strconv.ParseInt(limitStr, 10, 64)
    if err != nil {
        return users.ListUsersRequest{}, fmt.Errorf("invalid limit value: %w", err)
    }
    offsetStr := ctx.Query(queryOffset)
    offset, err := strconv.ParseInt(offsetStr, 10, 64)
    if err != nil {
        return users.ListUsersRequest{}, fmt.Errorf("invalid offset value: %w", err)
    }
    return users.ListUsersRequest{
        Limit:  int(limit),
        Offset: int(offset),
    }, nil
}

type SaveUserRequestHTTP struct {
    ID             int64  `json:"id" example:"1"`
    Email          string `json:"email" example:"emikohmann@gmail.com"`
    Username       string `json:"username" example:"ekohmann"`
    Password       string `json:"password" example:"abc123"`
    ProfilePicture string `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool   `json:"is_active" example:"true"`
}

// HTTPToSaveUserRequest turns the HTTP Request into a SaveUserRequest
func HTTPToSaveUserRequest(ctx *gin.Context) (users.SaveUserRequest, error) {
    var saveUserRequestHTTP SaveUserRequestHTTP
    if err := ctx.ShouldBindJSON(&saveUserRequestHTTP); err != nil {
        return users.SaveUserRequest{}, fmt.Errorf("invalid save user request: %w", err)
    }
    return users.SaveUserRequest{
        User: users.User{
            ID:             saveUserRequestHTTP.ID,
            Email:          saveUserRequestHTTP.Email,
            Username:       saveUserRequestHTTP.Username,
            Password:       saveUserRequestHTTP.Password,
            ProfilePicture: saveUserRequestHTTP.ProfilePicture,
            IsActive:       saveUserRequestHTTP.IsActive,
        },
    }, nil
}

type UpdateUserRequestHTTP struct {
    ID             int64  `json:"id" example:"1"`
    Email          string `json:"email" example:"emikohmann@gmail.com"`
    Username       string `json:"username" example:"ekohmann"`
    Password       string `json:"password" example:"abc123"`
    ProfilePicture string `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool   `json:"is_active" example:"true"`
}

// HTTPToUpdateUserRequest turns the HTTP Request into an UpdateUserRequest
func HTTPToUpdateUserRequest(ctx *gin.Context) (users.UpdateUserRequest, error) {
    userIDStr := ctx.Param(paramUserID)
    userID, err := strconv.ParseInt(userIDStr, 10, 64)
    if err != nil {
        return users.UpdateUserRequest{}, fmt.Errorf("invalid user ID: %w", err)
    }
    var updateUserRequestHTTP UpdateUserRequestHTTP
    if err := ctx.ShouldBindJSON(&updateUserRequestHTTP); err != nil {
        return users.UpdateUserRequest{}, fmt.Errorf("invalid update user request: %w", err)
    }
    return users.UpdateUserRequest{
        User: users.User{
            ID:             userID,
            Email:          updateUserRequestHTTP.Email,
            Username:       updateUserRequestHTTP.Username,
            Password:       updateUserRequestHTTP.Password,
            ProfilePicture: updateUserRequestHTTP.ProfilePicture,
            IsActive:       updateUserRequestHTTP.IsActive,
        },
    }, nil
}

// HTTPToDeleteUserRequest turns the HTTP request into a DeleteUserRequest
func HTTPToDeleteUserRequest(ctx *gin.Context) (users.DeleteUserRequest, error) {
    userIDStr := ctx.Param(paramUserID)
    userID, err := strconv.ParseInt(userIDStr, 10, 64)
    if err != nil {
        return users.DeleteUserRequest{}, fmt.Errorf("invalid user ID: %w", err)
    }
    return users.DeleteUserRequest{
        ID: userID,
    }, nil
}
