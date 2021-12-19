package tcg

// FillOpt - fill options
type FillOpt func(*fillOptions)

type fillOptions struct {
	pattern  *Buffer // buffer with pattern for fill instead of black color
	checkBuf *Buffer // buffer for check where we already fill pixels
	mask     *Buffer // mask buffer used for fill instead of original buffer
	allAreas bool    // fill in all areas, not necessarily continuous
}

// WithPattern - option for Fill method, which provide fill pattern from another buffer
func WithPattern(buf Buffer) FillOpt {
	return func(fo *fillOptions) {
		fo.pattern = &buf
	}
}

// WithMask - option for Fill method: add mask from Buffer
func WithMask(buf Buffer) FillOpt {
	return func(fo *fillOptions) {
		fo.mask = &buf
	}
}

// WithAllAreas - option for Fill method: fill in all areas, not necessarily continuous.
// Makes sense only when filled with a pattern.
func WithAllAreas() FillOpt {
	return func(fo *fillOptions) {
		fo.allAreas = true
	}
}
