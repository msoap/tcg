package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

const delay = time.Millisecond * 100

type cmds int

const (
	cmdExit cmds = iota
	cmdPause
	cmdNext
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		panic(err)
	}

	initRandom(tg)

	ticker := time.Tick(delay)
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
				tg.PutPixel(x, y, tcg.Black)
			} else {
				tg.PutPixel(x, y, tcg.White)
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
			oldCell := tg.GetPixel(x, y)
			switch {
			case oldCell == tcg.White && neighbors == 3:
				newGeneration.PutPixel(x, y, tcg.Black)
			case oldCell == tcg.Black && (neighbors == 2 || neighbors == 3):
				newGeneration.PutPixel(x, y, tcg.Black)
			default:
				newGeneration.PutPixel(x, y, tcg.White)
			}
		}
	}

	// copy to screen
	tg.Buffer.BitBlt(0, 0, newGeneration)

	tg.Show()
}

func getNeighbors(tg tcg.Tcg, x, y int) int {
	return tg.GetPixel(x-1, y-1) +
		tg.GetPixel(x, y-1) +
		tg.GetPixel(x+1, y-1) +
		tg.GetPixel(x-1, y) +
		tg.GetPixel(x+1, y) +
		tg.GetPixel(x-1, y+1) +
		tg.GetPixel(x, y+1) +
		tg.GetPixel(x+1, y+1)
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
