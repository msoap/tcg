package tcg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// Opt - options type for New tcg screen
type Opt func(*tcgConfig) error

type tcgConfig struct {
	width, height int  // screen size in characters
	clip          geom // clip, width == 0 - without clip
	style         tcell.Style
}

type geom struct {
	x, y, width, height int
}

// WithClip - set clip of screen,
// x, y, w, h - is in screen character coordinates, not pixels
func WithClip(x, y, width, height int) Opt {
	return func(cfg *tcgConfig) error {
		if width < 1 || height < 1 {
			return fmt.Errorf("width (%d) or height (%d) is less than 1 x 1", width, height)
		}
		if x+width > cfg.width || y+height > cfg.height {
			return fmt.Errorf("clip size (%d, %d; %d x %d) does not fit in screen size (%d x %d)", x, y, width, height, cfg.width, cfg.height)
		}

		cfg.clip = geom{
			x:      x,
			y:      y,
			width:  width,
			height: height,
		}
		return nil
	}
}

// WithClipCenter - set clip of screen, placed in the center of screen
// w, h - is in screen character coordinates, not pixels
func WithClipCenter(width, height int) Opt {
	return func(cfg *tcgConfig) error {
		return WithClip((cfg.width-width)/2, (cfg.height-height)/2, width, height)(cfg)
	}
}

// WithColor - set default color of pixels, this will affect the block of pixels per full symbol
// color can be in the form of different options: "blue", "yellow" or "#ffaa11", see tcell.GetColor
func WithColor(colorName string) Opt {
	return func(cfg *tcgConfig) error {
		cfg.style = cfg.style.Foreground(tcell.GetColor(colorName))
		return nil
	}
}

// WithBackgroundColor - set default background color of pixels, this will affect the block of pixels per full symbol
// color can be in the form of different options: "blue", "yellow" or "#ffaa11", see tcell.GetColor
func WithBackgroundColor(colorName string) Opt {
	return func(cfg *tcgConfig) error {
		cfg.style = cfg.style.Background(tcell.GetColor(colorName))
		return nil
	}
}

// ParseSizeString - parse size in "80x25" format
func ParseSizeString(in string) (int, int, error) {
	parts := strings.SplitN(in, "x", 2)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("size not in 'd x d' format: %s, %v", in, parts)
	}

	w, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, err
	}

	h, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, err
	}

	if w < 1 || h < 1 {
		return 0, 0, fmt.Errorf("width (%d) or height (%d) less than 1", w, h)
	}

	return w, h, nil
}
