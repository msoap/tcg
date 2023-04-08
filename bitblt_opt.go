package tcg

// Options for BitBlt
type BitBltOpt func(*bitBltOptions)

type bitBltOptions struct {
	// if true, then BitBlt will not copy pixels with color 0 (transparent)
	transparent bool
	// mask - if not nil, then BitBlt will copy only pixels with color != 0 in mask
	mask *Buffer
	// list of operations for each pixel
	operations []func(orig, src int) int
}

// BBTransparent - if true, then BitBlt will not copy pixels with color 0 (transparent)
func BBTransparent() BitBltOpt {
	return func(o *bitBltOptions) {
		o.transparent = true
	}
}

// BBMask - copy only pixels with color != 0 in mask
func BBMask(mask *Buffer) BitBltOpt {
	return func(o *bitBltOptions) {
		o.mask = mask
	}
}

// BBOpFn - add function for each pixel of source and destination
func BBOpFn(fn func(orig, src int) int) BitBltOpt {
	return func(o *bitBltOptions) {
		o.operations = append(o.operations, fn)
	}
}

// BBAnd - AND operation for each pixel of source and destination
func BBAnd() BitBltOpt {
	return func(o *bitBltOptions) {
		o.operations = append(o.operations, func(orig, src int) int {
			return orig & src
		})
	}
}

// BBOr - OR operation for each pixel of source and destination
func BBOr() BitBltOpt {
	return func(o *bitBltOptions) {
		o.operations = append(o.operations, func(orig, src int) int {
			return orig | src
		})
	}
}

// BBXor - XOR operation for each pixel of source and destination
func BBXor() BitBltOpt {
	return func(o *bitBltOptions) {
		o.operations = append(o.operations, func(orig, src int) int {
			return orig ^ src
		})
	}
}
