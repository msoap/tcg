package tcg

// BitBltAll - copy whole buffer into this buffer
func (b *Buffer) BitBltAll(x, y int, from Buffer, opts ...BitBltOpt) {
	if x == 0 && y == 0 && len(opts) == 0 {
		for i := 0; i < from.Height && i < b.Height; i++ {
			copy(b.buffer[i], from.buffer[i])
		}
		return
	}

	b.BitBlt(x, y, from.Width, from.Height, from, 0, 0, opts...)
}

// BitBlt - copy part of buffer into this buffer
// xd, yd - destination coordinates
// xs, ys - source coordinates
func (b *Buffer) BitBlt(xd, yd, width, height int, from Buffer, xs, ys int, opts ...BitBltOpt) {
	if len(opts) == 0 {
		for i := 0; i+ys < from.Height && i < height && i+yd < b.Height; i++ {
			for j := 0; j+xs < from.Width && j < width && j+xd < b.Width; j++ {
				b.Set(j+xd, i+yd, from.At(j+xs, i+ys))
			}
		}
		return
	}

	cfg := bitBltOptions{}
	for _, opt := range opts {
		opt(&cfg)
	}

	for i := 0; i+ys < from.Height && i < height && i+yd < b.Height; i++ {
		for j := 0; j+xs < from.Width && j < width && j+xd < b.Width; j++ {
			newColor := from.At(j+xs, i+ys)

			if cfg.transparent && newColor == White {
				continue
			}
			if cfg.mask != nil && cfg.mask.At(j+xs, i+ys) == White {
				continue
			}

			srcColor := b.At(j+xd, i+yd)
			for _, fn := range cfg.operations {
				newColor = fn(srcColor, newColor)
			}

			b.Set(j+xd, i+yd, newColor)
		}
	}
}
