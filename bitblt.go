package tcg

// BitBltAll - copy whole buffer into this buffer
func (b *Buffer) BitBltAll(x, y int, from Buffer) {
	if x == 0 && y == 0 {
		for i := 0; i < from.Height && i < b.Height; i++ {
			copy(b.buffer[i], from.buffer[i])
		}
		return
	}

	b.BitBlt(x, y, from.Width, from.Height, from, 0, 0)
}

// BitBlt - copy part of buffer into this buffer
func (b *Buffer) BitBlt(xd, yd, width, height int, from Buffer, xs, ys int) {
	for i := 0; i+ys < from.Height && i < height && i+yd < b.Height; i++ {
		for j := 0; j+xs < from.Width && j < width && j+xd < b.Width; j++ {
			b.Set(j+xd, i+yd, from.At(j+xs, i+ys))
		}
	}
}
