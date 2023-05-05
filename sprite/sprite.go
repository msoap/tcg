/*
Package sprite - sprite object for drawing/moving on tcg.Buffer
*/
package sprite

import (
	"github.com/msoap/tcg"
)

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

// Put - put sprite on buffer, save background under sprite, change state of sprite to drawn
func (s *Sprite) Put(buf tcg.Buffer) *Sprite {
	s.draw(buf)
	s.isDrawn = true

	return s
}

// draw - draw sprite on buffer, save background under sprite
func (s *Sprite) draw(buf tcg.Buffer) {
	// copy background
	s.bg.BitBlt(0, 0, s.width, s.height, buf, s.x, s.y)

	// draw sprite
	var opts []tcg.BitBltOpt
	if s.mask != nil {
		opts = append(opts, tcg.BBMask(s.mask))
	}
	buf.BitBlt(s.x, s.y, s.width, s.height, s.Buf, 0, 0, opts...)
}

// Withdraw - withdraw sprite from buffer
func (s *Sprite) Withdraw(buf tcg.Buffer) *Sprite {
	if s.isDrawn {
		s.clear(buf)
		s.isDrawn = false
	}

	return s
}

// clear sprite on buffer
func (s *Sprite) clear(buf tcg.Buffer) {
	buf.BitBlt(s.x, s.y, s.width, s.height, s.bg, 0, 0)
}

// MoveAbs - move sprite on buffer to absolute position, coordinates can be negative
func (s *Sprite) MoveAbs(buf tcg.Buffer, x, y int) *Sprite {
	if s.isDrawn {
		s.clear(buf)
	}

	s.x = x
	s.y = y

	if s.isDrawn {
		s.draw(buf)
	}

	return s
}

// Move - move sprite on buffer to relative position
func (s *Sprite) Move(buf tcg.Buffer, x, y int) *Sprite {
	return s.MoveAbs(buf, s.x+x, s.y+y)
}
