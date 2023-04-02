package tcg

import "testing"

func TestBuffer_BitBltAllSrc(t *testing.T) {
	t.Run("same buffers", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			"..*..",
			".*.*.",
			"..*..",
		})
		b := MustNewBufferFromStrings([]string{
			"..*..",
			"*****",
			"..*..",
		})
		b.BitBltAll(0, 0, src)

		expected := MustNewBufferFromStrings([]string{
			"..*..",
			".*.*.",
			"..*..",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("different buffers", func(t *testing.T) {
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
		b.BitBltAll(2, 3, src)

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
	})
}

func TestBuffer_BitBlt(t *testing.T) {
	b := MustNewBufferFromStrings([]string{
		"..........",
		".*******..",
		".*.*.*.*..",
		".*******..",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
		"..........",
	})
	b.BitBlt(3, 5, 5, 4, b, 1, 1)

	expected := MustNewBufferFromStrings([]string{
		"..........",
		".*******..",
		".*.*.*.*..",
		".*******..",
		"..........",
		"...*****..",
		"...*.*.*..",
		"...*****..",
		"..........",
		"..........",
	})
	assertEqBuffers(t, b, expected)
}
