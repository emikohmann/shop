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
    ID           int64
    Name         string
    Description  string
    Thumbnail    string
    Images       []string
    IsActive     bool
    Restrictions []string
    Price        float64
    Stock        int
    DateCreated  time.Time
    LastUpdated  time.Time
}
