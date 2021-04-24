/*
Implement Turtle drawing
*/
package turtle

import (
	"github.com/msoap/tcg"
)

// Turtle object
type Turtle struct {
	buf   *tcg.Buffer
	x, y  int  // current position
	color int  // current color
	pen   bool // now we are drawing by default
}

// New turtle object
func New(buf *tcg.Buffer) *Turtle {
	return &Turtle{buf: buf, color: tcg.Black, pen: true}
}

// Set - set pixel at current position (script: "S")
func (t *Turtle) Set() *Turtle {
	t.buf.Set(t.x, t.y, t.color)
	return t
}

// SetColor - set current color (script: "C1", "C0" for Black and White)
func (t *Turtle) SetColor(color int) *Turtle {
	t.color = color
	return t
}

// Raise pen and don't draw further (script: "N")
func (t *Turtle) Raise() *Turtle {
	t.pen = false
	return t
}

// Put pen and draw further (script: "Y")
func (t *Turtle) Put() *Turtle {
	t.pen = true
	return t
}

// Up - goto up and draw, or not (script: "U3")
func (t *Turtle) Up(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.y > 0 {
			t.y--
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// Down - goto down and draw, or not (script: "D3")
func (t *Turtle) Down(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.y < t.buf.Width-1 {
			t.y++
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// Right - goto right and draw, or not (script: "R3")
func (t *Turtle) Right(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x < t.buf.Height-1 {
			t.x++
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// Left - goto left and draw, or not (script: "L3")
func (t *Turtle) Left(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x > 0 {
			t.x--
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// UpRight - goto up and right and draw (script: "UR3")
func (t *Turtle) UpRight(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x < t.buf.Height-1 {
			t.x++
		}
		if t.y > 0 {
			t.y--
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// DownRight - goto down and right and draw (script: "DR3")
func (t *Turtle) DownRight(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x < t.buf.Height-1 {
			t.x++
		}
		if t.y < t.buf.Width-1 {
			t.y++
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// UpLeft - goto up and left and draw (script: "UL3")
func (t *Turtle) UpLeft(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x > 0 {
			t.x--
		}
		if t.y > 0 {
			t.y--
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// DownLeft - goto down and left and draw (script: "DL3")
func (t *Turtle) DownLeft(cnt int) *Turtle {
	for i := 0; i < cnt; i++ {
		if t.x > 0 {
			t.x--
		}
		if t.y < t.buf.Width-1 {
			t.y++
		}
		if t.pen {
			t.buf.Set(t.x, t.y, t.color)
		}
	}
	return t
}

// Goto to relative point without draw (script: "G3,-1")
func (t *Turtle) GoTo(x, y int) *Turtle {
	t.x += x
	t.y += y
	return t
}

// Goto to absulute point without draw (script: "GA2,1")
func (t *Turtle) GoToAbs(x, y int) *Turtle {
	t.x = x
	t.y = y
	return t
}

// LineTo - draw line to relative point (script: "LT-2,1")
func (t *Turtle) LineTo(x, y int) *Turtle {
	t.buf.Line(t.x, t.y, t.x+x, t.y+y, t.color)
	t.x += x
	t.y += y
	return t
}
