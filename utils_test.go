package tcg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AbsSgn(t *testing.T) {
	assert.Equal(t, 5, abs(5))
	assert.Equal(t, 5, abs(-5))

	assert.Equal(t, -1, sgn(-5))
	assert.Equal(t, 1, sgn(5))
	assert.Equal(t, 0, sgn(0))
}
