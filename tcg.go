package tcg

import (
	"github.com/gdamore/tcell/v2"
)

// pixel colors
const (
	White = 0 // without pixel
	Black = 1 // it will be black on a terminal with light theme, and white on dark terminals
)

var defaultStyle = tcell.StyleDefault.Foreground(tcell.ColorDefault)

// Tcg - tcell graphics object
type Tcg struct {
	mode          PixelMode
	config        tcgConfig
	scrW, scrH    int          // screen or clip of screen width/height in characters
	Width, Height int          // screen or clip of screen width/height in pixels
	TCellScreen   tcell.Screen // tcell object for keyboard interactions, or low level interactions with terminal screen
	Buf           Buffer       // buffer presents current screen
	style         tcell.Style
}

// New - get new object with tcell inside
func New(mode PixelMode, opts ...Opt) (*Tcg, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	scrW, scrH := screen.Size()
	config := tcgConfig{width: scrW, height: scrH, style: defaultStyle}
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
		style:       config.style,
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
// it updates the screen with the current buffer content and can erase any text printed before by PrintStr or PrintStrStyle.
// finalisers - should be used to apply additional changes to the screen after pixels are updated, for example, to print some text on the screen.
func (tg *Tcg) Show(finalisers ...func(tg *Tcg)) {
	tg.updateScreen()
	for _, finaliser := range finalisers {
		finaliser(tg)
	}
	tg.TCellScreen.Show()
}

func (tg *Tcg) updateScreen() {
	blockW, blockH := tg.mode.Width(), tg.mode.Height()

	for x := 0; x < tg.scrW; x++ {
		for y := 0; y < tg.scrH; y++ {
			charIndex := tg.Buf.getPixelsBlock(x*blockW, y*blockH, blockW, blockH)
			tg.TCellScreen.SetContent(tg.config.clip.x+x, tg.config.clip.y+y, tg.mode.charMapping[charIndex], nil, tg.style)
		}
	}
}

// Finish application
func (tg *Tcg) Finish() {
	tg.TCellScreen.Fini()
}

// SetClip - set new clip of screen, this allows to limit the area of the screen that will be used for drawing.
// x, y - is in screen character coordinates, not pixels.
// width, height - is in screen character coordinates, not pixels.
// It can be run any time, but clearing internal buffer screen.
func (tg *Tcg) SetClip(x, y, width, height int) error {
	if err := WithClip(x, y, width, height)(&tg.config); err != nil {
		return err
	}
	tg.applyClip()
	return nil
}

// SetClipCenter - set new clip in the center of screen
func (tg *Tcg) SetClipCenter(width, height int) error {
	if err := WithClipCenter(width, height)(&tg.config); err != nil {
		return err
	}
	tg.applyClip()
	return nil
}

// PrintStr - print string on screen, with white on black style
// string don't save in the buffer of pixels!
// x, y - is in screen character coordinates, not pixels.
// Also x/y coordinates is not use Clip of the screen, it's always absolute.
func (tg *Tcg) PrintStr(x, y int, str string) {
	tg.PrintStrStyle(x, y, str, tg.style)
}

// PrintStrStyle - print string on screen
// see the PrintStr about restrictions
func (tg *Tcg) PrintStrStyle(x, y int, str string, style tcell.Style) {
	for i, ch := range []rune(str) {
		tg.TCellScreen.SetContent(x+i, y, ch, nil, style)
	}
}

// ScreenSize - returns terminal screen size in chars (width, height)
func (tg *Tcg) ScreenSize() (int, int) {
	return tg.TCellScreen.Size()
}
