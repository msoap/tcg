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

	tg.PutPixel(3, 7, tcg.Black)
	tg.PutPixel(4, 7, tcg.Black)
	tg.PutPixel(5, 7, tcg.Black)

	tg.Show()
	time.Sleep(5 * time.Second)
	tg.Finish()
}
