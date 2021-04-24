package tcg

import (
	"fmt"
	"math"
)

// Buffer - implement base screen pixel buffer
type Buffer struct {
	Width  int
	Height int
	buffer [][]byte
}

// NewBuffer - get new buffer object
func NewBuffer(width, height int) Buffer {
	return Buffer{
		Width:  width,
		Height: height,
		buffer: allocateBuffer(width, height),
	}
}

// NewBufferFromStrings - get new buffer object from list of strings (00110001, 0001001, ...)
// such symbols are also valid: "  **  ", "..##.."
func NewBufferFromStrings(in []string) (*Buffer, error) {
	if len(in) == 0 {
		return nil, fmt.Errorf("got empty string list")
	}

	width, height := len(in[0]), len(in)

	buf := Buffer{
		Width:  width,
		Height: height,
		buffer: allocateBuffer(width, height),
	}

	for y, line := range in {
		if len(line) != width {
			return nil, fmt.Errorf("got line with different width (%d != %d) on line %d", width, len(line), y)
		}
		for x, char := range line {
			switch char {
			case '0', '.', ' ':
				// pass
			case '1', '*', '#':
				buf.Set(x, y, Black)
			default:
				return nil, fmt.Errorf("got not valid char %v on %d:%d", char, x, y)
			}
		}
	}

	return &buf, nil
}

// Strings - render as slice of string, ["00011000", ...]
func (b Buffer) Strings() []string {
	result := make([]string, 0, b.Height)
	for y := 0; y < b.Height; y++ {
		line := ""
		for x := 0; x < b.Width; x++ {
			if b.At(x, y) == Black {
				line += "1"
			} else {
				line += "0"
			}
		}
		result = append(result, line)
	}

	return result
}

// MustNewBufferFromStrings - get new buffer object from list of strings and die on error
func MustNewBufferFromStrings(in []string) Buffer {
	buf, err := NewBufferFromStrings(in)
	if err != nil {
		panic(err)
	}
	return *buf
}

func allocateBuffer(w, h int) [][]byte {
	buffer := make([][]byte, h)
	bytesLen := int(math.Ceil(float64(w) / 8))
	for i := range buffer {
		buffer[i] = make([]byte, bytesLen)
	}

	return buffer
}

// Set - put pixel into buffer
func (b *Buffer) Set(x, y int, color int) {
	if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
		return
	}

	// [0][1][2][3] [4][5][6][7]
	i, mask := x/8, byte(0x80>>byte(x%8))
	switch color {
	case Black:
		b.buffer[y][i] |= mask
	case White:
		b.buffer[y][i] &^= mask
	}
}

// At - get pixel from buffer
func (b Buffer) At(x, y int) int {
	if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
		return White
	}

	// [0][1][2][3] [4][5][6][7]
	i, mask := x/8, byte(0x80>>byte(x%8))

	if b.buffer[y][i]&mask > 0 {
		return Black
	}

	return White
}

// getPixelsBlock - get rectangular block of pixels as linear bits
func (b Buffer) getPixelsBlock(x, y, width, height int) int {
	if x < 0 || x+width > b.Width || y < 0 || y+height > b.Height {
		return 0
	}

	result, num := 0, width*height-1

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			result |= b.At(x+w, y+h) << num
			num--
		}
	}

	return result
}

// BitBltAllSrc - copy whole buffer into this buffer
func (b *Buffer) BitBltAllSrc(x, y int, from Buffer) {
	if x == 0 && y == 0 {
		for i := 0; i < from.Height && i < b.Height; i++ {
			copy(b.buffer[i], from.buffer[i])
		}
		return
	}

	for i := 0; i < from.Height && i+y < b.Height; i++ {
		for j := 0; j < from.Width && j+x < b.Width; j++ {
			b.Set(j+x, i+y, from.At(j, i))
		}
	}
}

// BitBlt - copy part of buffer into this buffer
func (b *Buffer) BitBlt(xd, yd, width, height int, from Buffer, xs, ys int) {
	for i := 0; i+ys < from.Height && i < height && i+yd < b.Height; i++ {
		for j := 0; j+xs < from.Width && j < width && j+xd < b.Width; j++ {
			b.Set(j+xd, i+yd, from.At(j+xs, i+ys))
		}
	}
}
