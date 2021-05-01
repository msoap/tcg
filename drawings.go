package tcg

import "math"

// HLine - draw horizontal line
func (b *Buffer) HLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		b.Set(x+i, y, color)
	}
}

// VLine - draw vertical line
func (b *Buffer) VLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		b.Set(x, y+i, color)
	}
}

// Rect - draw rectangle with 1px frame
func (b *Buffer) Rect(x, y int, width, height int, color int) {
	b.HLine(x, y, width, color)
	b.HLine(x, y+height-1, width, color)
	b.VLine(x, y, height, color)
	b.VLine(x+width-1, y, height, color)
}

// FillRect - draw filled rectangle
func (b *Buffer) FillRect(x, y int, width, height int, color int) {
	for i := 0; i < height; i++ {
		b.HLine(x, y+i, width, color)
	}
}

// Line - draw line
func (b *Buffer) Line(x1, y1, x2, y2 int, color int) {
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)

	if math.Abs(dx) > math.Abs(dy) {
		if x2 < x1 {
			x1, y1, x2, y2 = x2, y2, x1, y1 // swap
		}
		for x := x1; x <= x2; x++ {
			y := int(math.Round(float64(y1) + dy*float64(x-x1)/dx))
			b.Set(x, y, color)
		}
	} else {
		if y2 < y1 {
			x1, y1, x2, y2 = x2, y2, x1, y1 // swap
		}
		for y := y1; y <= y2; y++ {
			x := int(math.Round(float64(x1) + dx*float64(y-y1)/dy))
			b.Set(x, y, color)
		}
	}
}

// Fill an area with black color
func (b *Buffer) Fill(x, y int) {
	b.fillNBPixel(x, y)
}

// fill neighboring pixels, up/down/right/left
func (b *Buffer) fillNBPixel(x, y int) {
	if b.At(x, y) == Black {
		return
	}

	b.Set(x, y, Black)

	// up
	if y > 0 {
		b.fillNBPixel(x, y-1)
	}
	// down
	if y < b.Height-1 {
		b.fillNBPixel(x, y+1)
	}
	// left
	if x > 0 {
		b.fillNBPixel(x-1, y)
	}
	// right
	if x < b.Width-1 {
		b.fillNBPixel(x+1, y)
	}
}
