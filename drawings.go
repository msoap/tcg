package tcg

// HLine - draw horizontal line
func (tg *Tcg) HLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		tg.Buffer.Set(x+i, y, color)
	}
}

// VLine - draw vertical line
func (tg *Tcg) VLine(x, y int, length int, color int) {
	for i := 0; i < length; i++ {
		tg.Buffer.Set(x, y+i, color)
	}
}

// Box - draw box line
func (tg *Tcg) Box(x, y int, width, height int, color int) {
	tg.HLine(x, y, width, color)
	tg.HLine(x, y+height-1, width, color)
	tg.VLine(x, y, height, color)
	tg.VLine(x+width-1, y, height, color)
}
