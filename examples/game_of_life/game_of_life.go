package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		panic(err)
	}

	initRandom(tg)

	ticker := time.NewTicker(time.Millisecond * 100)
	escape := getEscape(tg)

LOOP:
	for {
		select {
		case <-ticker.C:
			nextStep(tg)
		case <-escape:
			break LOOP
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
	// TODO: implement bit blt from buffer to buffer
	for y := 0; y < tg.Height; y++ {
		for x := 0; x < tg.Width; x++ {
			tg.PutPixel(x, y, newGeneration.GetPixel(x, y))
		}
	}

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

func getEscape(tg tcg.Tcg) chan struct{} {
	resultCh := make(chan struct{})

	go func() {
		for {
			ev := tg.TCellScreen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					resultCh <- struct{}{}
				}
			}
		}
	}()

	return resultCh
}
