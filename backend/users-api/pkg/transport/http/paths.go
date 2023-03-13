package http

import "fmt"

const (
    paramUserID = "userID"
)

const (
    queryLimit  = "limit"
    queryOffset = "offset"
)

var (
    PathDocs       = "/docs/*any"
    PathMetrics    = "/metrics"
    PathGetUser    = fmt.Sprintf("/users/:%s", paramUserID)
    PathListUsers  = "/users"
    PathSaveUser   = "/users"
    PathUpdateUser = fmt.Sprintf("/users/:%s", paramUserID)
    PathDeleteUser = fmt.Sprintf("/users/:%s", paramUserID)
)
