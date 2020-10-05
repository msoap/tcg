package tcg

import "math"

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

func allocateBuffer(w, h int) [][]byte {
	buffer := make([][]byte, h)
	bytesLen := int(math.Ceil(float64(w) / 8))
	for i := range buffer {
		buffer[i] = make([]byte, bytesLen)
	}

	return buffer
}

// PutPixel - put pixel into buffer
func (b *Buffer) PutPixel(x, y int, color int) {
	if x < 0 || x > b.Width-1 || y < 0 || y > b.Height-1 {
		return
	}

	// [0][1][2][3] [4][5][6][7]
	i, mask := x/8, byte(0x80>>byte(x%8))
	switch color {
	case Black:
		b.buffer[y][i] = b.buffer[y][i] | mask
	case White:
		b.buffer[y][i] = b.buffer[y][i] &^ mask
	}
}

// GetPixel - get pixel from buffer
func (b Buffer) GetPixel(x, y int) int {
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
