package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

const defaultDelay = time.Millisecond * 100

type cmds int

const (
	cmdExit cmds = iota
	cmdPause
	cmdNext
)

func main() {
	delay := flag.Duration("delay", defaultDelay, "delay between steps")
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3")
	flag.Parse()

	tg, err := tcg.New(mode)
	if err != nil {
		panic(err)
	}

	initRandom(tg)

	ticker := time.Tick(*delay)
	command := getCommand(tg)
	paused := false

LOOP:
	for {
		select {
		case <-ticker:
			if !paused {
				nextStep(tg)
			}
		case cmd := <-command:
			switch cmd {
			case cmdExit:
				break LOOP
			case cmdPause:
				paused = !paused
			case cmdNext:
				nextStep(tg)
			}
		}
	}

	tg.Finish()
}

func initRandom(tg tcg.Tcg) {
	rand.Seed(time.Now().UnixNano())
	for y := 0; y < tg.Height; y++ {
		for x := 0; x < tg.Width; x++ {
			if rand.Float64() < 0.2 {
				tg.Buf.Set(x, y, tcg.Black)
			} else {
				tg.Buf.Set(x, y, tcg.White)
			}
		}
	}
	tg.Show()
}

func nextStep(tg tcg.Tcg) {
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
}

func getNeighbors(tg tcg.Tcg, x, y int) int {
	return tg.Buf.At(x-1, y-1) +
		tg.Buf.At(x, y-1) +
		tg.Buf.At(x+1, y-1) +
		tg.Buf.At(x-1, y) +
		tg.Buf.At(x+1, y) +
		tg.Buf.At(x-1, y+1) +
		tg.Buf.At(x, y+1) +
		tg.Buf.At(x+1, y+1)
}

func getCommand(tg tcg.Tcg) chan cmds {
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
