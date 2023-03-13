package http

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParams(t *testing.T) {
    assert.Equal(t, "userID", paramUserID)
}

func TestQueries(t *testing.T) {
    assert.Equal(t, "limit", queryLimit)
    assert.Equal(t, "offset", queryOffset)
}

func TestPaths(t *testing.T) {
    assert.Equal(t, "/docs/*any", PathDocs)
    assert.Equal(t, "/metrics", PathMetrics)
    assert.Equal(t, fmt.Sprintf("/users/:%s", paramUserID), PathGetUser)
    assert.Equal(t, "/users", PathListUsers)
    assert.Equal(t, "/users", PathSaveUser)
    assert.Equal(t, fmt.Sprintf("/users/:%s", paramUserID), PathUpdateUser)
    assert.Equal(t, fmt.Sprintf("/users/:%s", paramUserID), PathDeleteUser)
}
