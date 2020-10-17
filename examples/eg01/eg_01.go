package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func main() {
	hideCursor()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		showCursor()
		os.Exit(0)
	}()
	defer showCursor()

	framesBy3 := []string{"▌ ", "🬆 ", "🬂🬀", "🬁🬄", " ▌", "🬞🬓", "🬭🬏", "🬱 "}

	for i := 0; i < 10; i++ {
		for _, item := range framesBy3 {
			fmt.Print(item)
			time.Sleep(time.Millisecond * 100)
			fmt.Print("\r")
		}
	}

	fmt.Printf("\n")
	time.Sleep(time.Second * 1)
}

func hideCursor() {
	tput("civis")
}

func showCursor() {
	tput("cvvis")
}

func tput(arg string) error {
	cmd := exec.Command("tput", arg)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
