package tcg

import (
	"errors"
	"fmt"
	"strconv"
)

// PixelMode - graphics mode, size and symbols for graphics
type PixelMode struct {
	width, height int    // dimension of one character in pixels
	charMapping   []rune // charMapping used to render the pixels on screen
}

var (
	// predefined modes with 1x1, 1x2, 2x2 and 2x3 pixels per character
	Mode1x1 = PixelMode{width: 1, height: 1, charMapping: pixelChars1x1}
	Mode1x2 = PixelMode{width: 1, height: 2, charMapping: pixelChars1x2}
	Mode2x2 = PixelMode{width: 2, height: 2, charMapping: pixelChars2x2}
	Mode2x3 = PixelMode{width: 2, height: 3, charMapping: pixelChars2x3}

	// Simple mode for debug or unittests with "." and "*" chars as colors
	Mode1x1Simple = PixelMode{width: 1, height: 1, charMapping: pixelChars1x1Simple}
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
		*pm = Mode1x1
	case "1x2":
		*pm = Mode1x2
	case "2x2":
		*pm = Mode2x2
	case "2x3":
		*pm = Mode2x3
	default:
		return errors.New("not valid screen mode")
	}

	return nil
}

// NewPixelMode - create new custom mode,
// charMapping slice must contain all combinations of pixels that the symbol can show.
// For example for 2x3 mode you need provide 64 symbols: 2^(2*3), for 3x3: 512
func NewPixelMode(width, height int, charMapping []rune) (*PixelMode, error) {
	if err := checkCharMapping(width, height, charMapping); err != nil {
		return nil, err
	}

	return &PixelMode{width: width, height: height, charMapping: charMapping}, nil
}

func checkCharMapping(width, height int, charMapping []rune) error {
	if width == 0 || height == 0 {
		return fmt.Errorf("zero width or height")
	}

	if len(charMapping) == 0 {
		return fmt.Errorf("empty char mapping list")
	}

	if 1<<(width*height) != len(charMapping) {
		return fmt.Errorf("char list length: %d not equal %d (2 ^ (%d * %d))", len(charMapping), 1<<(width*height), width, height)
	}

	return nil
}
