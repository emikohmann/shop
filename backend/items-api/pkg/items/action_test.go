package items

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAction_String(t *testing.T) {
	assert.Equal(t, "GET", ActionGet.String())
	assert.Equal(t, "SAVE", ActionSave.String())
	assert.Equal(t, "UPDATE", ActionUpdate.String())
	assert.Equal(t, "DELETE", ActionDelete.String())
	assert.Equal(t, "UNKNOWN", Action(-1).String())
}
