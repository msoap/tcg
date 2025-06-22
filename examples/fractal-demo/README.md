# Fractal demo

This is a simple animated fractal demo in terminal. It can show Mandelbrot and Julia fractals, allows to zoom in and out, pan and pause the animation.

## How to run
```bash
go run examples/fractal-demo/fractal-demo.go [-mode mode]
```

Screen mode, one of `1x1`, `1x2`, `2x2`, `2x3`, `2x4Braille` (default `2x3`).

## Controls

   - `SPACE` - switch fractal type (Mandelbrot/Julia)
   - `+` or `=`, `-` - zoom in, zoom out
   - `ARROWS` - pan the fractal
   - `p` - pause/resume animation
   - `q` - quit

## Screenshots

<img width="623" alt="Ghostty terminal with fractal demo" src="https://github.com/user-attachments/assets/d27c3b3b-edbc-4b89-9c5f-5c27356f8056" />
<img width="623" alt="Ghostty terminal with fractal demo" src="https://github.com/user-attachments/assets/74c34b2a-f646-4683-b018-68bd07c8bd7c" />
