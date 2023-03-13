package users

type GetUserRequest struct {
    ID int64
}

type ListUsersRequest struct {
    Limit  int
    Offset int
}

type SaveUserRequest struct {
    User User
}

type UpdateUserRequest struct {
    User User
}

type DeleteUserRequest struct {
    ID int64
}
