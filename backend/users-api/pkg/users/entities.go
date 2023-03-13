package users

import "time"

type UserList struct {
    Paging Paging
    Users  []User
}

type Paging struct {
    Total  int
    Limit  int
    Offset int
}

type User struct {
    ID             int64
    Email          string
    Username       string
    Password       string
    ProfilePicture string
    IsActive       bool
    DateCreated    time.Time
    LastUpdated    time.Time
}
