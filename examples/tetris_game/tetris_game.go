// Tetris - classic Tetris game in a terminal, an example of using the tcg library.
//
// The game area is drawn as pixels in the tcg buffer, so blocks fall smoothly pixel by pixel.
// TCG can use only one color per terminal character, so after the pixels are rendered
// the characters of each block are repainted with the color of the block (see Tcg.Show finalisers).
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	tcell "github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/msoap/tcg"
)

const (
	// size of one square in characters, so in pixels it depends on the screen mode:
	// 4x3 pixels in the 2x3 mode, 4x4 pixels in the 2x4Braille mode
	squareCharsW = 2
	squareCharsH = 1

	fieldCols = 12 // width of the game area in squares
	fieldRows = 25 // height of the game area in squares

	fallSpeed     = 4.0 // initial speed of a falling block in pixels per second
	speedUpFactor = 1.2 // speed multiplier for each next level

	linesPerLevel  = 10 // lines to clear to get the next level
	lineScore      = 10 // score for a cleared line on the first level
	levelLineBonus = 5  // additional score for a cleared line for each level above the first

	// animation of a cleared line: each step erases one more pixel row of the line (top to bottom),
	// the last step drops everything above the line, the game continues during the animation
	clearStepTime = 30 * time.Millisecond

	frameRate    = 60                     // screen updates per second
	maxFrameTime = 100 * time.Millisecond // limit of one time step, for example after the terminal was suspended

	panelGap   = 2  // gap between the game area and the info panel in characters
	panelWidth = 16 // width of the info panel in characters
)

// sizes in characters
const (
	fieldCharsW = fieldCols * squareCharsW
	fieldCharsH = fieldRows * squareCharsH
	layoutW     = fieldCharsW + 2 + panelGap + panelWidth // game area with borders and the panel
	layoutH     = fieldCharsH + 2
)

var (
	minScreenW = max(layoutW, bannerWidth(helpLines))
	minScreenH = max(layoutH, len(helpLines)+2)
)

// sizes in pixels, they depend on the screen mode, see setMode
var (
	screenMode       tcg.PixelMode
	modeW, modeH     int // size of a character
	squareW, squareH int // size of a square
)

func init() {
	setMode(tcg.Mode2x3)
}

// setMode - set the screen mode of the game, it must be called before the start of the game
func setMode(mode tcg.PixelMode) {
	screenMode = mode
	modeW, modeH = mode.Width(), mode.Height()
	squareW, squareH = squareCharsW*modeW, squareCharsH*modeH
}

// modeFlag - screen mode as a command line flag, only modes with enough pixels per character are suitable for the game
type modeFlag struct{ tcg.PixelMode }

func (m *modeFlag) Set(in string) error {
	if in != "2x3" && in != "2x4Braille" {
		return fmt.Errorf("unsupported mode %q, use 2x3 or 2x4Braille", in)
	}

	return m.PixelMode.Set(in)
}

var helpLines = []string{
	"←  →    move",
	"↑       rotate",
	"space   drop",
	"p       pause",
	"r       restart",
	"q       quit",
	"?       show/hide help",
}

var (
	baseStyle    = tcell.StyleDefault.Foreground(color.Default).Background(color.Default)
	captionStyle = baseStyle.Foreground(color.Silver) // frame and captions of the panel
)

type point struct{ x, y int }

// shape of a block, it's rotated inside the square box size x size
type shape struct {
	size  int
	cells [4]point
	color color.Color
}

// all kinds of blocks, as drawn in the rules
var shapes = []shape{
	// I
	newShape(color.Yellow,
		"....",
		"****",
		"....",
		"...."),
	// O
	newShape(color.Red,
		"**",
		"**"),
	// T
	newShape(color.Lime,
		"***",
		".*.",
		"..."),
	// S
	newShape(color.Aqua,
		"**.",
		".**",
		"..."),
	// Z
	newShape(color.Fuchsia,
		".**",
		"**.",
		"..."),
	// J
	newShape(color.Blue,
		"*..",
		"***",
		"..."),
	// L
	newShape(color.White,
		"..*",
		"***",
		"..."),
}

func newShape(c color.Color, rows ...string) shape {
	s := shape{size: len(rows), color: c}
	n := 0
	for y, row := range rows {
		for x, char := range row {
			if char == '*' {
				s.cells[n] = point{x, y}
				n++
			}
		}
	}
	if n != len(s.cells) {
		panic(fmt.Sprintf("shape must have %d squares, got %d", len(s.cells), n))
	}

	return s
}

