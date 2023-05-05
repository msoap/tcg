package sprite

import (
	"fmt"
	"testing"

	"github.com/msoap/tcg"
	"github.com/stretchr/testify/assert"
)

var chess = tcg.MustNewBufferFromStrings([]string{
	"*.",
	".*",
})

func TestSprite(t *testing.T) {
	bg := tcg.NewBuffer(5, 5)
	bg.Fill(0, 0, tcg.WithPattern(chess))

	spr := New(tcg.MustNewBufferFromStrings([]string{
		"**",
		"**",
	}))

	spr.Put(bg, 1, 1)
	assertEqBuffers(t, bg, tcg.MustNewBufferFromStrings([]string{
		"*.*.*",
		".***.",
		"***.*",
		".*.*.",
		"*.*.*",
	}))

	spr.Withdraw(bg)
	assertEqBuffers(t, bg, tcg.MustNewBufferFromStrings([]string{
		"*.*.*",
		".*.*.",
		"*.*.*",
		".*.*.",
		"*.*.*",
	}))

	spr.Put(bg, 1, 1)

	spr.Move(bg, 1, 1)
	assertEqBuffers(t, bg, tcg.MustNewBufferFromStrings([]string{
		"*.*.*",
		".*.*.",
		"*.***",
		".***.",
		"*.*.*",
	}))

	spr.MoveAbs(bg, 0, -1)
	assertEqBuffers(t, bg, tcg.MustNewBufferFromStrings([]string{
		"***.*",
		".*.*.",
		"*.*.*",
		".*.*.",
		"*.*.*",
	}))
}

func assertEqBuffers(t *testing.T, got, expected tcg.Buffer) {
	if got.Width != expected.Width {
		t.Errorf("buffer width of got (%d) != expected (%d)", got.Width, expected.Width)
		return
	}
	if got.Height != expected.Height {
		t.Errorf("buffer height of got (%d) != expected (%d)", got.Height, expected.Height)
		return
	}

	if !expected.IsEqual(got) {
		gotStrings, expectedStrings := got.Strings(), expected.Strings()
		msg := fmt.Sprintf("buffers isn't equal:\n%-*s | %-*s\n", got.Width, "got", expected.Width, "expected")
		for y := 0; y < got.Height; y++ {
			msg += fmt.Sprintf("%-*s | %-*s\n", got.Width, gotStrings[y], expected.Width, expectedStrings[y])
		}
		assert.True(t, false, msg)
	}
}
