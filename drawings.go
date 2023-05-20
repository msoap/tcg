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

// Line - draw line using the Bresenham's algorithm
func (b *Buffer) Line(x1, y1, x2, y2 int, color int) {
	dx := abs(x2 - x1)
	dy := -abs(y2 - y1)

	sx := sgn(x2 - x1)
	sy := sgn(y2 - y1)

	e := dx + dy
	x0, y0 := x1, y1

	for {
		b.Set(x0, y0, color)

		if x0 == x2 && y0 == y2 {
			break
		}
		e2 := 2 * e

		if e2 >= dy {
			if x0 == x2 {
				break
			}
			e = e + dy
			x0 = x0 + sx
		}

		if e2 <= dx {
			if y0 == y2 {
				break
			}
			e = e + dx
			y0 = y0 + sy
		}
	}
}

// Circle - draw a circle using the Midpoint Circle Algorithm
func (b *Buffer) Circle(x, y, r int, color int) {
	if r < 0 {
		return
	}

	x1, y1, err := -r, 0, 2-2*r
	for {
		b.Set(x-x1, y+y1, color)
		b.Set(x-y1, y-x1, color)
		b.Set(x+x1, y-y1, color)
		b.Set(x+y1, y+x1, color)
		rr := err
		if rr > x1 {
			x1++
			err += x1*2 + 1
		}
		if rr <= y1 {
			y1++
			err += y1*2 + 1
		}
		if x1 >= 0 {
			break
		}
	}
}

// circleUsingArc - draw circle via arc method
func (b *Buffer) circleUsingArc(x, y, r int, color int) {
	b.Arc(x, y, float64(r), 0, 360, color)
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