// piece - falling block
type piece struct {
	kind  int      // index in shapes
	size  int      // size of the box with squares
	cells [4]point // squares of the block in the box, after rotations
	col   int      // x of the box in squares
	y     int      // y of the box in pixels, the block moves down pixel by pixel
}

// newPiece - get block at the top center of the game area
func newPiece(kind int) piece {
	s := shapes[kind]
	p := piece{kind: kind, size: s.size, cells: s.cells, col: (fieldCols - s.size) / 2}
	p.y = -p.minY() * squareH
	return p
}

func (p piece) minY() int {
	result := p.cells[0].y
	for _, c := range p.cells {
		result = min(result, c.y)
	}
	return result
}

// rotated - get the block rotated clockwise
func (p piece) rotated() piece {
	for i, c := range p.cells {
		p.cells[i] = point{p.size - 1 - c.y, c.x}
	}
	return p
}

// game - state of the game without any drawing
type game struct {
	field               [fieldRows][fieldCols]uint8 // fallen squares: 0 - empty, otherwise shape index + 1
	cur                 piece
	next                int     // kind of the next block
	fall                float64 // part of the pixel the block has already fallen
	score, lines, level int
	over                bool

	clearing  []int         // full rows that are being removed, they are solid until the end of the animation
	clearStep int           // steps of the animation, equals to the number of erased pixel rows of the lines
	clearTime time.Duration // time from the last step of the animation
}

func newGame() *game {
	g := &game{level: 1, next: rand.IntN(len(shapes))}
	g.spawn()
	return g
}

func (g *game) spawn() {
	g.cur = newPiece(g.next)
	g.next = rand.IntN(len(shapes))
	g.fall = 0
	if g.collides(g.cur) {
		g.finishClearing() // the lines that are being removed can be the only obstacle
		g.startClearing()
		g.over = g.collides(g.cur)
	}
}

func floorDiv(a, b int) int {
	q := a / b
	if a%b != 0 && a < 0 {
		q--
	}
	return q
}

// collides - does the block overlap the fallen squares, walls, or floor?
func (g *game) collides(p piece) bool {
	for _, c := range p.cells {
		col, top := p.col+c.x, p.y+c.y*squareH
		if col < 0 || col >= fieldCols {
			return true
		}
		// a square between pixel rows can overlap two rows of the field
		for row := floorDiv(top, squareH); row*squareH < top+squareH; row++ {
			if row < 0 || row >= fieldRows || g.field[row][col] != 0 {
				return true
			}
		}
	}

	return false
}

// move - move the block by one square left or right
func (g *game) move(dx int) bool {
	p := g.cur
	p.col += dx
	if g.collides(p) {
		return false
	}
	g.cur = p

	return true
}

// rotate - rotate the block clockwise, it can be shifted aside to fit near walls and other squares
func (g *game) rotate() bool {
	p := g.cur.rotated()
	if top := p.y + p.minY()*squareH; top < 0 {
		p.y -= top // keep the block inside the game area
	}

	for _, dx := range []int{0, -1, 1, -2, 2} {
		shifted := p
		shifted.col += dx
		if !g.collides(shifted) {
			g.cur = shifted
			return true
		}
	}

	return false
}

// hardDrop - move the block down until it lands
func (g *game) hardDrop() {
	for g.step() {
	}
	g.lock()
}

// step - move the block down by one pixel
func (g *game) step() bool {
	p := g.cur
	p.y++
	if g.collides(p) {
		return false
	}
	g.cur = p

	return true
}

// speed of falling in pixels per second
func (g *game) speed() float64 {
	return fallSpeed * math.Pow(speedUpFactor, float64(g.level-1))
}

// update - move the block down according to the elapsed time, returns true if something was changed
func (g *game) update(dt time.Duration) bool {
	if g.over {
		return false
	}

	changed := g.animate(dt)

	g.fall += g.speed() * dt.Seconds()
	pixels := int(g.fall)
	g.fall -= float64(pixels)

	for range pixels {
		if !g.step() {
			g.lock()
			return true
		}
		changed = true
	}

	return changed
}

// lock - add the block to the fallen squares, start removing of full lines and the next block
func (g *game) lock() {
	for _, c := range g.cur.cells {
		// the block always stops aligned to the squares grid, it can't be stopped between pixel rows
		row := floorDiv(g.cur.y+c.y*squareH, squareH)
		g.field[row][g.cur.col+c.x] = uint8(g.cur.kind + 1)
	}

	// the block can land while the previous lines are still being removed, it's above them, so it falls with the rest
	g.finishClearing()
	g.startClearing()

	g.spawn()
}

