package tcg

// HLine - draw horizontal line
func (b *Buffer) HLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		b.Set(x+i, y, color)
	}
}

// VLine - draw vertical line
func (b *Buffer) VLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		b.Set(x, y+i, color)
	}
}

// Box - draw box line
func (b *Buffer) Box(x, y int, width, height int, color int) {
	b.HLine(x, y, width, color)
	b.HLine(x, y+height-1, width, color)
	b.VLine(x, y, height, color)
	b.VLine(x+width-1, y, height, color)
}
