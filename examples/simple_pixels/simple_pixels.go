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

	tg.Buffer.Set(0, 0, tcg.Black)
	tg.Buffer.Set(2, 0, tcg.Black)
	tg.Buffer.Set(1, 1, tcg.Black)
	tg.Buffer.Set(3, 1, tcg.Black)
	tg.Buffer.Set(6, 8, tcg.Black)
	tg.Buffer.Set(8, 8, tcg.Black)
	tg.Buffer.Set(10, 8, tcg.Black)
	tg.Buffer.Set(12, 8, tcg.Black)

	tg.Box(3, 10, 20, 5, tcg.Black)

	tg.Show()
	time.Sleep(1 * time.Second)

	// first moving
	buf := tcg.NewBuffer(10, 10)
	buf.Set(0, 0, tcg.Black)
	buf.Set(1, 0, tcg.Black)
	buf.Set(2, 0, tcg.Black)
	buf.Set(0, 0, tcg.Black)
	buf.Set(0, 1, tcg.Black)
	buf.Set(0, 2, tcg.Black)

	tg.Buffer.BitBltAllSrc(40, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buffer.BitBltAllSrc(55, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buffer.BitBltAllSrc(70, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	// second moving
	buf.BitBlt(0, 0, 10, 10, tg.Buffer, 5, 5)
	tg.Buffer.BitBlt(40, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)
	tg.Buffer.BitBlt(55, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)
	tg.Buffer.BitBlt(70, 40, 10, 10, buf, 0, 0)
	tg.Show()

	time.Sleep(5 * time.Second)
	tg.Finish()
}
