package http

import (
    "time"
    "users-api/internal/apierrors"
    "users-api/pkg/users"
)

type GetUserResponseHTTP struct {
    ID             int64     `json:"id" example:"1"`
    Email          string    `json:"email" example:"emikohmann@gmail.com"`
    Username       string    `json:"username" example:"ekohmann"`
    Password       string    `json:"password" example:"abc123"`
    ProfilePicture string    `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool      `json:"is_active" example:"true"`
    DateCreated    time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
    LastUpdated    time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type ListUsersResponseHTTP struct {
    Paging PagingResponseHTTP `json:"paging"`
    Users  []UserResponseHTTP `json:"users"`
}

type PagingResponseHTTP struct {
    Total  int `json:"total" example:"500"`
    Limit  int `json:"limit" example:"10"`
    Offset int `json:"offset" example:"50"`
}

type UserResponseHTTP struct {
    ID             int64     `json:"id" example:"1"`
    Email          string    `json:"email" example:"emikohmann@gmail.com"`
    Username       string    `json:"username" example:"ekohmann"`
    Password       string    `json:"password" example:"abc123"`
    ProfilePicture string    `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool      `json:"is_active" example:"true"`
    DateCreated    time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
    LastUpdated    time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type SaveUserResponseHTTP struct {
    ID             int64     `json:"id" example:"1"`
    Email          string    `json:"email" example:"emikohmann@gmail.com"`
    Username       string    `json:"username" example:"ekohmann"`
    Password       string    `json:"password" example:"abc123"`
    ProfilePicture string    `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool      `json:"is_active" example:"true"`
    DateCreated    time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
    LastUpdated    time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type UpdateUserResponseHTTP struct {
    ID             int64     `json:"id" example:"1"`
    Email          string    `json:"email" example:"emikohmann@gmail.com"`
    Username       string    `json:"username" example:"ekohmann"`
    Password       string    `json:"password" example:"abc123"`
    ProfilePicture string    `json:"profile_picture" example:"https://contactcenter.macstation.com.ar/web/image?unique=ed3cc51"`
    IsActive       bool      `json:"is_active" example:"true"`
    DateCreated    time.Time `json:"date_created" example:"2023-02-23T21:46:28.366Z"`
    LastUpdated    time.Time `json:"last_updated" example:"2023-02-23T21:46:28.366Z"`
}

type DeleteUserResponseHTTP struct {
    ID int64 `json:"id" example:"1"`
}

type APIErrorHTTP struct {
    Status  int    `json:"status" example:"404"`
    Message string `json:"message" example:"Some information not found"`
}

// GetUserResponseToHTTP prepares the GetUserResponse to be presented as HTTP
func GetUserResponseToHTTP(response users.GetUserResponse) GetUserResponseHTTP {
    return GetUserResponseHTTP{
        ID:             response.User.ID,
        Email:          response.User.Email,
        Username:       response.User.Username,
        Password:       response.User.Password,
        ProfilePicture: response.User.ProfilePicture,
        IsActive:       response.User.IsActive,
        DateCreated:    response.User.DateCreated,
        LastUpdated:    response.User.LastUpdated,
    }
}

// ListUsersResponseToHTTP prepares the GetUserResponse to be presented as HTTP
func ListUsersResponseToHTTP(response users.ListUsersResponse) ListUsersResponseHTTP {
    users := make([]UserResponseHTTP, 0)
    for _, user := range response.Users {
        users = append(users, UserResponseHTTP{
            ID:             user.ID,
            Email:          user.Email,
            Username:       user.Username,
            Password:       user.Password,
            ProfilePicture: user.ProfilePicture,
            IsActive:       user.IsActive,
            DateCreated:    user.DateCreated,
            LastUpdated:    user.LastUpdated,
        })
    }
    return ListUsersResponseHTTP{
        Paging: PagingResponseHTTP{
            Total:  response.Paging.Total,
            Limit:  response.Paging.Limit,
            Offset: response.Paging.Offset,
        },
        Users: users,
    }
}

// SaveUserResponseToHTTP prepares the SaveUserResponse to be presented as HTTP
func SaveUserResponseToHTTP(response users.SaveUserResponse) SaveUserResponseHTTP {
    return SaveUserResponseHTTP{
        ID:             response.User.ID,
        Email:          response.User.Email,
        Username:       response.User.Username,
        Password:       response.User.Password,
        ProfilePicture: response.User.ProfilePicture,
        IsActive:       response.User.IsActive,
        DateCreated:    response.User.DateCreated,
        LastUpdated:    response.User.LastUpdated,
    }
}

// UpdateUserResponseToHTTP prepares the UpdateUserResponse to be presented as HTTP
func UpdateUserResponseToHTTP(response users.UpdateUserResponse) UpdateUserResponseHTTP {
    return UpdateUserResponseHTTP{
        ID:             response.User.ID,
        Email:          response.User.Email,
        Username:       response.User.Username,
        Password:       response.User.Password,
        ProfilePicture: response.User.ProfilePicture,
        IsActive:       response.User.IsActive,
        DateCreated:    response.User.DateCreated,
        LastUpdated:    response.User.LastUpdated,
    }
}

// DeleteUserResponseToHTTP prepares the DeleteUserResponse to be presented as HTTP
func DeleteUserResponseToHTTP(response users.DeleteUserResponse) DeleteUserResponseHTTP {
    return DeleteUserResponseHTTP{
        ID: response.ID,
    }
}

// APIErrorToHTTP prepares the APIError to be presented as HTTP
func APIErrorToHTTP(apiError apierrors.APIError) APIErrorHTTP {
    return APIErrorHTTP{
        Status:  apiError.Status(),
        Message: apiError.Message(),
    }
}
