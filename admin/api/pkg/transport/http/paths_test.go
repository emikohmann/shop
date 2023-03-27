package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaths(t *testing.T) {
	assert.Equal(t, "/docs/*any", PathDocs)
}
