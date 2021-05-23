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
			screen.Fini()
			return nil, err
		}
	}

	width := scrW * mode.Width()
	height := scrH * mode.Height()

	result := Tcg{
		mode:        mode,
		config:      config,
		scrW:        scrW,
		scrH:        scrH,
		Width:       width,
		Height:      height,
		TCellScreen: screen,
	}
	result.applyClip()

	return &result, nil
}

func (tg *Tcg) applyClip() {
	if tg.config.clip.width > 0 && tg.config.clip.height > 0 {
		tg.scrW = tg.config.clip.width
		tg.scrH = tg.config.clip.height
		tg.Width = tg.config.clip.width * tg.mode.Width()
		tg.Height = tg.config.clip.height * tg.mode.Height()
	}

	tg.Buf = NewBuffer(tg.Width, tg.Height)
}

// Show - update screen
func (tg *Tcg) Show() {
	tg.updateScreen()
	tg.TCellScreen.Show()
}

func (tg *Tcg) updateScreen() {
	chatMapping := pixelChars[tg.mode]
	blockW, blockH := tg.mode.Width(), tg.mode.Height()

	for x := 0; x < tg.scrW; x++ {
		for y := 0; y < tg.scrH; y++ {
			charIndex := tg.Buf.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
			tg.TCellScreen.SetContent(tg.config.clip.x+x, tg.config.clip.y+y, chatMapping[charIndex], nil, defaultStyle)
		}
	}
}

// RenderAsStrings - render buffer as slice of strings
func RenderAsStrings(buf Buffer, mode PixelsInChar) []string {
	chatMapping := pixelChars[mode]
	blockW, blockH := mode.Width(), mode.Height()

	var result []string

	width := buf.Width / blockW
	if buf.Width%blockW != 0 {
		width++
	}

	height := buf.Height / blockH
	if buf.Height%blockH != 0 {
		height++
	}

	for y := 0; y < height; y++ {
		line := ""
		for x := 0; x < width; x++ {
			charIndex := buf.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
			line += string(chatMapping[charIndex])
		}
		result = append(result, line)
	}

	return result
}

// Finish application
func (tg *Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// SetClip - set new clip of screen
func (tg *Tcg) SetClip(x, y, width, height int) error {
	if err := WithClip(x, y, width, height)(&tg.config); err != nil {
		return err
	}
	tg.applyClip()
	return nil
}

// SetClipCenter - set new clip in center of screen
func (tg *Tcg) SetClipCenter(width, height int) error {
	if err := WithClipCenter(width, height)(&tg.config); err != nil {
		return err
	}
	tg.applyClip()
	return nil
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
// x, y - is in screen character coordinates, not pixels
func (tg *Tcg) PrintStr(x, y int, str string) {
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, y, ch, nil, defaultStyle)
	}
}
