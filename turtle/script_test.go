package turtle

import (
	"testing"

	"github.com/msoap/tcg"
	"github.com/stretchr/testify/require"
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
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, tcg.MustNewBufferFromStrings(expected).IsEqual(b))
	}
}