// startClearing - start removing of the full lines and score them
func (g *game) startClearing() {
	rows := g.fullRows()
	if len(rows) == 0 {
		return
	}

	g.clearing = rows
	g.score += len(rows) * (lineScore + (g.level-1)*levelLineBonus)
	g.lines += len(rows)
	g.level = 1 + g.lines/linesPerLevel
}

// fullRows - get rows without empty squares, from top to bottom
func (g *game) fullRows() []int {
	var rows []int
	for row := range fieldRows {
		if !slices.Contains(g.field[row][:], 0) {
			rows = append(rows, row)
		}
	}

	return rows
}

// animate - make the next steps of the animation of removing lines, returns true if something was changed
func (g *game) animate(dt time.Duration) bool {
	if len(g.clearing) == 0 {
		return false
	}

	g.clearTime += dt
	steps := int(g.clearTime / clearStepTime)
	if steps == 0 {
		return false
	}
	g.clearTime -= time.Duration(steps) * clearStepTime

	g.clearStep += steps
	if g.clearStep > squareH { // all pixel rows are erased
		g.finishClearing()
		g.startClearing() // fallen squares can complete new lines
	}

	return true
}

// finishClearing - remove the lines, everything above falls down in one step,
// squares left without support fall down too
func (g *game) finishClearing() {
	for _, row := range g.clearing { // from top to bottom, so numbers of the rest of the lines are not changed
		copy(g.field[1:row+1], g.field[:row])
		g.field[0] = [fieldCols]uint8{}
	}
	g.clearing, g.clearStep, g.clearTime = nil, 0, 0

	g.settle()
}

// neighbors - squares that are connected to the square, diagonal ones are not
var neighbors = [...]point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

// settle - let the squares that hang in the air fall down.
// Squares that touch each other by sides are one group, the group stands if it's connected to the floor,
// otherwise it falls as a whole until it lands on something.
func (g *game) settle() {
	for g.dropFloating() {
	}
}

// dropFloating - drop every group that isn't connected to the floor, returns true if something was moved.
// A group can land on another floating group and then they fall together, so one pass may be not enough.
func (g *game) dropFloating() bool {
	var seen [fieldRows][fieldCols]bool
	for col := range fieldCols {
		g.collectGroup(&seen, point{col, fieldRows - 1}) // the squares that stand
	}

	// from bottom to top, so the lower groups fall first and the upper ones can land on them
	var floating [][]point
	for row := fieldRows - 1; row >= 0; row-- {
		for col := range fieldCols {
			if group := g.collectGroup(&seen, point{col, row}); group != nil {
				floating = append(floating, group)
			}
		}
	}

	moved := false
	for _, group := range floating {
		if g.dropGroup(group) {
			moved = true
		}
	}

	return moved
}

// collectGroup - get the group of squares that contains the square, squares that were seen before are skipped,
// returns nil if there is no new square
func (g *game) collectGroup(seen *[fieldRows][fieldCols]bool, start point) []point {
	if g.field[start.y][start.x] == 0 || seen[start.y][start.x] {
		return nil
	}

	seen[start.y][start.x] = true
	group := []point{start}
	for i := 0; i < len(group); i++ { // the group grows in the loop
		for _, d := range neighbors {
			n := point{group[i].x + d.x, group[i].y + d.y}
			if n.x < 0 || n.x >= fieldCols || n.y < 0 || n.y >= fieldRows || g.field[n.y][n.x] == 0 || seen[n.y][n.x] {
				continue
			}
			seen[n.y][n.x] = true
			group = append(group, n)
		}
	}

	return group
}

// dropGroup - move the group of squares down until it lands on something, returns true if it was moved
func (g *game) dropGroup(group []point) bool {
	squares := make([]uint8, len(group))
	for i, p := range group {
		squares[i] = g.field[p.y][p.x]
		g.field[p.y][p.x] = 0 // the group doesn't block itself
	}

	dist := fieldRows
	for _, p := range group {
		free := 0
		for p.y+free+1 < fieldRows && g.field[p.y+free+1][p.x] == 0 {
			free++
		}
		dist = min(dist, free)
	}

	for i, p := range group {
		g.field[p.y+dist][p.x] = squares[i]
	}

	return dist > 0
}

// erasedPixels - number of erased pixel rows (from the top) of the squares in the row of the field
func (g *game) erasedPixels(row int) int {
	if slices.Contains(g.clearing, row) {
		return g.clearStep
	}

	return 0
}

