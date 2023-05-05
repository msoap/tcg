package sprite

import "github.com/msoap/tcg"

// Sprite - sprite object
type Sprite struct {
	Buf     tcg.Buffer // sprite image
	x, y    int        // position on buffer, can be negative
	width   int
	height  int
	mask    *tcg.Buffer // mask for sprite
	bg      tcg.Buffer  // saved
	isDrawn bool        // is sprite drawn on buffer
}

// New - get new sprite object from tcg.Buffer
func New(buf tcg.Buffer) *Sprite {
	return &Sprite{
		width:  buf.Width,
		height: buf.Height,
		Buf:    buf,
		bg:     tcg.NewBuffer(buf.Width, buf.Height),
	}
}

// WithMask - add mask to sprite
func (s *Sprite) WithMask(mask tcg.Buffer) *Sprite {
	s.mask = &mask
	return s
}

// Put - put sprite on buffer, save background under sprite, coordinates can be negative
func (s *Sprite) Put(buf tcg.Buffer, x, y int) {
	// restore previous background
	if s.isDrawn {
		buf.BitBlt(s.x, s.y, s.width, s.height, s.bg, 0, 0)
	}

	// copy background
	s.bg.BitBlt(0, 0, s.width, s.height, buf, x, y)

	// draw sprite
	var opts []tcg.BitBltOpt
	if s.mask != nil {
		opts = append(opts, tcg.BBMask(s.mask))
	}
	buf.BitBlt(x, y, s.width, s.height, s.Buf, 0, 0, opts...)

	s.isDrawn = true
	s.x = x
	s.y = y
}

// Withdraw - withdraw sprite from buffer
func (s *Sprite) Withdraw(buf tcg.Buffer) {
	if s.isDrawn {
		buf.BitBlt(s.x, s.y, s.width, s.height, s.bg, 0, 0)
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
	s.MoveAbs(buf, s.x+x, s.y+y)
}
