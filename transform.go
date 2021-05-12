package tcg

// Clear - fill whole buffer with White
func (b *Buffer) Clear() {
	for y := 0; y < len(b.buffer); y++ {
		for x := 0; x < len(b.buffer[y]); x++ {
			b.buffer[y][x] = 0
		}
	}
}

// Invert pixels in the buffer
func (b *Buffer) Invert() {
	for y := 0; y < len(b.buffer); y++ {
		for x := 0; x < len(b.buffer[y]); x++ {
			b.buffer[y][x] = ^b.buffer[y][x]
		}
	}
}

// BitBltAllSrc - copy whole buffer into this buffer
func (b *Buffer) BitBltAllSrc(x, y int, from Buffer) {
	if x == 0 && y == 0 {
		for i := 0; i < from.Height && i < b.Height; i++ {
			copy(b.buffer[i], from.buffer[i])
		}
		return
	}

	for i := 0; i < from.Height && i+y < b.Height; i++ {
		for j := 0; j < from.Width && j+x < b.Width; j++ {
			b.Set(j+x, i+y, from.At(j, i))
		}
	}
}

// BitBlt - copy part of buffer into this buffer
func (b *Buffer) BitBlt(xd, yd, width, height int, from Buffer, xs, ys int) {
	for i := 0; i+ys < from.Height && i < height && i+yd < b.Height; i++ {
		for j := 0; j+xs < from.Width && j < width && j+xd < b.Width; j++ {
			b.Set(j+xd, i+yd, from.At(j+xs, i+ys))
		}
	}
}

// Fill an area with black color
func (b *Buffer) Fill(x, y int) {
	b.fillNBPixel(x, y, 0)
}

// fill neighboring pixels, up/down/right/left, whatPrev - where did we come from?
func (b *Buffer) fillNBPixel(x, y int, whatPrev int) {
	if b.At(x, y) == Black {
		return
	}

	b.Set(x, y, Black)

	// up
	if whatPrev != 2 && y > 0 {
		b.fillNBPixel(x, y-1, 1)
	}
	// down
	if whatPrev != 1 && y < b.Height-1 {
		b.fillNBPixel(x, y+1, 2)
	}
	// left
	if whatPrev != 4 && x > 0 {
		b.fillNBPixel(x-1, y, 3)
	}
	// right
	if whatPrev != 3 && x < b.Width-1 {
		b.fillNBPixel(x+1, y, 4)
	}
}

// FlipH - horizontal flip image buffer
func (b *Buffer) FlipH() {
	for y := 0; y < b.Height/2; y++ {
		b.buffer[y], b.buffer[b.Height-y-1] = b.buffer[b.Height-y-1], b.buffer[y]
	}
}
