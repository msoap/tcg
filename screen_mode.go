package tcg

import (
	"errors"
	"fmt"
	"strconv"
)

// PixelMode - graphics mode, size and symbols for graphics
type PixelMode struct {
	width, height int    // dimension of one character in pixels
	charMapping   []rune // chatMap used to render the pixels on screen
}

var (
	Mode1x1v = PixelMode{width: 1, height: 1, charMapping: pixelChars1x1}
	Mode1x2v = PixelMode{width: 1, height: 2, charMapping: pixelChars1x2}
	Mode2x2v = PixelMode{width: 2, height: 2, charMapping: pixelChars2x2}
	Mode2x3v = PixelMode{width: 2, height: 3, charMapping: pixelChars2x3}
)

// Width - returns the width in pixels of one character in the text console
func (pm PixelMode) Width() int {
	return pm.width
}

// Height - returns the height in pixels of one character in the text console
func (pm PixelMode) Height() int {
	return pm.height
}

// String representation like "2x3"
func (pm PixelMode) String() string {
	return strconv.Itoa(pm.width) + "x" + strconv.Itoa(pm.height)
}

// Set from string (for using with flag.Var())
func (pm *PixelMode) Set(in string) error {
	switch in {
	case "1x1":
		*pm = Mode1x1v
	case "1x2":
		*pm = Mode1x2v
	case "2x2":
		*pm = Mode2x2v
	case "2x3":
		*pm = Mode2x3v
	default:
		return errors.New("not valid screen mode")
	}

	return nil
}

func NewPixelMode(width, height int, charMapping []rune) (*PixelMode, error) {
	if err := checkCharMapping(width, height, charMapping); err != nil {
		return nil, err
	}

	return &PixelMode{width: width, height: height, charMapping: charMapping}, nil
}

func checkCharMapping(width, height int, cm []rune) error {
	if width == 0 || height == 0 {
		return fmt.Errorf("zero width or height")
	}

	if len(cm) == 0 {
		return fmt.Errorf("empty char mapping list")
	}

	if 1<<(width*height) != len(cm) {
		return fmt.Errorf("char list length: %d not equal %d (2 ^ (%d * %d))", len(cm), 1<<(width*height), width, height)
	}

	return nil
}

// PixelsInChar - a type representing the graphics mode
type PixelsInChar int

// graphics modes
const (
	Mode1x1 PixelsInChar = iota
	Mode1x2
	Mode2x2
	Mode2x3
)

// Width - returns the width in pixels of one character in the text console
func (pic PixelsInChar) Width() int {
	switch pic {
	case Mode1x1, Mode1x2:
		return 1
	case Mode2x2, Mode2x3:
		return 2
	default:
		return 0
	}
}

// Height - returns the height in pixels of one character in the text console
func (pic PixelsInChar) Height() int {
	switch pic {
	case Mode1x1:
		return 1
	case Mode1x2, Mode2x2:
		return 2
	case Mode2x3:
		return 3
	default:
		return 0
	}
}

// String represent
func (pic PixelsInChar) String() string {
	switch pic {
	case Mode1x1:
		return "1x1"
	case Mode1x2:
		return "1x2"
	case Mode2x2:
		return "2x2"
	case Mode2x3:
		return "2x3"
	default:
		return "-"
	}
}

// Set from string (for using with flag.Var())
func (pic *PixelsInChar) Set(in string) error {
	switch in {
	case "1x1":
		*pic = Mode1x1
	case "1x2":
		*pic = Mode1x2
	case "2x2":
		*pic = Mode2x2
	case "2x3":
		*pic = Mode2x3
	default:
		return errors.New("not valid screen mode")
	}

	return nil
}
