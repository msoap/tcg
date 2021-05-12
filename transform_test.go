package tcg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuffer_Fill(t *testing.T) {
	{
		b := MustNewBufferFromStrings([]string{
			"..........",
			".*******..",
			".*.....*..",
			".*.....*..",
			".*.....*..",
			".*.....*..",
			".*******..",
			"..........",
			"..........",
			"..........",
		})
		b.Fill(3, 3)

		expected := []string{
			"..........",
			".*******..",
			".*******..",
			".*******..",
			".*******..",
			".*******..",
			".*******..",
			"..........",
			"..........",
			"..........",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
	{
		b := MustNewBufferFromStrings([]string{
			"..........",
			".*******..",
			".*.....*..",
			".*...**...",
			".*....**..",
			".*......*.",
			".*.***.*..",
			".*.*.*.*..",
			"..*..***..",
			"..........",
		})
		b.Fill(3, 3)

		expected := []string{
			"..........",
			".*******..",
			".*******..",
			".******...",
			".*******..",
			".********.",
			".*******..",
			".***.***..",
			"..*..***..",
			"..........",
		}
		// t.Log("\n" + strings.Join(b.Strings(), "\n"))
		require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
	}
}

func TestBuffer_Invert(t *testing.T) {
	b := MustNewBufferFromStrings([]string{
		"..........",
		".*******..",
		".*.....*..",
		".*.....*..",
		".*.....*..",
		".*.....*..",
		".*******..",
		"..........",
		"..........",
		"..........",
	})
	b.Invert()

	expected := []string{
		"**********",
		"*.......**",
		"*.*****.**",
		"*.*****.**",
		"*.*****.**",
		"*.*****.**",
		"*.......**",
		"**********",
		"**********",
		"**********",
	}
	// t.Log("\n" + strings.Join(b.Strings(), "\n"))
	require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
}

func TestBuffer_FlipH(t *testing.T) {
	b := MustNewBufferFromStrings([]string{
		"..........",
		"*********.",
		".*.....*..",
		"..*...*...",
		"...*.*....",
		"....*.....",
		"..........",
		"..........",
		"..........",
	})
	b.FlipH()

	expected := []string{
		"..........",
		"..........",
		"..........",
		"....*.....",
		"...*.*....",
		"..*...*...",
		".*.....*..",
		"*********.",
		"..........",
	}
	// t.Log("\n" + strings.Join(b.Strings(), "\n"))
	require.True(t, MustNewBufferFromStrings(expected).IsEqual(b))
}
