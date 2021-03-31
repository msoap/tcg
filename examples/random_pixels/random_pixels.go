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

	for i := 0; i < 100; i++ {
		for y := 0; y < tg.Height; y++ {
			for x := 0; x < tg.Width; x++ {
				if rand.Float64() < 0.5 {
					tg.Buffer.Set(x, y, tcg.Black)
				} else {
					tg.Buffer.Set(x, y, tcg.White)
				}
			}
		}
		tg.Show()
		time.Sleep(10 * time.Millisecond)
	}

	tg.Finish()
}
