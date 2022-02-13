[![Go Reference](https://pkg.go.dev/badge/github.com/msoap/tcg.svg)](https://pkg.go.dev/github.com/msoap/tcg)
[![GitHub Action](https://github.com/msoap/tcg/actions/workflows/go.yml/badge.svg)](https://github.com/msoap/tcg/actions/workflows/go.yml)

# TCG - terminal cell graphics

Go Graphics library for use in a text terminal. Only 1bit graphics can be used with two colors. Used unicode block symbols for drawing. 2x3 mode is supported by the latest versions of the Iosevka font.

## Features

  * Available 4 [graphics mode](https://pkg.go.dev/github.com/msoap/tcg#PixelsInChar), from 2x3 pixels grid for terminal symbol to 1x1, 1x2 and 2x2
  * [Set](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Set)/[get](https://pkg.go.dev/github.com/msoap/tcg#Buffer.At) one pixel
  * Drawings: [lines](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Line) (vertical, horizontal, or with any angle), [boxes](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Rect), [circles](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Circle), [arcs](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Arc)
  * [Fill](https://pkg.go.dev/github.com/msoap/tcg#Buffer.Fill) area with a different [options](https://pkg.go.dev/github.com/msoap/tcg#FillOpt), for example fill with [patterns](https://pkg.go.dev/github.com/msoap/tcg#WithPattern)
  * Buffer manipulating, copy, cut, clone, convert to/from stdlib [Image](https://pkg.go.dev/github.com/msoap/tcg#Buffer.ToImage) or text
  * Buffer transform: [BitBlt](https://pkg.go.dev/github.com/msoap/tcg#Buffer.BitBlt), clear, flip, invert, scroll (vertical, horizontal)
  * Sub-package for [turtle graphics](https://pkg.go.dev/github.com/msoap/tcg/turtle), also available drawing by text script

## Install

    go get -u github.com/msoap/tcg

## Usage

```go
import (
    "github.com/msoap/tcg"
)

main () {
    tg := tcg.New(tcg.Mode2x3) // each terminal symbol contains a 2x3 pixels grid, also you can use 1x1, 1x2, and 2x2 modes
    for {
        tg.Set(10, 10, tcg.Black)  // draw one pixel
        pix := tg.At(10, 10)       // get color of pixel
        tg.Show()                  // synchronize buffer with screen

        time.Sleep(time.Millisecond * 100) // 10 FPS
        if doExit {
            break
        }
    }
    tg.Finish()                // finish application and restore screen
}
```

See more examples in `examples` folder.

## Screenshot

Game of Life exmple in iTerm2 terminal with Iosevka font:

<img width="663" alt="Screenshot 2022-02-13 at 19 12 58 " src="https://user-images.githubusercontent.com/844117/153767605-76dd1552-9424-49b9-9bf3-9163132af9b2.png">

## TODO

  * [ ] fonts support
  * [ ] sprites, maybe with animation

## See also

  * [Go library for terminal - github.com/gdamore/tcell](https://github.com/gdamore/tcell/)
  * [Turtle graphics](https://en.wikipedia.org/wiki/Turtle_graphics)

Unicode symbols:

  * [Block Elements - wikipedia](https://en.wikipedia.org/wiki/Block_Elements)
  * [Block Elements - unicode.org](https://www.unicode.org/charts/PDF/U2580.pdf)
  * [Symbols for Legacy Computing - wikipedia](https://en.wikipedia.org/wiki/Symbols_for_Legacy_Computing)
  * [Symbols for Legacy Computing - unicode.org](http://unicode.org/charts/PDF/U1FB00.pdf)

Supported fonts:

  * [Iosevka font](https://github.com/be5invis/Iosevka)
