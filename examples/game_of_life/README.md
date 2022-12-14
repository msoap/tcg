# Game Of Life

This is small example of "Game Of Life", see [Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) for more information.

## Install

```
go install github.com/msoap/tcg/examples/game_of_life@latest
```

## Usage

```
game_of_life [options]
options:
  -color string
    	redefine color, it can be: 'yellow', 'red' or like '#ffaa11'
  -delay duration
    	delay between steps (default 100ms)
  -fill float
    	how much to fill the area initially (default 0.2)
  -mode value
    	screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille (default 2x3)
  -out string
    	save screenshot to file (default "game_of_life.png")
  -size string
    	screen size in chars, in 'width x height' format, example: '80x25'
```

## Shortcuts

  * `esc` / `q` - quit
  * `s` - save screenshot
  * `space` / `p` - pause
  * `right` - step forward

## Screenshots

<img width="663" alt="TCG library example screenshot for Game of Life" src="https://user-images.githubusercontent.com/844117/153767605-76dd1552-9424-49b9-9bf3-9163132af9b2.png">

<img width="906" alt="TCG library example screenshot for Game of Life" src="https://user-images.githubusercontent.com/844117/207433839-7a15d70e-9258-4943-a9ef-34ea9c139ca5.png">
