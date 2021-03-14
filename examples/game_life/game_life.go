package main

import (
	"math/rand"
	"time"

	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		panic(err)
	}

	initRandom(tg)

	tg.Show()
	time.Sleep(5 * time.Second)
	tg.Finish()
}

func initRandom(tg tcg.Tcg) {
	for y := 0; y < tg.Height; y++ {
		for x := 0; x < tg.Width; x++ {
			if rand.Float64() < 0.1 {
				tg.PutPixel(x, y, tcg.Black)
			} else {
				tg.PutPixel(x, y, tcg.White)
			}
		}
	}
}
