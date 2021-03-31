package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY := tg.Width, tg.Height
	x, y := maxX/2, maxY/2
	drawNext := func() {
		tg.Buffer.Set(x, y, tcg.Black)
		tg.Show()
	}

LOOP:
	for {
		ev := tg.TCellScreen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Rune() == 'q' {
				break LOOP
			}
			switch ev.Key() {
			case tcell.KeyDown:
				y++
				drawNext()
			case tcell.KeyUp:
				y--
				drawNext()
			case tcell.KeyLeft:
				x--
				drawNext()
			case tcell.KeyRight:
				x++
				drawNext()
			case tcell.KeyEscape:
				break LOOP
			}
		}
	}

	tg.Finish()
}
