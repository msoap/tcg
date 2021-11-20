package tcg

import (
	"testing"
)

func BenchmarkBuffer_BitBlt_OldSimple(b *testing.B) {
	buf := getTestBuffer()

	xd, yd, width, height, xs, ys := 0, 0, 10, 10, 10, 10
	for i := 0; i < b.N; i++ {
		// buf.BitBlt(0, 0, 10, 10, buf, 10, 10)
		for i := 0; i+ys < buf.Height && i < height && i+yd < buf.Height; i++ {
			for j := 0; j+xs < buf.Width && j < width && j+xd < buf.Width; j++ {
				buf.Set(j+xd, i+yd, buf.At(j+xs, i+ys))
			}
		}
	}
}

func getTestBuffer() Buffer {
	return MustNewBufferFromStrings([]string{
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
		".*.*.*.*.*.*.*.*.*.*",
		"*.*.*.*.*.*.*.*.*.*.",
	})
}
