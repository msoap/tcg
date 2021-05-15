package tcg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
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
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
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
	assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
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
	assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
}

func TestBuffer_FlipV(t *testing.T) {
	b := MustNewBufferFromStrings([]string{
		"..........",
		"..........",
		".*........",
		".**.......",
		".*.*......",
		".*..*.....",
		".*.*......",
		".**.......",
		".*........",
		"..........",
	})
	b.FlipV()

	expected := []string{
		"..........",
		"..........",
		"........*.",
		".......**.",
		"......*.*.",
		".....*..*.",
		"......*.*.",
		".......**.",
		"........*.",
		"..........",
	}
	assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
}

func TestBuffer_ScrollV(t *testing.T) {
	{
		b := MustNewBufferFromStrings([]string{
			".*........",
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
			".*........",
			"..........",
			"..........",
			"..........",
		})

		b.ScrollV(1)
		expected := []string{
			"..........",
			".*........",
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
			".*........",
			"..........",
			"..........",
		}
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))

		b.ScrollV(3)
		expected = []string{
			"..........",
			"..........",
			"..........",
			"..........",
			".*........",
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
		}
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
	}
	{
		b := MustNewBufferFromStrings([]string{
			"..........",
			"..........",
			"..........",
			".*........",
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
			".*........",
		})

		b.ScrollV(-1)
		expected := []string{
			"..........",
			"..........",
			".*........",
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
			".*........",
			"..........",
		}
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))

		b.ScrollV(-3)
		expected = []string{
			".**.......",
			".*.*......",
			".*..*.....",
			".*.*......",
			".**.......",
			".*........",
			"..........",
			"..........",
			"..........",
			"..........",
		}
		assert.True(t, MustNewBufferFromStrings(expected).IsEqual(b), "expected:\n"+strings.Join(b.Strings(), "\n"))
	}
}

func TestBuffer_ScrollH(t *testing.T) {
	{
		b := MustNewBufferFromStrings([]string{
			".*........",
			".**.......",
			"**.*......",
			"**..*.....",
			"**.*......",
			".**.......",
			".*........",
			"..........",
			"..........",
			"..........",
		})

		b.ScrollH(1)
		expected := MustNewBufferFromStrings([]string{
			"..*.......",
			"..**......",
			".**.*.....",
			".**..*....",
			".**.*.....",
			"..**......",
			"..*.......",
			"..........",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)

		b.ScrollH(3)
		expected = MustNewBufferFromStrings([]string{
			".....*....",
			".....**...",
			"....**.*..",
			"....**..*.",
			"....**.*..",
			".....**...",
			".....*....",
			"..........",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
	{
		b := MustNewBufferFromStrings([]string{
			"..........",
			"..........",
			"........*.",
			".......**.",
			"......*.**",
			".....*..**",
			"......*.**",
			".......**.",
			"........*.",
			"..........",
		})

		b.ScrollH(-1)
		expected := MustNewBufferFromStrings([]string{
			"..........",
			"..........",
			".......*..",
			"......**..",
			".....*.**.",
			"....*..**.",
			".....*.**.",
			"......**..",
			".......*..",
			"..........",
		})
		assertEqBuffers(t, b, expected)

		b.ScrollH(-3)
		expected = MustNewBufferFromStrings([]string{
			"..........",
			"..........",
			"....*.....",
			"...**.....",
			"..*.**....",
			".*..**....",
			"..*.**....",
			"...**.....",
			"....*.....",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}
