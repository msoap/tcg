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
	t.Run("simeple from self", func(t *testing.T) {
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
	})

	t.Run("simeple from other", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*.*.*.",
			".*****.",
			".......",
		})
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
		b.BitBlt(3, 5, 5, 3, src, 1, 1)

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
	})

	t.Run("with transparent", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			".**.**.",
			".......",
			".**.**.",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBTransparent())

		expected := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*****.",
			".*****.",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("with mask", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		mask := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*...*.",
			"...*...",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			".**.**.",
			".......",
			".**.**.",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBMask(&mask))

		expected := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*...*.",
			".*****.",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("with custom operation", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			".**.**.",
			".*...*.",
			".**.**.",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBOpFn(func(src, dst int) int {
			return src | dst
		}))

		expected := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*****.",
			".*****.",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("with AND operation", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*...*.",
			".*****.",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBAnd())

		expected := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*...*.",
			"...*...",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("with OR operation", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			".**.**.",
			".*...*.",
			".**.**.",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBOr())

		expected := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".*****.",
			".*****.",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})

	t.Run("with XOR operation", func(t *testing.T) {
		src := MustNewBufferFromStrings([]string{
			".......",
			"...*...",
			".*****.",
			"...*...",
			".......",
		})
		b := MustNewBufferFromStrings([]string{
			".......",
			".*****.",
			".**.**.",
			".*****.",
			".......",
		})
		b.BitBlt(1, 1, 5, 3, src, 1, 1, BBXor())

		expected := MustNewBufferFromStrings([]string{
			".......",
			".**.**.",
			"...*...",
			".**.**.",
			".......",
		})
		assertEqBuffers(t, b, expected)
	})
}
