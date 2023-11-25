# Game Of Life

This is an example of "Game Of Life", see [Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) for more information.

Features:
  
  - predefined maps from image files
  - save map as screenshot to image file, later you can load it
  - edit map with keyboard, by one pixel or with pen mode, can show/hide cursor
  - infinite map (wrap around edges)
  - mouse support (scroll up/down to step forward/backward)
  - different screen modes (1x1, 1x2, 2x2, 2x3, 2x4Braille)
  - allow to configure color, delay, fill, size, etc.

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
  -in string
    	load map from image file (*.png)
  -mode value
    	screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille (default 2x3)
  -out string
    	save map as screenshot to file (default "game_of_life.png")
  -size string
    	screen size in chars, in 'width x height' format, example: '80x25'
  -wait
    	wait for start
  -inf
    	infinite map (wrap around edges)
```

Use predefined map from image file:


```shell
game_of_life -in examples/game_of_life/glider.png -inf
game_of_life -in examples/game_of_life/gospers_glider_gun.png
game_of_life -in examples/game_of_life/penta_decathlon.png
```

## Shortcuts

  * `esc` / `q` - quit
  * `s` - save map as screenshot (to file, see `-out` option)
  * `p` - pause
  * `l` - step forward, or mouse scroll up with terminal mouse support
  * `h` - step to previous state, or mouse scroll down
  * `Space` - toggle pixel under cursor
  * `a` - toggle pen mode, when pen mode is on, you can draw with arrows keys
  * `w` - wipe map
  * `c` - toggle cursor visibility
  * `←`, `↑`, `→`, `↓` - move cursor

## Screenshots

MacOS iTerm2 (2x3 mode):

<img width="826" alt="TCG library example screenshot for Game of Life" src="https://user-images.githubusercontent.com/844117/226143976-4db0f377-0195-4c8d-8dea-799eee29d9ce.png">

Ubuntu GNOME Terminal (2x4Braille mode):

<img width="666" alt="TCG library example screenshot for Game of Life on Ubuntu" src="https://user-images.githubusercontent.com/844117/222967488-3c07917e-f90f-4843-b987-fc97b3397a19.png">
