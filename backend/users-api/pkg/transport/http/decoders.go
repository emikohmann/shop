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
    ID           int64    `json:"id" example:"1"`
    Name         string   `json:"name" example:"Iphone 13 128GB 4GB RAM"`
    Description  string   `json:"description" example:"The iPhone 13 display has rounded corners"`
    Thumbnail    string   `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    Images       []string `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
    IsActive     bool     `json:"is_active" example:"true"`
    Restrictions []string `json:"restrictions"`
    Price        float64  `json:"price" example:"729.99"`
    Stock        int      `json:"stock" example:"1"`
}

// HTTPToSaveUserRequest turns the HTTP Request into a SaveUserRequest
func HTTPToSaveUserRequest(ctx *gin.Context) (users.SaveUserRequest, error) {
    var saveUserRequestHTTP SaveUserRequestHTTP
    if err := ctx.ShouldBindJSON(&saveUserRequestHTTP); err != nil {
        return users.SaveUserRequest{}, fmt.Errorf("invalid save user request: %w", err)
    }
    return users.SaveUserRequest{
        User: users.User{
            ID:           saveUserRequestHTTP.ID,
            Name:         saveUserRequestHTTP.Name,
            Description:  saveUserRequestHTTP.Description,
            Thumbnail:    saveUserRequestHTTP.Thumbnail,
            Images:       saveUserRequestHTTP.Images,
            IsActive:     saveUserRequestHTTP.IsActive,
            Restrictions: saveUserRequestHTTP.Restrictions,
            Price:        saveUserRequestHTTP.Price,
            Stock:        saveUserRequestHTTP.Stock,
        },
    }, nil
}

type UpdateUserRequestHTTP struct {
    Name         string   `json:"name" example:"Iphone 13 128GB 4GB RAM"`
    Description  string   `json:"description" example:"The iPhone 13 display has rounded corners"`
    Thumbnail    string   `json:"thumbnail" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    Images       []string `json:"images" example:"https://www.macstation.com.ar/img/productos/2599-2.jpg"`
    IsActive     bool     `json:"is_active" example:"true"`
    Restrictions []string `json:"restrictions"`
    Price        float64  `json:"price" example:"729.99"`
    Stock        int      `json:"stock" example:"1"`
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
            ID:           userID,
            Name:         updateUserRequestHTTP.Name,
            Description:  updateUserRequestHTTP.Description,
            Thumbnail:    updateUserRequestHTTP.Thumbnail,
            Images:       updateUserRequestHTTP.Images,
            IsActive:     updateUserRequestHTTP.IsActive,
            Restrictions: updateUserRequestHTTP.Restrictions,
            Price:        updateUserRequestHTTP.Price,
            Stock:        updateUserRequestHTTP.Stock,
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
