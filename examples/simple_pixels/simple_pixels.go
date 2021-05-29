package main

import (
	"time"

	"github.com/msoap/tcg"
)

var letterA = []string{
	"00011000",
	"00100100",
	"00111100",
	"00100100",
	"00100100",
}

var letterB = []string{
	"00111000",
	"00100100",
	"00111000",
	"00100100",
	"00111000",
}

var letterC = []string{
	"00011100",
	"00100010",
	"00100000",
	"00100010",
	"00011100",
}

var pattern = []string{
	"01",
	"10",
}

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

	tg.Buf.Rect(3, 10, 30, 20, tcg.Black)
	tg.Buf.Fill(5, 11, tcg.WithPattern(tcg.MustNewBufferFromStrings(pattern)))
	tg.Buf.FillRect(60, 1, 15, 14, tcg.Black)

	tg.Buf.Line(0, 40, 25, 60, tcg.Black)

	// letters
	tg.Buf.BitBltAllSrc(40, 30, tcg.MustNewBufferFromStrings(letterA))
	tg.Buf.BitBltAllSrc(48, 30, tcg.MustNewBufferFromStrings(letterB))
	tg.Buf.BitBltAllSrc(56, 30, tcg.MustNewBufferFromStrings(letterC))

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
	time.Sleep(2 * time.Second)

	// second moving
	buf.BitBlt(0, 0, 10, 10, tg.Buf, 5, 5)
	tg.Buf.BitBlt(40, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buf.BitBlt(55, 40, 10, 10, buf, 0, 0)
	tg.Show()
	time.Sleep(1 * time.Second)

	tg.Buf.BitBlt(70, 40, 10, 10, buf, 0, 0)
	tg.Show()

	tg.PrintStr(25, 8, "Hello World!")
	time.Sleep(3 * time.Second)

	tg.SetClipCenter(100, 30)
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
