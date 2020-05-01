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

	Mode1x1 PixelsInChar = 1
	Mode1x2 PixelsInChar = 1 * 2
	Mode2x2 PixelsInChar = 2 * 2
	Mode2x3 PixelsInChar = 2 * 3
)

const hPixelRatio = 2

var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorDefault)

// Tcg - tcell graphics object
type Tcg struct {
	TCellScreen tcell.Screen
	buffer      Buffer
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

var (
	pixelChars1x1 = [...]rune{
		' ', // 0
		'█', // 1
	}

	pixelChars1x2 = [...]rune{
		' ', // 00
		'▄', // 01
		'▀', // 10
		'█', // 11
	}

	pixelChars2x2 = [...]rune{
		' ', // 0000
		'▗', // 0001
		'▖', // 0010
		'▄', // 0011
		'▝', // 0100
		'▐', // 0101
		'▞', // 0110
		'▟', // 0111
		'▘', // 1000
		'▚', // 1001
		'▌', // 1010
		'▙', // 1011
		'▀', // 1100
		'▜', // 1101
		'▛', // 1110
		'█', // 1111
	}

	pixelChars2x3 = [...]rune{
		' ', // 000000
		'🬞', // 000001
		'🬏', // 000010
		'🬭', // 000011
		'🬇', // 000100
		'🬦', // 000101
		'🬖', // 000110
		'🬵', // 000111

		'🬃', // 001000
		'🬢', // 001001
		'🬓', // 001010
		'🬱', // 001011
		'🬋', // 001100
		'🬩', // 001101
		'🬚', // 001110
		'🬹', // 001111

		'🬁', // 010000
		'🬠', // 010001
		'🬑', // 010010
		'🬯', // 010011
		'🬉', // 010100
		'▐', // 010101
		'🬘', // 010110
		'🬷', // 010111

		'🬅', // 011000
		'🬤', // 011001
		'🬔', // 011010
		'🬳', // 011011
		'🬍', // 011100
		'🬫', // 011101
		'🬜', // 011110
		'🬻', // 011111

		'🬀', // 100000
		'🬟', // 100001
		'🬐', // 100010
		'🬮', // 100011
		'🬈', // 100100
		'🬧', // 100101
		'🬗', // 100110
		'🬶', // 100111

		'🬄', // 101000
		'🬣', // 101001
		'▌', // 101010
		'🬲', // 101011
		'🬌', // 101100
		'🬪', // 101101
		'🬛', // 101110
		'🬺', // 101111

		'🬂', // 110000
		'🬡', // 110001
		'🬒', // 110010
		'🬰', // 110011
		'🬊', // 110100
		'🬨', // 110101
		'🬙', // 110110
		'🬸', // 110111

		'🬆', // 111000
		'🬥', // 111001
		'🬕', // 111010
		'🬴', // 111011
		'🬎', // 111100
		'🬬', // 111101
		'🬝', // 111110
		'█', // 111111
	}
)

// PutPixel - put pixel on the screen
func (tg *Tcg) PutPixel(x, y int, color int) {
	tg.buffer.PutPixel(x, y, color)

	//        x
	// y: 0: [0][1][2][3] [4][5][6][7]
	// y: 1: [0][1][2][3] [4][5][6][7]
	var index int
	scrY, remY := y/hPixelRatio, y%hPixelRatio
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
// string don't save in buffer!
func (tg *Tcg) PrintStr(x, y int, str string) {
	scrY := y / hPixelRatio
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, scrY, ch, nil, defaultStyle)
	}
}
