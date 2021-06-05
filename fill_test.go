package tcg

import "testing"

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

func TestBuffer_FillWithPattern(t *testing.T) {
	pattern := MustNewBufferFromStrings([]string{
		".*",
		"*.",
	})

	{
		b := MustNewBufferFromStrings([]string{
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
		b.Fill(2, 2, WithPattern(pattern))

		expected := MustNewBufferFromStrings([]string{
			"..........",
			".********.",
			".*.*.*.**.",
			".**.*.*.*.",
			".*.*.*.**.",
			".**.*.*.*.",
			".*.*.*.**.",
			".**.*.*.*.",
			".********.",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
	{
		b := MustNewBufferFromStrings([]string{
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
		b.Fill(3, 2, WithPattern(pattern))

		expected := MustNewBufferFromStrings([]string{
			"..........",
			".********.",
			".**.*.*.*.",
			".*.*.*.**.",
			".**.*.*.*.",
			".*.*.*.**.",
			".**.*.*.*.",
			".*.*.*.**.",
			".********.",
			"..........",
		})
		assertEqBuffers(t, b, expected)
	}
}

func TestBuffer_FillWithMask(t *testing.T) {
	pattern := MustNewBufferFromStrings([]string{
		"*.",
		".*",
	})
	mask := MustNewBufferFromStrings([]string{
		"**********",
		"*........*",
		"*........*",
		"*........*",
		"*........*",
		"*........*",
		"*........*",
		"*........*",
		"*........*",
		"**********",
	})
	b := NewBuffer(10, 10)
	b.Fill(2, 2, WithPattern(pattern), WithMask(mask))

	expected := MustNewBufferFromStrings([]string{
		"..........",
		".*.*.*.*..",
		"..*.*.*.*.",
		".*.*.*.*..",
		"..*.*.*.*.",
		".*.*.*.*..",
		"..*.*.*.*.",
		".*.*.*.*..",
		"..*.*.*.*.",
		"..........",
	})
	assertEqBuffers(t, b, expected)
}

func TestBuffer_FillWithAllAreas(t *testing.T) {
	pattern := MustNewBufferFromStrings([]string{
		"*.",
		".*",
	})
	mask := MustNewBufferFromStrings([]string{
		"*****.....",
		"*****.....",
		"*****.....",
		"*****.....",
		"*****.....",
		".....*****",
		".....*****",
		".....*****",
		".....*****",
		".....*****",
	})
	b := NewBuffer(10, 10)
	b.Fill(0, 0, WithPattern(pattern), WithMask(mask), WithAllAreas())

	expected := MustNewBufferFromStrings([]string{
		"......*.*.",
		".....*.*.*",
		"......*.*.",
		".....*.*.*",
		"......*.*.",
		".*.*......",
		"*.*.*.....",
		".*.*......",
		"*.*.*.....",
		".*.*......",
	})
	assertEqBuffers(t, b, expected)
}
