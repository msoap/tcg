# TCG - terminal graphics library

Used unicode block symbols for drawing. 2x3 mode is supported by the latest versions of the Iosevka and JetBrains Mono fonts.

See:

  * [Block Elements - wikipedia](https://en.wikipedia.org/wiki/Block_Elements)
  * [Block Elements - unicode.org](https://www.unicode.org/charts/PDF/U2580.pdf)
  * [Symbols for Legacy Computing - wikipedia](https://en.wikipedia.org/wiki/Symbols_for_Legacy_Computing)
  * [Symbols for Legacy Computing - unicode.org](http://unicode.org/charts/PDF/U1FB00.pdf)

## Install

    go get -u github.com/msoap/tcg

## Usage

```go
import (
    "github.com/msoap/tcg"
)

main () {
    tg := tcg.New(tcg.Mode2x3)
    tg.PutPixel(10, 10, tcg.Black)
    pix := tg.GetPixel(10, 10) // tcg.Black
    tg.PrintStr(20, 20, "Hello world!")
    tg.Show()
    tg.Finish()
}
```

## See also

  * [Go library for terminal - github.com/gdamore/tcell](https://github.com/gdamore/tcell/)
  * [Iosevka font](https://github.com/be5invis/Iosevka)
  * [JetBrains Mono font](https://github.com/JetBrains/JetBrainsMono)
