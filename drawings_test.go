package tcg

import (
	"testing"
)

func TestBuffer_HLine(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.HLine(1, 5, 7, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			"..........",
			"..........",
			"..........",
			"..........",
			".*******..",
			"..........",
			"..........",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_VLine(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.VLine(5, 1, 7, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			".....*....",
			".....*....",
			".....*....",
			".....*....",
			".....*....",
			".....*....",
			".....*....",
			"..........",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_Rect(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.Rect(1, 1, 8, 8, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			".********.",
			".*......*.",
			".*......*.",
			".*......*.",
			".*......*.",
			".*......*.",
			".*......*.",
			".********.",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_FillRect(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.FillRect(1, 1, 8, 8, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			".********.",
			".********.",
			".********.",
			".********.",
			".********.",
			".********.",
			".********.",
			".********.",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}

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

func TestBuffer_LineFast(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.LineFast(0, 0, 9, 9, Black)
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
		b.LineFast(9, 9, 0, 0, Black)
		assertEqBuffers(t, b, expected)
	}
	{
		b := NewBuffer(10, 10)
		b.LineFast(0, 0, 9, 4, Black)
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
		b.LineFast(1, 9, 0, 0, Black)
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
		b.LineFast(0, 1, 9, 2, Black)
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

func BenchmarkBuffer_Line(b *testing.B) {
	buf := NewBuffer(10, 10)
	for n := 0; n < b.N; n++ {
		buf.Line(0, 0, 9, 4, Black)
	}
}

func BenchmarkBuffer_LineFast(b *testing.B) {
	buf := NewBuffer(10, 10)
	for n := 0; n < b.N; n++ {
		buf.LineFast(0, 0, 9, 4, Black)
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

func TestBuffer_CircleFast(t *testing.T) {
	{
		b := NewBuffer(10, 10)
		b.CircleFast(5, 5, 4, Black)

		expected := MustNewBufferFromStrings([]string{
			"..........",
			"....***...",
			"..**...**.",
			"..*.....*.",
			".*.......*",
			".*.......*",
			".*.......*",
			"..*.....*.",
			"..**...**.",
			"....***...",
		})
		assertEqBuffers(t, b, expected)
	}
}

func BenchmarkBuffer_Circle(b *testing.B) {
	buf := NewBuffer(10, 10)
	for n := 0; n < b.N; n++ {
		buf.Circle(5, 5, 4, Black)
	}
}

func BenchmarkBuffer_CircleFast(b *testing.B) {
	buf := NewBuffer(10, 10)
	for n := 0; n < b.N; n++ {
		buf.CircleFast(5, 5, 4, Black)
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
