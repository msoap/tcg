package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/byline"
	"github.com/msoap/tcg"
)

/*
create PBM files with ImageMagick:
  convert img01.png -compress none -random-threshold 25% img01.pbm
*/

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("file name of .pbm file needed")
	}

	fReader, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file %s: %s", os.Args[1], err)
	}

	buf, err := reader2buffer(fReader)
	if err != nil {
		log.Fatalf("failed to load pbm file: %s", err)
	}
	_ = fReader.Close()

	tg, err := tcg.New(tcg.Mode2x3)
	if err != nil {
		log.Fatalf("failed to create tcg: %s", err)
	}

	for y := 0; y < buf.Height && y < tg.Height; y++ {
		for x := 0; x < buf.Width && x < tg.Width; x++ {
			tg.PutPixel(x, y, buf.GetPixel(x, y)) // TODO: implement BitBlt
		}
	}

	tg.Show()
	<-getEscape(tg)
	tg.Finish()
}

func reader2buffer(in io.Reader) (tcg.Buffer, error) {
	var buf tcg.Buffer
	status := "begin" // "begin" -> "header" -> "size"
	y := 0

	err := byline.NewReader(in).AWKMode(func(line string, fields []string, vars byline.AWKVars) (string, error) {
		switch {
		case strings.HasPrefix(line, "#"):
			return line, nil
		case status == "begin" && vars.NF >= 1 && fields[0] == "P1":
			status = "header"
		case status == "header" && vars.NF == 2 && len(fields) == 2:
			width, _ := strconv.Atoi(fields[0])
			height, _ := strconv.Atoi(fields[1])
			if width == 0 || height == 0 {
				return "", fmt.Errorf("failed to parse pbm size on line %d: %q", vars.NR, line)
			}

			status = "size"
			buf = tcg.NewBuffer(width, height)
		case status == "size" && vars.NF > 0 && len(fields) >= buf.Width:
			for x := 0; x < len(fields) && x < buf.Width; x++ {
				color := tcg.White
				if fields[x] == "1" {
					color = tcg.Black
				}
				buf.PutPixel(x, y, color)
			}
			y++
		default:
			return "", fmt.Errorf("failed to parse pbm file on line %d: %q", vars.NR, line)
		}

		return line, nil
	}).Discard()

	if err != nil {
		return buf, err
	}

	return buf, nil
}

func getEscape(tg tcg.Tcg) chan struct{} {
	resultCh := make(chan struct{})

	go func() {
		for {
			ev := tg.TCellScreen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					resultCh <- struct{}{}
				}
			}
		}
	}()

	return resultCh
}
