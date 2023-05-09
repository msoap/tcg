package main

import (
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
	"github.com/msoap/tcg/sprite"
)

type (
	cmds int
	mode int
	game struct {
		tg               *tcg.Tcg
		mode             mode // play or pause/edit
		generation       int  // current generation number
		cursorX, cursorY int  // cursor position for edit/pen mode
		pen              bool // continue drawing with arrows keys
		showCursor       bool // show cursor in edit mode
		curs             *sprite.Sprite
		scrH             int        // screen height in characters
		infMap           bool       // is map infinite?
		history          *list.List // history of last generations
	}
)

const (
	defaultDelay          = time.Millisecond * 100
	defaultInitFillFactor = 0.2
	historySize           = 1000
	maxFPS                = 9999
	curHalf               = 3

	modePlay mode = iota
	modePause

	cmdExit         cmds = iota // exit
	cmdPause                    // pause/play
	cmdNext                     // next step
	cmdPrev                     // prev step
	cmdPixel                    // toggle one pixel in current position
	cmdPen                      // toggle pen mode, when pen mode is on, you can draw with arrows keys
	cmdToggleCursor             // toggle showing cursor in edit mode
	cmdWipe                     // clear screen
	cmdUp                       // move cursor up
	cmdDown                     // -/-
	cmdLeft                     // -/-
	cmdRight                    // -/-
	cmdScreenshot               // save screenshot
)

var (
	cursorImage = tcg.MustNewBufferFromStrings([]string{
		"...*...",
		"...*...",
		".......",
		"**...**",
		".......",
		"...*...",
		"...*...",
	})
	cursorMask = tcg.MustNewBufferFromStrings([]string{
		"..***..",
		"..***..",
		"**...**",
		"**...**",
		"**...**",
		"..***..",
		"..***..",
	})
)

func main() {
	delay := flag.Duration("delay", defaultDelay, "delay between steps")
	size := flag.String("size", "", "screen size in chars, in 'width x height' format, example: '80x25'")
	colorName := flag.String("color", "", "redefine color, it can be: 'yellow', 'red' or like '#ffaa11'")
	fillFactor := flag.Float64("fill", defaultInitFillFactor, "how much to fill the area initially")
	inFileName := flag.String("in", "", "load map from image file (*.png)")
	screenshotName := flag.String("out", "game_of_life.png", "save map as screenshot to file")
	infMap := flag.Bool("inf", false, "infinite map (wrap around edges)")
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille")
	wait := flag.Bool("wait", false, "wait for start")
	flag.Parse()

	var (
		width, height int
		err           error
	)
	if *size != "" {
		width, height, err = tcg.ParseSizeString(*size)
		if err != nil {
			log.Fatal(err)
		}
	}

	var initImg image.Image
	if *inFileName != "" {
		if initImg, err = loadImage(*inFileName); err != nil {
			log.Fatal(err)
		}
	}

	opts := []tcg.Opt{}
	if *colorName != "" {
		opts = append(opts, tcg.WithColor(*colorName))
	}

	tg, err := tcg.New(mode, opts...)
	if err != nil {
		log.Fatal(err)
	}

	if tg.TCellScreen.HasMouse() {
		tg.TCellScreen.EnableMouse(tcell.MouseMotionEvents)
	}

	_, scrH := tg.ScreenSize()

	pattern := tcg.MustNewBufferFromStrings([]string{
		" *",
		"* ",
	})
	tg.Buf.Fill(0, 0, tcg.WithPattern(pattern))
	tg.Buf.Rect(0, 0, tg.Width, tg.Height, tcg.Black) // coordinates in pixels
	tg.Show()

	if width == 0 {
		width, height = tg.ScreenSize()
		width -= 6
		height -= 4
	}

	if err := tg.SetClipCenter(width, height); err != nil {
		tg.Finish()
		log.Fatal(err)
	}

	tg.Buf.Rect(0, 0, tg.Width, tg.Height, tcg.Black) // coordinates in pixels
	tg.PrintStr(4, 1, " Game of Life ")               // coordinates in chars, not pixels
	tg.PrintStr(24, scrH-1, `| <q> - Quit | <p> - Pause | <l>/<h> - Next/Prev step | <Space>/<a> pixel/pen | <c> show cursor | <s> - Screenshot `)
	tg.Show()

	if err := tg.SetClipCenter(width-2, height-2); err != nil {
		tg.Finish()
		log.Fatal(err)
	}

	curMode := modePlay
	if *wait {
		curMode = modePause
	}
	game := newGame(tg, curMode)
	game.infMap = *infMap
	if initImg != nil {
		game.initFromImage(initImg)
	} else {
		game.initRandom(*fillFactor)
	}

	ticker := time.Tick(*delay)
	command := getCommand(tg)

LOOP:
	for {
		select {
		case <-ticker:
			if game.mode == modePlay {
				game.nextStep()
			}
		case cmd := <-command:
			switch cmd {
			case cmdExit:
				break LOOP
			case cmdPause:
				game.doPause()
			case cmdNext:
				if game.mode != modePlay {
					game.nextStep()
				}
			case cmdPrev:
				if game.mode != modePlay {
					game.prevStep()
				}
			case cmdPixel:
				if game.mode != modePlay {
					game.togglePixel()
				}
			case cmdLeft:
				if game.mode != modePlay {
					game.moveCursor(-1, 0)
				}
			case cmdRight:
				if game.mode != modePlay {
					game.moveCursor(1, 0)
				}
			case cmdUp:
				if game.mode != modePlay {
					game.moveCursor(0, -1)
				}
			case cmdDown:
				if game.mode != modePlay {
					game.moveCursor(0, 1)
				}
			case cmdPen:
				if game.mode != modePlay {
					game.pen = !game.pen
				}
			case cmdToggleCursor:
				if game.mode != modePlay {
					game.toggleCursor()
				}
			case cmdWipe:
				if game.mode != modePlay {
					game.tg.Buf.Clear()
					game.tg.Show()
				}
			case cmdScreenshot:
				if err := saveScreenshot(*screenshotName, tg.Buf); err != nil {
					tg.PrintStr(0, 0, fmt.Sprintf("save: %s", err))
					tg.Show()
				}
			}
		}
	}

	tg.Finish()
}

