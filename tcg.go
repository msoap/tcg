//go:generate minimock -i github.com/gdamore/tcell/v2.Screen -o .

package tcg

import (
	"github.com/gdamore/tcell/v2"
)

// pixel colors
const (
	White = 0
	Black = 1
)

var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorDefault)

// Tcg - tcell graphics object
type Tcg struct {
	mode          PixelsInChar
	scrW, scrH    int // screen width/height in characters
	Width, Height int // screen width/height in pixels
	TCellScreen   tcell.Screen
	Buf           Buffer
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
		Buf:         NewBuffer(w*mode.Width(), h*mode.Height()),
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
			charIndex := tg.Buf.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
			tg.TCellScreen.SetContent(x, y, chatMapping[charIndex], nil, defaultStyle)
		}
	}
}

// Finish application
func (tg Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
func (tg *Tcg) PrintStr(x, y int, str string) {
	scrY := y / tg.mode.Height()
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, scrY, ch, nil, defaultStyle)
	}
}
