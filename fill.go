package tcg

// Fill an area with black color
func (b *Buffer) Fill(x, y int, opts ...FillOpt) {
	fo := fillOptions{}
	for _, fn := range opts {
		fn(&fo)
	}

	if fo.mask != nil {
		fo.checkBuf = fo.mask.Clone()
	} else if fo.pattern != nil || fo.allAreas {
		fo.checkBuf = b.Clone()
	}

	if fo.allAreas {
		b.fillAllPixel(x, y, x, y, fo)
	} else {
		b.fillNBPixel(x, y, x, y, 0, fo)
	}
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

// getFillColor - get color for fill, returns color and flag whether to put a pixel
func (b *Buffer) getFillColor(x, y, xs, ys int, fo fillOptions) (int, bool) {
	var color int

	if fo.pattern != nil {
		if fo.checkBuf.At(x, y) == Black {
			return 0, false
		}
		fo.checkBuf.Set(x, y, Black)
		color = fo.pattern.At(abs(x-xs)%fo.pattern.Width, abs(y-ys)%fo.pattern.Height)
	} else {
		if b.At(x, y) == Black {
			return 0, false
		}
		color = Black
	}

	return color, true
}

// fill neighboring pixels, up/down/right/left, whatPrev - where did we come from?
func (b *Buffer) fillNBPixel(x, y, xs, ys int, whatPrev int, fo fillOptions) {
	color, ok := b.getFillColor(x, y, xs, ys, fo)
	if !ok {
		return
	}

	if fo.pattern != nil {
		fo.checkBuf.Set(x, y, Black)
	}

	b.Set(x, y, color)

	// up
	if whatPrev != 2 && y > 0 {
		b.fillNBPixel(x, y-1, xs, ys, 1, fo)
	}
	// down
	if whatPrev != 1 && y < b.Height-1 {
		b.fillNBPixel(x, y+1, xs, ys, 2, fo)
	}
	// left
	if whatPrev != 4 && x > 0 {
		b.fillNBPixel(x-1, y, xs, ys, 3, fo)
	}
	// right
	if whatPrev != 3 && x < b.Width-1 {
		b.fillNBPixel(x+1, y, xs, ys, 4, fo)
	}
}

// fillAllPixel - fill all white pixels
func (b *Buffer) fillAllPixel(x, y, xs, ys int, fo fillOptions) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if color, ok := b.getFillColor(x, y, xs, ys, fo); ok {
				b.Set(x, y, color)
			}
		}
	}
}
