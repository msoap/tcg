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
    gr := tcg.New()
    gr.PutPixel(10, 10, tcg.Black)
    gr.tcell.Show()
}
```

## See also

  * [github.com/gdamore/tcell](https://github.com/gdamore/tcell/)