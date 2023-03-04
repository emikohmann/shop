package items

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriority_Value(t *testing.T) {
	assert.Equal(t, uint8(1), PriorityLow.Value())
	assert.Equal(t, uint8(2), PriorityMedium.Value())
	assert.Equal(t, uint8(3), PriorityHigh.Value())
	assert.Equal(t, uint8(0), Priority(-1).Value())
}
