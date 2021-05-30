package tcg

// FillOpt - fill options
type FillOpt func(*fillOptions)

type fillOptions struct {
	pattern  *Buffer
	checkBuf *Buffer // buffer for check where we already fill pixels
	mask     *Buffer // mask buffer used for fill instead of original buffer
}

// WithPattern - option for Fill method, which provide fill pattern from another buffer
func WithPattern(buf Buffer) FillOpt {
	return func(fo *fillOptions) {
		fo.pattern = &buf
	}
}

// WithMask - option for Fill method: add mask
func WithMask(buf Buffer) FillOpt {
	return func(fo *fillOptions) {
		fo.mask = &buf
	}
}
