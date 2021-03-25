package main

import (
	"time"

	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		panic(err)
	}

	tg.HLine(0, 4, 5, tcg.Black)
	tg.HLine(5, 5, 5, tcg.Black)
	tg.HLine(10, 6, 5, tcg.Black)

	tg.PutPixel(0, 0, tcg.Black)
	tg.PutPixel(2, 0, tcg.Black)
	tg.PutPixel(1, 1, tcg.Black)
	tg.PutPixel(3, 1, tcg.Black)
	tg.PutPixel(6, 8, tcg.Black)
	tg.PutPixel(8, 8, tcg.Black)
	tg.PutPixel(10, 8, tcg.Black)
	tg.PutPixel(12, 8, tcg.Black)

	tg.Box(3, 10, 20, 5, tcg.Black)

	tg.Show()
	time.Sleep(1 * time.Second)

	buf := tcg.NewBuffer(10, 10)
	buf.PutPixel(0, 0, tcg.Black)
	buf.PutPixel(1, 0, tcg.Black)
	buf.PutPixel(2, 0, tcg.Black)
	buf.PutPixel(0, 0, tcg.Black)
	buf.PutPixel(0, 1, tcg.Black)
	buf.PutPixel(0, 2, tcg.Black)

	tg.Buffer.BitBltAllSrc(40, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buffer.BitBltAllSrc(55, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buffer.BitBltAllSrc(70, 20, buf)
	tg.Show()

	time.Sleep(5 * time.Second)
	tg.Finish()
}
