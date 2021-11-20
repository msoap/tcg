package tcg

import (
	"testing"
)

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
	b.HFlip()

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
	b.VFlip()

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

		b.VScroll(1)
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

		b.VScroll(3)
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

		b.VScroll(-1)
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

		b.VScroll(-3)
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

		b.HScroll(1)
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

		b.HScroll(3)
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

		b.HScroll(-1)
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

		b.HScroll(-3)
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

func TestBuffer_BitBltAllSrc(t *testing.T) {
	src := MustNewBufferFromStrings([]string{
		"..*..",
		".*.*.",
		"..*..",
	})
	b := MustNewBufferFromStrings([]string{
		"..........",
		".*******..",
		".*.....*..",
		".*.....*..",
		".*.....*..",
		".*.....*..",
		".*.....*..",
		".*******..",
		"..........",
		"..........",
	})
	b.BitBltAllSrc(2, 3, src)

	expected := MustNewBufferFromStrings([]string{
		"..........",
		".*******..",
		".*.....*..",
		".*..*..*..",
		".*.*.*.*..",
		".*..*..*..",
		".*.....*..",
		".*******..",
		"..........",
		"..........",
	})
	assertEqBuffers(t, b, expected)
}
