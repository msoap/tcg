//go:generate minimock -i github.com/gdamore/tcell/v2.Screen -o .

package tcg

import (
	"errors"

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
	mode          PixelsInChar
	scrW, scrH    int // screen width/height in characters
	Width, Height int // screen width/height in pixels
	TCellScreen   tcell.Screen
	Buffer        Buffer
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

// String represent
func (pic PixelsInChar) String() string {
	switch pic {
	case Mode1x1:
		return "1x1"
	case Mode1x2:
		return "1x2"
	case Mode2x2:
		return "2x2"
	case Mode2x3:
		return "2x3"
	default:
		return "-"
	}
}

// Set from string (for use with flag.Var())
func (pic *PixelsInChar) Set(in string) error {
	switch in {
	case "1x1":
		*pic = Mode1x1
	case "1x2":
		*pic = Mode1x2
	case "2x2":
		*pic = Mode2x2
	case "2x3":
		*pic = Mode2x3
	default:
		return errors.New("not valid screen mode")
	}

	return nil
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
		Width:       w * mode.Width(),
		Height:      h * mode.Height(),
		TCellScreen: screen,
		Buffer:      NewBuffer(w*mode.Width(), h*mode.Height()),
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
			charIndex := tg.Buffer.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
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
	tg.Buffer.PutPixel(x, y, color)
}

// GetPixel - get pixel from the screen
func (tg *Tcg) GetPixel(x, y int) int {
	return tg.Buffer.GetPixel(x, y)
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
func (tg *Tcg) PrintStr(x, y int, str string) {
	scrY := y / tg.mode.Height()
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, scrY, ch, nil, defaultStyle)
	}
}
