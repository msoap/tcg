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

// Circle - draw circle
func (b *Buffer) Circle(x, y int, r float64, color int) {
	b.Arc(x, y, r, 0, 360, color)
}

// Arc - draw circle arc, from and to: 0 .. 360
func (b *Buffer) Arc(x, y int, r float64, from, to float64, color int) {
	step := 2 * math.Pi / (r * 8) // in radians
	fromR := from / 180 * math.Pi
	toR := to / 180 * math.Pi
	for θ := fromR; θ < toR; θ += step {
		nx := x + int(math.Round(r*math.Cos(θ)))
		ny := y - int(math.Round(r*math.Sin(θ)))
		b.Set(nx, ny, color)
	}
}
