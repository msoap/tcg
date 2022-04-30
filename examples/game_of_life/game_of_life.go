package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

const (
	defaultDelay          = time.Millisecond * 100
	defaultInitFillFactor = 0.2
)

type cmds int

const (
	cmdExit cmds = iota
	cmdPause
	cmdNext
)

func main() {
	delay := flag.Duration("delay", defaultDelay, "delay between steps")
	size := flag.String("size", "", "screen size, in 'width x height' format, example: '80x25'")
	fillFactor := flag.Float64("fill", defaultInitFillFactor, "how much to fill the area initially")
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3")
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

	tg, err := tcg.New(mode)
	if err != nil {
		log.Fatal(err)
	}
	_, scrH := tg.ScreenSize()

	pattern := tcg.MustNewBufferFromStrings([]string{
		" *",
		"* ",
	})
	tg.Buf.Fill(0, 0, tcg.WithPattern(pattern))
	tg.Show()

	if width == 0 {
		width, height = tg.ScreenSize()
	} else {
		if err := tg.SetClipCenter(width, height); err != nil {
			tg.Finish()
			log.Fatal(err)
		}
	}

	tg.Buf.Rect(0, 0, tg.Width, tg.Height, tcg.Black) // coordinates in pixels
	tg.PrintStr(5, 1, " Game of Life ")               // coordinates in chars, not pixels
	tg.PrintStr(15, scrH-1, ` <q> - Quit <p> - Pause <Right> Next step `)
	tg.Show()

	if err := tg.SetClipCenter(width-2, height-2); err != nil {
		tg.Finish()
		log.Fatal(err)
	}
	initRandom(tg, *fillFactor)

	ticker := time.Tick(*delay)
	command := getCommand(tg)
	paused := false

LOOP:
	for {
		select {
		case <-ticker:
			if !paused {
				nextStep(scrH, tg)
			}
		case cmd := <-command:
			switch cmd {
			case cmdExit:
				break LOOP
			case cmdPause:
				paused = !paused
			case cmdNext:
				nextStep(scrH, tg)
			}
		}
	}

	tg.Finish()
}

func initRandom(tg *tcg.Tcg, fillFact float64) {
	rand.Seed(time.Now().UnixNano())
	for y := 0; y < tg.Height; y++ {
		for x := 0; x < tg.Width; x++ {
			if rand.Float64() < fillFact {
				tg.Buf.Set(x, y, tcg.Black)
			} else {
				tg.Buf.Set(x, y, tcg.White)
			}
		}
	}
	tg.Show()
}

func nextStep(scrH int, tg *tcg.Tcg) {
	startedAt := time.Now()

	newGeneration := tcg.NewBuffer(tg.Width, tg.Height)

	for y := 0; y < tg.Height; y++ {
		for x := 0; x < tg.Width; x++ {
			neighbors := getNeighbors(tg, x, y)
			oldCell := tg.Buf.At(x, y)
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

	// copy to screen
	tg.Buf.BitBltAllSrc(0, 0, newGeneration)

	tg.Show()

	tg.PrintStr(3, scrH-1, fmt.Sprintf(" %-3d FPS ", time.Second/time.Since(startedAt)))
}

func getNeighbors(tg *tcg.Tcg, x, y int) int {
	return tg.Buf.At(x-1, y-1) +
		tg.Buf.At(x, y-1) +
		tg.Buf.At(x+1, y-1) +
		tg.Buf.At(x-1, y) +
		tg.Buf.At(x+1, y) +
		tg.Buf.At(x-1, y+1) +
		tg.Buf.At(x, y+1) +
		tg.Buf.At(x+1, y+1)
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
				case ev.Rune() == 'p' || ev.Rune() == ' ':
					resultCh <- cmdPause
				case ev.Key() == tcell.KeyRight:
					resultCh <- cmdNext
				}
			}
		}
	}()

	return resultCh
}
