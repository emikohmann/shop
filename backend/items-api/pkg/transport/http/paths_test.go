package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParams(t *testing.T) {
	assert.Equal(t, "itemID", paramItemID)
}

func TestPaths(t *testing.T) {
	assert.Equal(t, "/docs/*any", PathDocs)
	assert.Equal(t, "/metrics", PathMetrics)
	assert.Equal(t, fmt.Sprintf("/items/:%s", paramItemID), PathGetItem)
	assert.Equal(t, "/items", PathSaveItem)
	assert.Equal(t, fmt.Sprintf("/items/:%s", paramItemID), PathUpdateItem)
	assert.Equal(t, fmt.Sprintf("/items/:%s", paramItemID), PathDeleteItem)
}
