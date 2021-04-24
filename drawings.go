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

// Box - draw box
func (b *Buffer) Box(x, y int, width, height int, color int) {
	b.HLine(x, y, width, color)
	b.HLine(x, y+height-1, width, color)
	b.VLine(x, y, height, color)
	b.VLine(x+width-1, y, height, color)
}

// FillBox - draw filled box
func (b *Buffer) FillBox(x, y int, width, height int, color int) {
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
