package tcg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuffer_Line(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.Line(0, 0, 9, 9, Black)
		expected := []string{
			"*.........",
			".*........",
			"..*.......",
			"...*......",
			"....*.....",
			".....*....",
			"......*...",
			".......*..",
			"........*.",
			".........*",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))

		b = NewBuffer(10, 10)
		b.Line(9, 9, 0, 0, Black)
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
	{
		b := NewBuffer(10, 10)
		b.Line(0, 0, 9, 4, Black)
		expected := []string{
			"**........",
			"..**......",
			"....**....",
			"......**..",
			"........**",
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
	{
		b := NewBuffer(10, 10)
		b.Line(1, 9, 0, 0, Black)
		expected := []string{
			"*.........",
			"*.........",
			"*.........",
			"*.........",
			"*.........",
			".*........",
			".*........",
			".*........",
			".*........",
			".*........",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
	{
		b := NewBuffer(10, 10)
		b.Line(0, 1, 9, 2, Black)
		expected := []string{
			"..........",
			"*****.....",
			".....*****",
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
}
