package main

import (
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

	list := buf.RenderAsStrings(tcg.Mode2x3)
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
