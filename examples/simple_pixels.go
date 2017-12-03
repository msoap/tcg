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

	tg.PutPixel(4, 8, tcg.Black)
	tg.PutPixel(6, 8, tcg.Black)
	tg.PutPixel(8, 8, tcg.Black)
	tg.PutPixel(10, 8, tcg.Black)
	tg.PutPixel(12, 8, tcg.Black)

	tg.Box(3, 10, 20, 5, tcg.Black)

	tg.Show()
	time.Sleep(5 * time.Second)
	tg.Finish()
}
