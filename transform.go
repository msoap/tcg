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

// HFlip - horizontal flip image buffer
func (b *Buffer) HFlip() {
	for y := 0; y < b.Height/2; y++ {
		b.buffer[y], b.buffer[b.Height-y-1] = b.buffer[b.Height-y-1], b.buffer[y]
	}
}

// VFlip - vertical flip image buffer
func (b *Buffer) VFlip() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width/2; x++ {
			leftColor, rightColor := b.At(x, y), b.At(b.Width-x-1, y)
			b.Set(x, y, rightColor)
			b.Set(b.Width-x-1, y, leftColor)
		}
	}
}

// VScroll - vertical scroll image buffer by cnt pixels, cnt > 0 - scroll down, cnt < 0 - up
func (b *Buffer) VScroll(cnt int) {
	zeroLine := make([]byte, widthInBytes(b.Width))

	if cnt > 0 {
		for y := b.Height - 1; y > cnt-1; y-- {
			copy(b.buffer[y], b.buffer[y-cnt])
		}
		// clear rest
		for y := 0; y < cnt; y++ {
			copy(b.buffer[y], zeroLine)
		}
	} else if cnt < 0 {
		for y := 0; y < b.Height+cnt; y++ {
			copy(b.buffer[y], b.buffer[y-cnt])
		}
		// clear rest
		for y := b.Height + cnt; y < b.Height; y++ {
			copy(b.buffer[y], zeroLine)
		}
	}
}

// HScroll - horizontal scroll image buffer by cnt pixels, cnt > 0 - scroll right, cnt < 0 - left
func (b *Buffer) HScroll(cnt int) {
	if cnt > 0 {
		for y := 0; y < b.Height; y++ {
			for x := b.Width - 1; x > cnt-1; x-- {
				b.Set(x, y, b.At(x-cnt, y))
			}
		}
		// clear rest
		for x := 0; x < cnt; x++ {
			b.VLine(x, 0, b.Height, White)
		}
	} else if cnt < 0 {
		for y := 0; y < b.Height; y++ {
			for x := 0; x < b.Width+cnt; x++ {
				b.Set(x, y, b.At(x-cnt, y))
			}
		}
		// clear rest
		for x := b.Width + cnt; x < b.Width; x++ {
			b.VLine(x, 0, b.Height, White)
		}
	}
}
