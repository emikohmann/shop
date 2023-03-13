package users

type GetUserResponse struct {
    User User
}

type ListUsersResponse struct {
    Paging Paging
    Users  []User
}

type SaveUserResponse struct {
    User User
}

type UpdateUserResponse struct {
    User User
}

type DeleteUserResponse struct {
    ID int64
}
