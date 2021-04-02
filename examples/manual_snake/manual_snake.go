package main

import (
	"flag"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3")
	flag.Parse()

	tg, err := tcg.New(mode)
	if err != nil {
		log.Fatal(err)
	}

	maxX, maxY := tg.Width, tg.Height
	x, y := maxX/2, maxY/2
	drawNext := func() {
		tg.Buf.Set(x, y, tcg.Black)
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
