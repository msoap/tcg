//go:generate minimock -i github.com/gdamore/tcell.Screen -o .

package tcg

import (
	"github.com/gdamore/tcell"
)

type PixelsInChar int

// pixel colors
const (
	White = 0
	Black = 1
)

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
	TCellScreen tcell.Screen
	buffer      Buffer
}

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
	h *= mode.Height()

	return Tcg{
		mode:        mode,
		TCellScreen: screen,
		buffer:      NewBuffer(w, h),
	}, nil
}

// Show - update screen
func (tg Tcg) Show() {
	tg.TCellScreen.Show()
}

// Finish application
func (tg Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// PutPixel - put pixel on the screen
func (tg *Tcg) PutPixel(x, y int, color int) {
	tg.buffer.PutPixel(x, y, color)

	//        x
	// y: 0: [0][1][2][3] [4][5][6][7]
	// y: 1: [0][1][2][3] [4][5][6][7]
	var index int
	scrY, remY := y/tg.mode.Height(), y%tg.mode.Height()
	if remY == 0 {
		pairedPx := tg.GetPixel(x, y+1)
		index = color<<1 | pairedPx
	} else {
		pairedPx := tg.GetPixel(x, y-1)
		index = pairedPx<<1 | color
	}

	tg.TCellScreen.SetContent(x, scrY, pixelChars1x2[index], nil, defaultStyle)
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
