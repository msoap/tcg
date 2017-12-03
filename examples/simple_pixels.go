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
	tg.Show()
	time.Sleep(5 * time.Second)
	tg.Finish()
}
