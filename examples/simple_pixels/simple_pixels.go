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

	tg.Buf.HLine(0, 4, 5, tcg.Black)
	tg.Buf.HLine(5, 5, 5, tcg.Black)
	tg.Buf.HLine(10, 6, 5, tcg.Black)

	tg.Buf.Set(0, 0, tcg.Black)
	tg.Buf.Set(2, 0, tcg.Black)
	tg.Buf.Set(1, 1, tcg.Black)
	tg.Buf.Set(3, 1, tcg.Black)
	tg.Buf.Set(6, 8, tcg.Black)
	tg.Buf.Set(8, 8, tcg.Black)
	tg.Buf.Set(10, 8, tcg.Black)
	tg.Buf.Set(12, 8, tcg.Black)

	tg.Buf.Box(3, 10, 20, 5, tcg.Black)
	tg.Buf.FillBox(60, 1, 15, 14, tcg.Black)

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

	tg.Buf.BitBltAllSrc(40, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buf.BitBltAllSrc(55, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buf.BitBltAllSrc(70, 20, buf)
	tg.Show()
	time.Sleep(1 * time.Second)

	// second moving
	buf.BitBlt(0, 0, 10, 10, tg.Buf, 5, 5)
	tg.Buf.BitBlt(40, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)
	tg.Buf.BitBlt(55, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)
	tg.Buf.BitBlt(70, 40, 10, 10, buf, 0, 0)

	tg.PrintStr(60, 10, "Hello")
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buf.FillBox(0, 0, tg.Width, tg.Height, tcg.White) // clear
	for _, step := range []int{5, 6, 7, 17, 33} {
		for y := 0; y < tg.Height; y++ {
			for x := 0; x < tg.Width; x++ {
				c := tcg.White
				if (x^y)%step == 0 {
					c = tcg.Black
				}
				tg.Buf.Set(x, y, c)
			}
		}
		tg.Show()
		time.Sleep(1 * time.Second)
	}

	tg.Finish()
}
