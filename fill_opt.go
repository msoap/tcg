package tcg

// FillOpt - fill options
type FillOpt func(*fillOptions)

type fillOptions struct {
	pattern *Buffer
	fillBuf *Buffer // buffer for check where we already fill pixels
}

// WithPattern - option for Fill method, which provide fill pattern from another buffer
func WithPattern(buf Buffer) FillOpt {
	return func(fo *fillOptions) {
		fo.pattern = &buf
	}
}
