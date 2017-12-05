package tcg

import (
	"github.com/gdamore/tcell"
)

// pixel colors
const (
	White = 0
	Black = 1
)

const hPixelRatio = 2

var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorDefault)

// Tcg - tcell graphics object
type Tcg struct {
	TCellScreen tcell.Screen
	buffer      Buffer
}

// New - get new object with tcell inside
func New() (Tcg, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return Tcg{}, err
	}

	if err := screen.Init(); err != nil {
		return Tcg{}, err
	}
	w, h := screen.Size()
	h *= hPixelRatio // each character cell contains two pixels

	return Tcg{
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

var pixelChars = [...]rune{' ', '▄', '▀', '█'}

// PutPixel - put pixel on the screen
func (tg *Tcg) PutPixel(x, y int, color int) {
	tg.buffer.PutPixel(x, y, color)

	//        x
	// y: 0: [0][1][2][3] [4][5][6][7]
	// y: 1: [0][1][2][3] [4][5][6][7]
	scrY, remY, index := y/hPixelRatio, y%hPixelRatio, 0
	if remY == 0 {
		pairedPx := tg.GetPixel(x, y+1)
		index = color<<1 | pairedPx
	} else {
		pairedPx := tg.GetPixel(x, y-1)
		index = pairedPx<<1 | color
	}

	tg.TCellScreen.SetContent(x, scrY, pixelChars[index], nil, defaultStyle)
}

// GetPixel - get pixel from the screen
func (tg *Tcg) GetPixel(x, y int) int {
	return tg.buffer.GetPixel(x, y)
}

// PrintStr - print string on screen, with white on black style
// string don't save in buffer!
func (tg *Tcg) PrintStr(x, y int, str string) {
	scrY := int(y / hPixelRatio)
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, scrY, ch, nil, defaultStyle)
	}
}
