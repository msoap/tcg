package sprite

import "github.com/msoap/tcg"

// Sprite - sprite object
type Sprite struct {
	X, Y    int
	Width   int
	Height  int
	Buf     tcg.Buffer  // sprite image
	Mask    *tcg.Buffer // mask for sprite
	bg      tcg.Buffer  // saved
	isDrawn bool        // is sprite drawn on buffer
}

// New - get new sprite object
func New(width, height int) *Sprite {
	return &Sprite{
		Width:  width,
		Height: height,
		Buf:    tcg.NewBuffer(width, height),
		bg:     tcg.NewBuffer(width, height),
	}
}

// NewWithBuffer - get new sprite object from tcg.Buffer
func NewWithBuffer(buf tcg.Buffer) *Sprite {
	return &Sprite{
		Width:  buf.Width,
		Height: buf.Height,
		Buf:    buf,
		bg:     tcg.NewBuffer(buf.Width, buf.Height),
	}
}

// Put - put sprite on buffer
func (s *Sprite) Put(buf tcg.Buffer, x, y int) {
	// restore previous background
	if s.isDrawn {
		buf.BitBlt(s.X, s.Y, s.Width, s.Height, s.bg, 0, 0)
	}

	// copy background
	s.bg.BitBlt(0, 0, s.Width, s.Height, buf, x, y)

	// draw sprite
	var opts []tcg.BitBltOpt
	if s.Mask != nil {
		opts = append(opts, tcg.BBMask(s.Mask))
	}
	buf.BitBlt(x, y, s.Width, s.Height, s.Buf, 0, 0, opts...)

	s.isDrawn = true
	s.X = x
	s.Y = y
}

// Withdraw - withdraw sprite from buffer
func (s *Sprite) Withdraw(buf tcg.Buffer) {
	if s.isDrawn {
		buf.BitBlt(s.X, s.Y, s.Width, s.Height, s.bg, 0, 0)
		s.isDrawn = false
	}
}

// MoveAbs - move sprite on buffer to absolute position
func (s *Sprite) MoveAbs(buf tcg.Buffer, x, y int) {
	s.Withdraw(buf)
	s.Put(buf, x, y)
}

// Move - move sprite on buffer to relative position
func (s *Sprite) Move(buf tcg.Buffer, x, y int) {
	s.MoveAbs(buf, s.X+x, s.Y+y)
}
