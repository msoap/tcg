package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msoap/byline"
	"github.com/msoap/tcg"
)

/*
create PBM files with ImageMagick:
  convert img01.png -compress none -random-threshold 70% img01.pbm

usage:
  go run examples/pbm_view/pbm_view.go examples/pbm_view/img01.pbm
*/

func main() {
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille")
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("file name of .pbm file needed: %+v", args)
	}

	fReader, err := os.Open(args[0])
	if err != nil {
		log.Fatalf("failed to open file %s: %s", args[0], err)
	}

	buf, err := reader2buffer(fReader)
	if err != nil {
		log.Fatalf("failed to load pbm file: %s", err)
	}
	_ = fReader.Close()

	list := buf.RenderAsStrings(mode)
	for _, line := range list {
		fmt.Println(line)
	}
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
				buf.Set(x, y, color)
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
