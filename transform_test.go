package tcg

import (
	"testing"
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

		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
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

		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
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

	expected := MustNewBufferFromStrings([]string{
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
	})
	assertEqBuffers(t, b, expected)
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

	expected := MustNewBufferFromStrings([]string{
		"..........",
		"..........",
		"..........",
		"....*.....",
		"...*.*....",
		"..*...*...",
		".*.....*..",
		"*********.",
		"..........",
	})
	assertEqBuffers(t, b, expected)
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

	expected := MustNewBufferFromStrings([]string{
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
	})
	assertEqBuffers(t, b, expected)
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
		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)

		b.ScrollV(3)
		expected = MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
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
		expected := MustNewBufferFromStrings([]string{
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
		assertEqBuffers(t, b, expected)

		b.ScrollV(-3)
		expected = MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
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
