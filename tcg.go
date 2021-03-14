//go:generate minimock -i github.com/gdamore/tcell/v2.Screen -o .

package tcg

import (
	"github.com/gdamore/tcell/v2"
)

// PixelsInChar - a type representing the graphics mode
type PixelsInChar int

// pixel colors
const (
	White = 0
	Black = 1
)

// graphics modes
const (
	Mode1x1 PixelsInChar = iota
	Mode1x2
	Mode2x2
	Mode2x3
)

var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorDefault)

// Tcg - tcell graphics object
type Tcg struct {
	mode        PixelsInChar
	scrW, scrH  int // screen width/height in characters
	TCellScreen tcell.Screen
	buffer      Buffer
}

// Width - returns the width in pixels of one character in the text console
func (pic PixelsInChar) Width() int {
	switch pic {
	case Mode1x1, Mode1x2:
		return 1
	case Mode2x2, Mode2x3:
		return 2
	default:
		return 0
	}
}

// Height - returns the height in pixels of one character in the text console
func (pic PixelsInChar) Height() int {
	switch pic {
	case Mode1x1:
		return 1
	case Mode1x2, Mode2x2:
		return 2
	case Mode2x3:
		return 3
	default:
		return 0
	}
}

// New - get new object with tcell inside
func New(mode PixelsInChar) (Tcg, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return Tcg{}, err
	}

	if err := screen.Init(); err != nil {
		return Tcg{}, err
	}
	w, h := screen.Size()

	return Tcg{
		mode:        mode,
		scrW:        w,
		scrH:        h,
		TCellScreen: screen,
		buffer:      NewBuffer(w*mode.Width(), h*mode.Height()),
	}, nil
}

// Show - update screen
func (tg Tcg) Show() {
	tg.updateScreen()
	tg.TCellScreen.Show()
}

func (tg Tcg) updateScreen() {
	chatMapping := pixelChars[tg.mode]
	blockW, blockH := tg.mode.Width(), tg.mode.Height()

	for x := 0; x < tg.scrW; x++ {
		for y := 0; y < tg.scrH; y++ {
			charIndex := tg.buffer.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
			tg.TCellScreen.SetContent(x, y, chatMapping[charIndex], nil, defaultStyle)
		}
	}
}

// Finish application
func (tg Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// PutPixel - put pixel on the screen
func (tg *Tcg) PutPixel(x, y int, color int) {
	tg.buffer.PutPixel(x, y, color)
}

// GetPixel - get pixel from the screen
func (tg *Tcg) GetPixel(x, y int) int {
	return tg.buffer.GetPixel(x, y)
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
func (tg *Tcg) PrintStr(x, y int, str string) {
	scrY := y / tg.mode.Height()
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, scrY, ch, nil, defaultStyle)
	}
}
