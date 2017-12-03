package main

import (
	"time"

	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New()
	if err != nil {
		panic(err)
	}

	tg.PrintStr(3, 4, "Hello world!")

	tg.HLine(3, 6, 12, tcg.Black)
	tg.VLine(3, 6, 5, tcg.Black)

	tg.PutPixel(5, 8, tcg.Black)
	tg.PutPixel(7, 8, tcg.Black)
	tg.PutPixel(9, 8, tcg.Black)
	tg.PutPixel(11, 8, tcg.Black)
	tg.PutPixel(13, 8, tcg.Black)

	tg.HLine(3, 10, 12, tcg.Black)
	tg.VLine(15, 6, 5, tcg.Black)

	tg.Box(3, 20, 20, 5, tcg.Black)

	tg.Show()
	time.Sleep(5 * time.Second)
	tg.Finish()
}
