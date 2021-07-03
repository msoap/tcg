package turtle

import (
	"fmt"
	"testing"

	"github.com/msoap/tcg"
	"github.com/stretchr/testify/assert"
)

func TestTurtle_DrawScript(t *testing.T) {
	{
		b := tcg.NewBuffer(10, 10)
		New(&b).DrawScript(`
			G 2,1
			R2 D2 L2 U2 # square 3x3
			G 3,3 LT 3,3
			DL2 UL2 UR2
		`)
		expected := []string{
			"..........",
			"..***.....",
			"..*.*.....",
			"..***.....",
			".....*....",
			"......*...",
			".....*.*..",
			"....*...*.",
			".....*.*..",
			"......*...",
		}
		assertEqBuffers(t, b, tcg.MustNewBufferFromStrings(expected))
	}
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