func newGame(tg *tcg.Tcg, mode mode) *game {
	_, scrH := tg.ScreenSize()

	cursorSprite := sprite.New(cursorImage).WithMask(cursorMask)

	return &game{
		mode:    mode,
		tg:      tg,
		curs:    cursorSprite,
		scrH:    scrH,
		history: list.New(),
	}
}

func (g *game) initRandom(fillFact float64) {
	rand.Seed(time.Now().UnixNano())
	for y := 0; y < g.tg.Height; y++ {
		for x := 0; x < g.tg.Width; x++ {
			if rand.Float64() < fillFact {
				g.tg.Buf.Set(x, y, tcg.Black)
			} else {
				g.tg.Buf.Set(x, y, tcg.White)
			}
		}
	}
	g.tg.Show()
}

func loadImage(in string) (image.Image, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("close file: %s", err)
		}
	}()

	img, _, err := image.Decode(file)
	return img, err
}

func (g *game) initFromImage(img image.Image) {
	buf := tcg.NewBufferFromImage(img)
	g.tg.Buf.BitBltAll(0, 0, buf)
	g.tg.Show()
}

func (g *game) doPause() {
	if g.mode == modePlay {
		g.mode = modePause
		if g.showCursor {
			g.curs.MoveAbs(g.tg.Buf, g.cursorX-curHalf, g.cursorY-curHalf).Put(g.tg.Buf)
		}
	} else {
		g.mode = modePlay
		if g.showCursor {
			g.curs.Withdraw(g.tg.Buf)
		}
	}
	g.tg.Show()
}

func (g *game) togglePixel() {
	defer g.handleCursor()()

	color := g.tg.Buf.At(g.cursorX, g.cursorY)
	color = color ^ 1
	g.tg.Buf.Set(g.cursorX, g.cursorY, color)
	g.tg.Show()
}

func (g *game) moveCursor(dx, dy int) {
	defer g.handleCursor()()
	oldX, oldY := g.cursorX, g.cursorY

	g.cursorX += dx
	g.cursorY += dy

	if g.cursorX < 0 {
		g.cursorX = 0
	}
	if g.cursorX >= g.tg.Width {
		g.cursorX = g.tg.Width - 1
	}
	if g.cursorY < 0 {
		g.cursorY = 0
	}
	if g.cursorY >= g.tg.Height {
		g.cursorY = g.tg.Height - 1
	}

	if (oldX != g.cursorX || oldY != g.cursorY) && g.pen {
		g.tg.Buf.Set(g.cursorX, g.cursorY, tcg.Black)
	}

	g.tg.Show()
}

func (g *game) toggleCursor() {
	defer g.handleCursor()()
	g.showCursor = !g.showCursor
}

func (g *game) handleCursor() func() {
	changed := false
	if g.showCursor {
		g.curs.Withdraw(g.tg.Buf)
		changed = true
	}

	return func() {
		if g.showCursor {
			g.curs.MoveAbs(g.tg.Buf, g.cursorX-curHalf, g.cursorY-curHalf).Put(g.tg.Buf)
			changed = true
		}
		if changed {
			g.tg.Show()
		}
	}
}

