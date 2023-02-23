package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, err := NewApplication()
	assert.NoError(t, err)
	assert.NotNil(t, app)
}
