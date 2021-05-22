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
	config        tcgConfig
	scrW, scrH    int // screen width/height in characters
	Width, Height int // screen width/height in pixels
	TCellScreen   tcell.Screen
	Buf           Buffer
}

// New - get new object with tcell inside
func New(mode PixelsInChar, opts ...Opt) (*Tcg, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	scrW, scrH := screen.Size()
	config := tcgConfig{width: scrW, height: scrH}
	for _, optFn := range opts {
		if err := optFn(&config); err != nil {
			return nil, err
		}
	}

	width := scrW * mode.Width()
	height := scrH * mode.Height()
	// setup clip
	if config.clip.width > 0 && config.clip.height > 0 {
		scrW = config.clip.width
		scrH = config.clip.height
		width = config.clip.width * mode.Width()
		height = config.clip.height * mode.Height()
	}

	return &Tcg{
		mode:        mode,
		config:      config,
		scrW:        scrW,
		scrH:        scrH,
		Width:       width,
		Height:      height,
		TCellScreen: screen,
		Buf:         NewBuffer(width, height),
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
			tg.TCellScreen.SetContent(tg.config.clip.x+x, tg.config.clip.y+y, chatMapping[charIndex], nil, defaultStyle)
		}
	}
}

// Finish application
func (tg Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
// x, y - is in screen character coordinates, not pixels
func (tg *Tcg) PrintStr(x, y int, str string) {
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, y, ch, nil, defaultStyle)
	}
}
