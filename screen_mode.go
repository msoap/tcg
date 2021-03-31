package tcg

import "errors"

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

// Set from string (for use with flag.Var())
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
