package tcg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxAbsSgn(t *testing.T) {
	assert.Equal(t, 3, min(3, 5))
	assert.Equal(t, 3, min(3, 3))

	assert.Equal(t, 5, max(3, 5))
	assert.Equal(t, 5, max(5, 5))

	assert.Equal(t, 5, abs(5))
	assert.Equal(t, 5, abs(-5))

	assert.Equal(t, -1, sgn(-5))
	assert.Equal(t, 1, sgn(5))
	assert.Equal(t, 0, sgn(0))
}
