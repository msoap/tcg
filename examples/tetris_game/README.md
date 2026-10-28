# Tetris

Classic [Tetris](https://en.wikipedia.org/wiki/Tetris) game in a terminal, an example of using the TCG library in the 2x3 and 2x4Braille modes.

Features:

  - 7 classic blocks, each with its own color, chosen randomly; the next block is shown
  - blocks fall smoothly pixel by pixel, not by characters
  - full lines are erased pixel row by pixel row, then everything above falls down in one step; the animation is quick (about 0.12s) and the game doesn't wait for it
  - score: +10 for each line, +5 more for each level above the first; each 10 lines give a new level and +20% to the falling speed
  - the screen is redrawn without flickering, keys are handled immediately

## Install

```
go install github.com/msoap/tcg/examples/tetris_game@latest
```

Or just run immediately without installing:

```
go run github.com/msoap/tcg/examples/tetris_game@latest
```

The 2x3 mode needs a font or terminal with support of "Symbols for Legacy Computing", see the [main README](../../README.md).
With the default sizes the game needs a terminal of at least 44x27 characters.

## Usage

```
tetris_game [-mode 2x3|2x4Braille]
```

  - `2x3` - default mode, it needs a font with the "Symbols for Legacy Computing" (see above)
  - `2x4Braille` - mode with Braille symbols, it works in most fonts, but the blocks look dotted

## Controls

| Key           | Action                    |
|---------------|---------------------------|
| `←` / `→`     | move the block            |
| `↑`           | rotate the block clockwise |
| `space`       | drop the block            |
| `p`           | pause                     |
| `r`           | restart                   |
| `?`           | show/hide help            |
| `q`           | quit                      |

## Settings

Sizes of the game area and a square, the initial speed, scoring, the speed of the animation of removing lines, etc. are constants at the top of [tetris_game.go](tetris_game.go).
A square is two characters wide and one character high (`squareCharsW`, `squareCharsH`), so it's 4x3 pixels in the 2x3 mode and 4x4 pixels in the 2x4Braille mode.
The speed of falling is in pixels per second, so in the 2x4Braille mode (4 pixels per character row) blocks fall a bit slower than in the 2x3 mode (3 pixels per row).

## Screenshot

<img width="752" height="608" alt="image" src="https://github.com/user-attachments/assets/6571d00b-9a45-4ba9-a6b9-c747e761de9b" />
