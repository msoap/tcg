package tcg

import (
	"testing"
)

func TestBuffer_Line(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.Line(0, 0, 9, 9, Black)
		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)

		b.Clear()
		b.Line(9, 9, 0, 0, Black)
		assertEqBuffers(t, b, expected)
	}
	{
		b := NewBuffer(10, 10)
		b.Line(0, 0, 9, 4, Black)
		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
	}
	{
		b := NewBuffer(10, 10)
		b.Line(1, 9, 0, 0, Black)
		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
	}
	{
		b := NewBuffer(10, 10)
		b.Line(0, 1, 9, 2, Black)
		expected := MustNewBufferFromStrings([]string{
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
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_Circle(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.Circle(5, 5, 4, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			"...*****..",
			"..**...**.",
			".**.....**",
			".*.......*",
			".*.......*",
			".*.......*",
			".**.....**",
			"..**...**.",
			"...*****..",
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_Arc(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.Arc(5, 5, 4, 0, 90, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			"......**..",
			".......**.",
			"........**",
			".........*",
			".........*",
			"..........",
			"..........",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
	{
		b := NewBuffer(10, 10)
		b.Arc(5, 5, 4, 90, 180, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			"...***....",
			"..**......",
			".**.......",
			".*........",
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}