// cellShape - what is drawn in the character of the game area, 0 - nothing, otherwise shape index + 1
func (g *game) cellShape(cx, cy int) uint8 {
	if s := g.field[cy/squareCharsH][cx/squareCharsW]; s != 0 {
		return s
	}
	if g.over {
		return 0
	}

	// the character is a rectangle of pixels, find overlapping with the squares of the block
	px, py := cx*modeW, cy*modeH
	for _, c := range g.cur.cells {
		x, y := (g.cur.col+c.x)*squareW, g.cur.y+c.y*squareH
		if x < px+modeW && x+squareW > px && y < py+modeH && y+squareH > py {
			return uint8(g.cur.kind + 1)
		}
	}

	return 0
}

type app struct {
	tg                     *tcg.Tcg
	g                      *game
	paused, help, tooSmall bool
	dirty                  bool // the screen should be redrawn
	x, y                   int  // top left corner of the game on the screen in characters
	startW, startH         int  // screen size on start
}

func main() {
	mode := modeFlag{tcg.Mode2x3}
	flag.Var(&mode, "mode", "screen mode, one of 2x3, 2x4Braille")
	flag.Parse()
	setMode(mode.PixelMode)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	tg, err := tcg.New(screenMode)
	if err != nil {
		return err
	}
	defer tg.Finish()

	w, h := tg.ScreenSize()
	if w < minScreenW || h < minScreenH {
		return fmt.Errorf("screen size %dx%d is too small, need at least %dx%d", w, h, minScreenW, minScreenH)
	}

	a := &app{tg: tg, g: newGame(), startW: w, startH: h}
	if err := a.layout(); err != nil {
		return err
	}

	return a.loop()
}

// layout - place the game in the center of the screen
func (a *app) layout() error {
	w, h := a.tg.ScreenSize()
	a.tooSmall = w < minScreenW || h < minScreenH
	if a.tooSmall {
		return nil
	}

	// tcg checks a clip against the screen size on start, so do not go out of it if the screen is bigger now
	a.x = min((w-layoutW)/2, a.startW-layoutW)
	a.y = min((h-layoutH)/2, a.startH-layoutH)

	return a.tg.SetClip(a.x+1, a.y+1, fieldCharsW, fieldCharsH)
}

func (a *app) frozen() bool {
	return a.paused || a.help || a.tooSmall || a.g.over
}

func (a *app) restart() {
	a.g = newGame()
	a.paused, a.help = false, false
}

func (a *app) loop() error {
	ticker := time.NewTicker(time.Second / frameRate)
	defer ticker.Stop()

	events := a.tg.TCellScreen.EventQ()
	last := time.Now()
	a.dirty = true

	for {
		select {
		case ev, ok := <-events:
			if !ok {
				return nil
			}
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if a.handleKey(ev) {
					return nil
				}
			case *tcell.EventResize:
				a.tg.TCellScreen.Sync()
				if err := a.layout(); err != nil {
					return err
				}
				a.dirty = true
			}
		case now := <-ticker.C:
			dt := min(now.Sub(last), maxFrameTime)
			last = now
			if !a.frozen() && a.g.update(dt) {
				a.dirty = true
			}
		}

		if a.dirty {
			a.draw()
			a.dirty = false
		}
	}
}

// handleKey - returns true if the game should be finished
func (a *app) handleKey(ev *tcell.EventKey) bool {
	if !ev.Pressed() {
		return false
	}

	key, str := ev.Key(), strings.ToLower(ev.Str())
	if key == tcell.KeyCtrlC || (ev.Modifiers()&tcell.ModCtrl != 0 && str == "c") || str == "q" {
		return true
	}

	a.dirty = true
	switch {
	case str == "r":
		a.restart()
	case str == "?":
		a.help = !a.help
	case key == tcell.KeyEscape && a.help:
		a.help = false
	case str == "p":
		a.paused = !a.paused && !a.g.over
	case a.frozen():
		a.dirty = false
	case key == tcell.KeyLeft:
		a.g.move(-1)
	case key == tcell.KeyRight:
		a.g.move(1)
	case key == tcell.KeyUp:
		a.g.rotate()
	case key == tcell.KeySpace || str == " ":
		a.g.hardDrop()
	default:
		a.dirty = false
	}

	return false
}

// draw - redraw the whole screen, tcell updates only changed characters on the terminal, so it doesn't flicker
func (a *app) draw() {
	a.tg.TCellScreen.Fill(' ', baseStyle)

	if a.tooSmall {
		a.tg.PrintStrStyled(0, 0, fmt.Sprintf("Screen is too small, need %dx%d", minScreenW, minScreenH), baseStyle)
		a.tg.TCellScreen.Show()
		return
	}

	a.tg.Buf.Clear()
	for row := range fieldRows {
		erased := a.g.erasedPixels(row)
		for col := range fieldCols {
			if a.g.field[row][col] != 0 {
				a.tg.Buf.FillRect(col*squareW, row*squareH+erased, squareW, squareH-erased, tcg.Black)
			}
		}
	}
	if !a.g.over {
		for _, c := range a.g.cur.cells {
			a.tg.Buf.FillRect((a.g.cur.col+c.x)*squareW, a.g.cur.y+c.y*squareH, squareW, squareH, tcg.Black)
		}
	}

	a.tg.Show(a.drawDecorations)
}

