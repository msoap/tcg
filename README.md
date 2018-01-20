# TCG - tcell graphics library

## Install

    go get -u github.com/msoap/tcg

## Usage

```go
import (
    "github.com/msoap/tcg"
    "github.com/gdamore/tcell"
)

main () {
    tg := tcg.New()
    tg.PutPixel(10, 10, tcg.Black)
    pix := tg.GetPixel(10, 10) // tcg.Black
    tg.PrintStr(20, 20, "Hello world!")
    tg.Show()
    tg.Finish()
}
```

## See also

  * [github.com/gdamore/tcell](https://github.com/gdamore/tcell/)