func (g *game) nextStep() {
	if g.showCursor && g.mode == modePause {
		defer g.handleCursor()()
	}

	startedAt := time.Now()
	g.generation++

	newGeneration := tcg.NewBuffer(g.tg.Width, g.tg.Height)

	for y := 0; y < g.tg.Height; y++ {
		for x := 0; x < g.tg.Width; x++ {
			neighbors := g.getNeighbors(x, y)
			oldCell := g.tg.Buf.At(x, y)
			switch {
			case oldCell == tcg.White && neighbors == 3:
				newGeneration.Set(x, y, tcg.Black)
			case oldCell == tcg.Black && (neighbors == 2 || neighbors == 3):
				newGeneration.Set(x, y, tcg.Black)
			default:
				newGeneration.Set(x, y, tcg.White)
			}
		}
	}

	// save to history
	g.history.PushFront(g.tg.Buf.Clone())
	if g.history.Len() > historySize {
		g.history.Remove(g.history.Back())
	}

	// copy to screen
	g.tg.Buf.BitBltAll(0, 0, newGeneration)
	g.updateStatMap(startedAt)
}

func (g *game) updateStatMap(startedAt time.Time) {
	fps := time.Second / time.Since(startedAt)
	if fps > maxFPS {
		fps = maxFPS
	}
	g.tg.PrintStr(3, g.scrH-1, fmt.Sprintf(" %4d FPS | %4d Gen ", fps, g.generation))
	g.tg.Show()
}

func (g *game) getNeighbors(x, y int) int {
	/*
		[-1, -1] [0, -1] [1, -1]
		[-1,  0]    *    [1,  0]
		[-1,  1] [0,  1] [1,  1]
	*/
	if !g.infMap {
		return g.tg.Buf.At(x-1, y-1) +
			g.tg.Buf.At(x, y-1) +
			g.tg.Buf.At(x+1, y-1) +
			g.tg.Buf.At(x-1, y) +
			g.tg.Buf.At(x+1, y) +
			g.tg.Buf.At(x-1, y+1) +
			g.tg.Buf.At(x, y+1) +
			g.tg.Buf.At(x+1, y+1)
	}

	sum := 0
	for xd := -1; xd <= 1; xd++ {
		for yd := -1; yd <= 1; yd++ {
			if xd == 0 && yd == 0 {
				continue
			}

			atX := x + xd
			if atX < 0 {
				atX = g.tg.Width - 1
			} else if atX >= g.tg.Width {
				atX = 0
			}

			atY := y + yd
			if atY < 0 {
				atY = g.tg.Height - 1
			} else if atY >= g.tg.Height {
				atY = 0
			}

			sum += g.tg.Buf.At(atX, atY)
		}
	}

	return sum
}

func (g *game) prevStep() {
	if g.history.Len() == 0 {
		return
	}

	defer g.handleCursor()()

	startedAt := time.Now()
	buf := g.history.Remove(g.history.Front()).(*tcg.Buffer)
	g.tg.Buf.BitBltAll(0, 0, *buf)
	g.generation--
	g.updateStatMap(startedAt)
}

func getCommand(tg *tcg.Tcg) chan cmds {
	resultCh := make(chan cmds)

	go func() {
		for {
			ev := tg.TCellScreen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:

				switch {
				case ev.Rune() == 'q' || ev.Key() == tcell.KeyEscape:
					resultCh <- cmdExit
				case ev.Rune() == 'p':
					resultCh <- cmdPause
				case ev.Rune() == 'l':
					resultCh <- cmdNext
				case ev.Rune() == 'h':
					resultCh <- cmdPrev
				case ev.Rune() == ' ':
					resultCh <- cmdPixel
				case ev.Rune() == 'a':
					resultCh <- cmdPen
				case ev.Rune() == 'c':
					resultCh <- cmdToggleCursor
				case ev.Key() == tcell.KeyRight:
					resultCh <- cmdRight
				case ev.Key() == tcell.KeyLeft:
					resultCh <- cmdLeft
				case ev.Key() == tcell.KeyUp:
					resultCh <- cmdUp
				case ev.Key() == tcell.KeyDown:
					resultCh <- cmdDown
				case ev.Rune() == 's':
					resultCh <- cmdScreenshot
				case ev.Rune() == 'w':
					resultCh <- cmdWipe
				}

			case *tcell.EventMouse:

				switch ev.Buttons() {
				case tcell.WheelUp:
					resultCh <- cmdNext
				case tcell.WheelDown:
					resultCh <- cmdPrev
				}

			}
		}
	}()

	return resultCh
}

func saveScreenshot(fileName string, buf tcg.Buffer) error {
	var bufBytes bytes.Buffer
	if err := png.Encode(&bufBytes, buf.ToImage()); err != nil {
		return err
	}

	return os.WriteFile(fileName, bufBytes.Bytes(), 0644)
}
