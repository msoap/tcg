package main

import (
	"flag"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/msoap/tcg"
)

const (
	maxIter = 100
)

// getColorWithDithering returns a dithered color based on iteration count
func getColorWithDithering(iter int, maxIter int, x, y int) int {
	if iter >= maxIter {
		return tcg.White
	}

	// Create multiple color bands based on iteration count
	// This creates more visible detail across the fractal
	ratio := float64(iter) / float64(maxIter)

	// Use a non-linear mapping to enhance contrast
	// Square root makes lower iteration counts more visible
	intensity := math.Sqrt(ratio)

	// Simple dithering pattern - much more permissive
	ditherPattern := (x + y) % 3
	threshold := float64(ditherPattern) / 3.0

	// Create multiple threshold levels for better visibility
	if intensity > 0.8 {
		return tcg.Black // Darkest areas
	} else if intensity > 0.6 {
		if threshold < 0.5 {
			return tcg.Black
		}
		return tcg.White
	} else if intensity > 0.4 {
		if threshold < 0.33 {
			return tcg.Black
		}
		return tcg.White
	} else if intensity > 0.2 {
		if threshold < 0.2 {
			return tcg.Black
		}
		return tcg.White
	} else {
		// Very light areas - mostly white with sparse black dots
		if (x+y*2)%7 == 0 {
			return tcg.Black
		}
		return tcg.White
	}
}

// mandelbrot calculates if a point belongs to the Mandelbrot set
func mandelbrot(c complex128, maxIter int) int {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return i
		}
		z = z*z + c
	}
	return maxIter
}

// drawMandelbrot draws the Mandelbrot fractal on the buffer
func drawMandelbrot(buf *tcg.Buffer, zoom float64, offsetX, offsetY float64) {
	buf.Clear()
	height := buf.Height
	width := buf.Width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Convert pixel coordinates to complex plane
			realPart := (float64(x-width/2))/zoom + offsetX
			imagPart := (float64(y-height/2))/zoom + offsetY

			c := complex(realPart, imagPart)
			iter := mandelbrot(c, maxIter)

			// Use dithered color based on iteration count
			color := getColorWithDithering(iter, maxIter, x, y)
			buf.Set(x, y, color)
		}
	}
}

// drawJulia draws a Julia set fractal
func drawJulia(buf *tcg.Buffer, zoom float64, c complex128, offsetX, offsetY float64) {
	buf.Clear()
	height := buf.Height
	width := buf.Width

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Convert pixel coordinates to complex plane
			realPart := (float64(x-width/2))/zoom + offsetX
			imagPart := (float64(y-height/2))/zoom + offsetY

			z := complex(realPart, imagPart)
			iter := 0

			for iter < maxIter {
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					break
				}
				z = z*z + c
				iter++
			}

			// Use dithered color based on iteration count
			color := getColorWithDithering(iter, maxIter, x, y)
			buf.Set(x, y, color)
		}
	}
}

func main() {
	mode := tcg.Mode2x3
	flag.Var(&mode, "mode", "screen mode, one of 1x1, 1x2, 2x2, 2x3, 2x4Braille")
	flag.Parse()

	tg, err := tcg.New(mode)
	if err != nil {
		log.Fatal(err)
	}
	defer tg.Finish()

	width, height := tg.Buf.Width, tg.Buf.Height
	buf := tcg.NewBuffer(width, height-6)

	zoom := 50.0
	offsetX := -0.7
	offsetY := 0.0
	offsetStepScale := 0.05
	animationStep := 0
	fractalType := 0 // 0 for Mandelbrot, 1 for Julia
	startedFrameAt := time.Now()
	isRunning := true

	for {
		select {
		case ev := <-tg.TCellScreen.EventQ():
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {

				case tcell.KeyEscape, tcell.KeyCtrlC:
					return

				case tcell.KeyLeft:
					offsetX -= offsetStepScale
				case tcell.KeyRight:
					offsetX += offsetStepScale
				case tcell.KeyUp:
					offsetY -= offsetStepScale
				case tcell.KeyDown:
					offsetY += offsetStepScale

				case tcell.KeyRune:
					switch ev.Str() {
					case "q":
						return
					case "p": // Pause/resume animation
						isRunning = !isRunning
					case " ": // Space to switch fractal type
						fractalType = (fractalType + 1) % 2
					case "+", "=": // Zoom in
						zoom *= 1.1
					case "-": // Zoom out
						zoom /= 1.1
					}
				}
			}
		default:
		}

		animOffsetX := offsetX + 0.1*math.Sin(float64(animationStep)*0.01)
		animOffsetY := offsetY + 0.05*math.Cos(float64(animationStep)*0.015)

		if fractalType == 0 {
			animZoom := zoom * (1.0 + 0.01*math.Sin(float64(animationStep)*0.02))

			drawMandelbrot(&buf, animZoom, animOffsetX, animOffsetY)
		} else {
			juliaReal := -0.7 + 0.3*math.Sin(float64(animationStep)*0.03)
			juliaImag := 0.27015 + 0.1*math.Cos(float64(animationStep)*0.02)
			juliaC := complex(juliaReal, juliaImag)

			drawJulia(&buf, zoom*0.8, juliaC, animOffsetX, animOffsetY)
		}

		tg.Buf.BitBltAll(0, 5, buf)

		// some decorations
		tg.Buf.Rect(0, 0, width, height, tcg.Black)
		tg.Buf.HLine(1, 2, width-2, tcg.Black)
		tg.Buf.HLine(1, 4, width-2, tcg.Black)

		tg.Show(func(tg *tcg.Tcg) {
			fps := strconv.Itoa(int(1.0 / time.Since(startedFrameAt).Seconds()))
			tg.PrintStr(2, 0, "   Fractal demo. Press SPACE - switch kind, +/- - zoom, arrows - pan, P - pause, Q - quit [ "+fps+" FPS ]   ")
		})

		time.Sleep(80 * time.Millisecond)
		startedFrameAt = time.Now()
		if isRunning {
			animationStep++
		}
	}
}
