[![Go Reference](https://pkg.go.dev/badge/github.com/msoap/tcg.svg)](https://pkg.go.dev/github.com/msoap/tcg)
[![GitHub Action](https://github.com/msoap/tcg/actions/workflows/go.yml/badge.svg)](https://github.com/msoap/tcg/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/msoap/tcg/badge.svg?branch=master)](https://coveralls.io/github/msoap/tcg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/msoap/tcg)](https://goreportcard.com/report/github.com/msoap/tcg)

# TCG - terminal cell graphics

Go Graphics library for use in a text terminal. Only 1bit graphics can be used with two colors. Used unicode block symbols for drawing. 2x3 mode is supported by the Iosevka font or some terminals (like default on Ubuntu).

## Features

  * Available 5 [graphics modes](https://pkg.go.dev/github.com/msoap/tcg#PixelMode), from 2x3 pixels grid for terminal symbol to 1x1, 1x2 and 2x2, or you can [create](https://pkg.go.dev/github.com/msoap/tcg#NewPixelMode) own custom mode with your graphics
  * [Set](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Set)/[get](https://pkg.go.dev/github.com/msoap/tcg#Buffer.At) one pixel
  * Drawings: [lines](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Line) (vertical, horizontal, or with any angle), [boxes](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Rect), [circles](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Circle), [arcs](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Arc)
  * [Fill](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Fill) area with a different [options](https://pkg.go.dev/github.com/msoap/tcg#FillOpt), for example fill with [patterns](https://pkg.go.dev/github.com/msoap/tcg#WithPattern)
  * Buffer manipulating: [cut](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Cut), [clone](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Clone), convert to/from stdlib [Image](https://pkg.go.dev/github.com/msoap/tcg#Buffer.ToImage) or text
  * Buffer transform: [BitBlt](https://pkg.go.dev/github.com/msoap/tcg#Buffer.BitBlt), [clear](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Clear), [flip](https://pkg.go.dev/github.com/msoap/tcg#Buffer.HFlip), [invert](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Invert), scroll ([vertical](https://pkg.go.dev/github.com/msoap/tcg#Buffer.VScroll), [horizontal](https://pkg.go.dev/github.com/msoap/tcg#Buffer.HScroll))
  * Sub-package for [turtle graphics](https://pkg.go.dev/github.com/msoap/tcg/turtle), also available [drawing](https://pkg.go.dev/github.com/msoap/tcg@v0.0.1/turtle#Turtle.DrawScript) by text script

## Install

    go get github.com/msoap/tcg

## Usage

```go
package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/msoap/tcg"
)

func main() {
	tg, err := tcg.New(tcg.Mode2x3) // each terminal symbol contains a 2x3 pixels grid, also you can use 1x1, 1x2, and 2x2 modes
	if err != nil {
		log.Fatalf("create tg: %s", err)
	}

	i := 0
	for {
		pixColor := tg.Buf.At(10, 10)       // get color of pixel
		tg.Buf.Set(11, 11, pixColor)        // draw one pixel with color from 10,10
		tg.Buf.Line(0, 0, 50, i, tcg.Black) // draw a diagonal line
		tg.Show()                           // synchronize buffer with screen

		if ev, ok := tg.TCellScreen.PollEvent().(*tcell.EventKey); ok && ev.Rune() == 'q' {
			break // exit by 'q' key
		}
		i++
	}

	tg.Finish() // finish application and restore screen
}
```

See more examples in [examples](https://github.com/msoap/tcg/tree/master/examples) folder.

## Screenshot

[Game of Life](https://github.com/msoap/tcg/tree/master/examples/game_of_life) example in iTerm2 terminal:

<img width="843" alt="TCG library example screenshot for Game of Life" src="https://user-images.githubusercontent.com/844117/222931477-2aac0d97-051a-4261-a445-c1d4c12705de.png">

See more [screenshots](https://github.com/msoap/tcg/wiki/Screenshots).

## TODO

  * [ ] fonts support
  * [ ] sprites, maybe with animation

## See also

  * Go terminal library - [github.com/gdamore/tcell](https://github.com/gdamore/tcell/)
  * [Turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics)
  * Try it in the Go [Playground](https://go.dev/play/p/iX_-qby2SbT)

Unicode symbols:

  * Block Elements - [wikipedia](https://en.wikipedia.org/wiki/Block_Elements)
  * Block Elements - [unicode.org](https://www.unicode.org/charts/PDF/U2580.pdf)
  * Symbols for Legacy Computing - [wikipedia](https://en.wikipedia.org/wiki/Symbols_for_Legacy_Computing)
  * Symbols for Legacy Computing - [unicode.org](http://unicode.org/charts/PDF/U1FB00.pdf)

Supported fonts (for "2x3" mode):

  * [Iosevka font](https://github.com/be5invis/Iosevka)