// drawDecorations - draw colors of the blocks, borders, info panel and banners
func (a *app) drawDecorations(tg *tcg.Tcg) {
	for cy := range fieldCharsH {
		for cx := range fieldCharsW {
			if s := a.g.cellShape(cx, cy); s != 0 {
				x, y := a.x+1+cx, a.y+1+cy
				char, _, _ := tg.TCellScreen.Get(x, y)
				tg.PrintStrStyled(x, y, char, baseStyle.Foreground(shapes[s-1].color))
			}
		}
	}

	drawBox(tg, a.x, a.y, fieldCharsW+2, layoutH, "TETRIS", captionStyle)
	a.drawPanel(tg)

	switch {
	case a.help:
		a.drawBanner(tg, "CONTROLS", helpLines)
	case a.g.over:
		a.drawBanner(tg, "GAME OVER", []string{"Score: " + strconv.Itoa(a.g.score), "", "r - restart", "q - quit"})
	case a.paused:
		a.drawBanner(tg, "PAUSED", []string{"p - continue"})
	}
}

func (a *app) drawPanel(tg *tcg.Tcg) {
	x, y := a.x+fieldCharsW+2+panelGap, a.y+1
	labelStyle, valueStyle := captionStyle, baseStyle.Bold(true)

	stats := []struct {
		label string
		value int
	}{{"SCORE", a.g.score}, {"LEVEL", a.g.level}, {"LINES", a.g.lines}}
	for i, stat := range stats {
		tg.PrintStrStyled(x, y+i*3, stat.label, labelStyle)
		tg.PrintStrStyled(x, y+i*3+1, strconv.Itoa(stat.value), valueStyle)
	}

	tg.PrintStrStyled(x, y+9, "NEXT", labelStyle)
	drawShape(tg, x, y+11, a.g.next)

	tg.PrintStrStyled(x, y+fieldCharsH-1, "? - help", labelStyle)
}

// drawShape - draw the block with characters, centered in the width of 4 squares
func drawShape(tg *tcg.Tcg, x, y, kind int) {
	s := shapes[kind]
	minX, minY, maxX := s.cells[0].x, s.cells[0].y, s.cells[0].x
	for _, c := range s.cells {
		minX, minY, maxX = min(minX, c.x), min(minY, c.y), max(maxX, c.x)
	}

	x += (4 - (maxX - minX + 1)) * squareCharsW / 2
	style, square := baseStyle.Foreground(s.color), strings.Repeat("█", squareCharsW)
	for _, c := range s.cells {
		for row := range squareCharsH {
			tg.PrintStrStyled(x+(c.x-minX)*squareCharsW, y+(c.y-minY)*squareCharsH+row, square, style)
		}
	}
}

// drawBox - draw a frame with an optional title, the inside of the frame is not changed
func drawBox(tg *tcg.Tcg, x, y, w, h int, title string, style tcell.Style) {
	top := strings.Repeat("─", w-2)
	if title = " " + title + " "; utf8.RuneCountInString(title) <= w-2 {
		left := (w - 2 - utf8.RuneCountInString(title)) / 2
		top = strings.Repeat("─", left) + title + strings.Repeat("─", w-2-left-utf8.RuneCountInString(title))
	}

	tg.PrintStrStyled(x, y, "┌"+top+"┐", style)
	for row := 1; row < h-1; row++ {
		tg.PrintStrStyled(x, y+row, "│", style)
		tg.PrintStrStyled(x+w-1, y+row, "│", style)
	}
	tg.PrintStrStyled(x, y+h-1, "└"+strings.Repeat("─", w-2)+"┘", style)
}

func bannerWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		width = max(width, utf8.RuneCountInString(line))
	}

	return width + 4
}

// drawBanner - draw a text box in the center of the game
func (a *app) drawBanner(tg *tcg.Tcg, title string, lines []string) {
	w, h := max(bannerWidth(lines), len(title)+6), len(lines)+2
	x, y := a.x+(layoutW-w)/2, a.y+(layoutH-h)/2

	drawBox(tg, x, y, w, h, title, baseStyle)
	for i, line := range lines {
		tg.PrintStrStyled(x+1, y+1+i, fmt.Sprintf(" %-*s", w-3, line), baseStyle)
	}
}
