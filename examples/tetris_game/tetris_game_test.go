package main

import (
	"testing"
	"time"

	"github.com/msoap/tcg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestGame - game with the block of given kind on the empty field
func newTestGame(kind int) *game {
	g := newGame()
	g.field = [fieldRows][fieldCols]uint8{}
	g.next = kind
	g.spawn()
	return g
}

func cellSet(p piece) map[point]bool {
	result := map[point]bool{}
	for _, c := range p.cells {
		result[c] = true
	}
	return result
}

func Test_shapes(t *testing.T) {
	require.Len(t, shapes, 7)

	for kind, s := range shapes {
		p := newPiece(kind)
		assert.Len(t, cellSet(p), 4, "kind %d: squares must be different", kind)
		for _, c := range p.cells {
			assert.True(t, c.x >= 0 && c.x < s.size && c.y >= 0 && c.y < s.size, "kind %d: square %v is out of the box", kind, c)
		}

		// four rotations return the block to the start position
		r := p
		for range 4 {
			r = r.rotated()
		}
		assert.Equal(t, cellSet(p), cellSet(r), "kind %d", kind)

		// new block is at the top, in the center, and fits the empty field
		assert.Zero(t, p.y+p.minY()*squareH, "kind %d: top square is on the first row", kind)
		assert.False(t, newTestGame(kind).collides(p), "kind %d", kind)
	}
}

func Test_rotate_clockwise(t *testing.T) {
	g := newTestGame(5) // J
	assert.Equal(t, map[point]bool{{0, 0}: true, {0, 1}: true, {1, 1}: true, {2, 1}: true}, cellSet(g.cur))

	require.True(t, g.rotate())
	assert.Equal(t, map[point]bool{{2, 0}: true, {1, 0}: true, {1, 1}: true, {1, 2}: true}, cellSet(g.cur))
}

func Test_move_walls(t *testing.T) {
	g := newTestGame(0) // I, 4 squares wide
	for range fieldCols {
		g.move(-1)
	}
	assert.Equal(t, 0, g.cur.col, "block is at the left wall")
	assert.False(t, g.move(-1))

	for range fieldCols {
		g.move(1)
	}
	assert.Equal(t, fieldCols-4, g.cur.col, "block is at the right wall")
	assert.False(t, g.move(1))
}

func Test_rotate_near_wall(t *testing.T) {
	g := newTestGame(0) // I
	require.True(t, g.rotate())
	for range fieldCols {
		g.move(1)
	}
	// vertical block is at the right wall, horizontal one doesn't fit without moving aside
	require.True(t, g.rotate())
	assert.False(t, g.collides(g.cur))
}

func Test_rotate_on_top(t *testing.T) {
	g := newTestGame(0) // I, the box with a vertical block starts above the game area
	require.True(t, g.rotate())
	assert.False(t, g.collides(g.cur))
	assert.Zero(t, g.cur.y+g.cur.minY()*squareH)
}

func Test_collides_pixels(t *testing.T) {
	g := newTestGame(1) // O
	g.field[10][g.cur.col] = 1

	p := g.cur
	p.y = 8*squareH - 1 // the block is 1 pixel higher than rows 8-9, it's between rows 7-8 and 8-9
	assert.False(t, g.collides(p))

	p.y++ // the block is on the rows 8-9, it touches the square on the row 10
	assert.False(t, g.collides(p))

	p.y++ // the block overlaps the square by one pixel
	assert.True(t, g.collides(p))
}

func Test_gravity_pixels(t *testing.T) {
	g := newTestGame(1)
	y := g.cur.y

	assert.False(t, g.update(100*time.Millisecond), "0.4 pixel, nothing to draw")
	assert.Equal(t, y, g.cur.y)

	assert.True(t, g.update(200*time.Millisecond), "0.4 + 0.8 pixels")
	assert.Equal(t, y+1, g.cur.y)

	for range 4 {
		g.update(250 * time.Millisecond) // 1 pixel per step
	}
	assert.Equal(t, y+5, g.cur.y)
}

func Test_hardDrop(t *testing.T) {
	g := newTestGame(0) // I
	g.hardDrop()

	for col := range fieldCols {
		want := uint8(0)
		if col >= 4 && col < 8 {
			want = 1
		}
		assert.Equal(t, want, g.field[fieldRows-1][col], "col %d", col)
	}
	assert.Zero(t, g.score)
	assert.False(t, g.over)
}

func Test_fall_and_lock_on_stack(t *testing.T) {
	g := newTestGame(1) // O
	g.field[fieldRows-1][g.cur.col] = 7
	col := g.cur.col
	g.next = 0

	for i := 0; i < 100 && g.cur.kind == 1; i++ {
		g.update(time.Second)
	}

	require.Equal(t, 0, g.cur.kind, "the block was locked and the next one started")
	assert.Equal(t, uint8(2), g.field[fieldRows-2][col], "block landed on the top of the stack")
	assert.Equal(t, uint8(2), g.field[fieldRows-3][col])
	assert.Equal(t, uint8(2), g.field[fieldRows-2][col+1], "the right half is above the empty floor, but it's the same block")
	assert.Equal(t, uint8(2), g.field[fieldRows-3][col+1])
}

func Test_remove_lines(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	for row := fieldRows - 2; row < fieldRows; row++ {
		for col := range fieldCols {
			if col < 4 || col >= 8 {
				g.field[row][col] = 3
			}
		}
	}
	g.field[fieldRows-3][0] = 5 // above the lines

	g.hardDrop()

	// the I block is horizontal, so only the bottom line is full
	assert.Equal(t, []int{fieldRows - 1}, g.clearing)
	assert.Equal(t, 10, g.score)
	assert.Equal(t, 1, g.lines)
	assert.Equal(t, 1, g.level)

	g.finishClearing()
	assert.Equal(t, uint8(5), g.field[fieldRows-2][0], "everything above the line fell down")
	assert.Equal(t, uint8(3), g.field[fieldRows-1][0], "the second line moved to the bottom")
	assert.Equal(t, uint8(0), g.field[fieldRows-1][5])
	assert.Zero(t, g.field[0][0])
}

func Test_remove_several_lines(t *testing.T) {
	g := newTestGame(0) // I
	g.rotate()
	for _, row := range []int{fieldRows - 4, fieldRows - 3, fieldRows - 2, fieldRows - 1} {
		for col := range fieldCols {
			if col != 5 {
				g.field[row][col] = 3
			}
		}
	}
	g.field[fieldRows-5][0] = 4
	require.Equal(t, 6, g.cur.col+g.cur.cells[0].x, "vertical block is in the column 6")
	require.True(t, g.move(-1))

	g.hardDrop()

	assert.Equal(t, []int{fieldRows - 4, fieldRows - 3, fieldRows - 2, fieldRows - 1}, g.clearing)
	assert.Equal(t, 4*lineScore, g.score)
	assert.Equal(t, 4, g.lines)

	g.finishClearing()
	assert.Equal(t, uint8(4), g.field[fieldRows-1][0])
	for col := range fieldCols {
		assert.Zero(t, g.field[fieldRows-2][col])
	}
}

func Test_level_and_score(t *testing.T) {
	g := newTestGame(0)
	assert.InDelta(t, fallSpeed, g.speed(), 1e-9)

	g.lines = linesPerLevel - 1
	for col := range fieldCols {
		if col < 4 || col >= 8 {
			g.field[fieldRows-1][col] = 3
		}
	}
	g.hardDrop()

	assert.Equal(t, 2, g.level)
	assert.Equal(t, lineScore, g.score, "the line is on the old level")
	assert.InDelta(t, fallSpeed*speedUpFactor, g.speed(), 1e-9)

	// the second level gives a bonus for each line
	g.finishClearing()
	g.next = 0
	g.spawn()
	for col := range fieldCols {
		if col < 4 || col >= 8 {
			g.field[fieldRows-1][col] = 3
		}
	}
	g.hardDrop()
	assert.Equal(t, lineScore+lineScore+levelLineBonus, g.score)
	assert.Equal(t, linesPerLevel+1, g.lines)
}

func Test_game_over(t *testing.T) {
	g := newTestGame(1)
	assert.False(t, g.over)

	// the stack under the block is almost to the top
	for row := 2; row < fieldRows; row++ {
		g.field[row][g.cur.col], g.field[row][g.cur.col+1] = 3, 3
	}
	g.next = 1
	g.hardDrop()

	assert.True(t, g.over)
	assert.False(t, g.update(time.Second), "nothing changes after the end")
}

func Test_cellShape(t *testing.T) {
	g := newTestGame(1) // O
	g.field[fieldRows-1][0] = 4

	assert.Equal(t, uint8(4), g.cellShape(0, (fieldRows-1)*squareCharsH))
	assert.Zero(t, g.cellShape(fieldCharsW-1, 0))

	// the block is on the first row, in the center
	cx, cy := g.cur.col*squareCharsW, 0
	assert.Equal(t, uint8(2), g.cellShape(cx, cy))
	assert.Zero(t, g.cellShape(cx, 2*squareCharsH))

	// one pixel lower, the block takes a part of the next character row too
	g.cur.y++
	assert.Equal(t, uint8(2), g.cellShape(cx, 2*squareCharsH))
	g.cur.y += squareH
	assert.Zero(t, g.cellShape(cx, 0))
}

// fillRowExceptSlot - fill the row except the place for the horizontal I block
func fillRowExceptSlot(g *game, row int) {
	for col := range fieldCols {
		if col < 4 || col >= 8 {
			g.field[row][col] = 3
		}
	}
}

func Test_clear_animation(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	fillRowExceptSlot(g, fieldRows-1)
	g.hardDrop()

	assert.Equal(t, []int{fieldRows - 1}, g.clearing)
	assert.Equal(t, lineScore, g.score, "the score is updated at once")
	assert.Zero(t, g.erasedPixels(fieldRows-1))
	assert.Zero(t, g.erasedPixels(fieldRows-2))
	assert.False(t, g.animate(clearStepTime/2), "the step is not finished")

	// each step erases one more pixel row of the line, but the line is still solid
	for erased := 1; erased <= squareH; erased++ {
		assert.True(t, g.animate(clearStepTime))
		assert.Equal(t, erased, g.erasedPixels(fieldRows-1))
		assert.Zero(t, g.erasedPixels(fieldRows-2))
		assert.NotZero(t, g.field[fieldRows-1][0])
	}
	assert.Equal(t, squareH, g.clearStep)

	// the last step: the squares above fall down
	g.field[fieldRows-2][0] = 5
	assert.True(t, g.animate(clearStepTime))
	assert.Empty(t, g.clearing)
	assert.Zero(t, g.clearStep)
	assert.Equal(t, uint8(5), g.field[fieldRows-1][0])
	assert.Zero(t, g.field[fieldRows-2][0])
	assert.False(t, g.animate(clearStepTime))
}

func Test_play_during_animation(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	fillRowExceptSlot(g, fieldRows-1)
	g.hardDrop()
	require.NotEmpty(t, g.clearing)

	// the next block is O, it's controlled and falls while the line is being removed
	require.Equal(t, 1, g.cur.kind)
	y := g.cur.y
	g.fall = 0.9
	assert.True(t, g.update(clearStepTime))
	assert.Equal(t, y+1, g.cur.y)
	assert.Equal(t, 1, g.clearStep, "the animation goes on")
	assert.True(t, g.move(-1))
	assert.True(t, g.rotate())

	// the block lands on the line that is still solid, and falls with everything above it
	g.hardDrop()
	assert.Empty(t, g.clearing)
	assert.Equal(t, lineScore, g.score)
	assert.Equal(t, uint8(2), g.field[fieldRows-1][4])
	assert.Equal(t, uint8(2), g.field[fieldRows-2][5])
	assert.Zero(t, g.field[fieldRows-3][4])
	assert.False(t, g.over)
}

func Test_animation_by_time(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	fillRowExceptSlot(g, fieldRows-1)
	g.hardDrop()

	for range squareH {
		g.update(clearStepTime)
		assert.NotEmpty(t, g.clearing)
	}
	g.update(clearStepTime)
	assert.Empty(t, g.clearing)
}

func Test_new_lines_during_animation(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 0
	fillRowExceptSlot(g, fieldRows-1)
	fillRowExceptSlot(g, fieldRows-2)
	g.hardDrop() // completes only the bottom line, the line above it still misses the slot
	require.Equal(t, []int{fieldRows - 1}, g.clearing)

	// the next I block completes the second line, the first one is removed at once
	g.hardDrop()
	assert.Equal(t, []int{fieldRows - 1}, g.clearing, "the second line is on the bottom now")
	assert.Equal(t, 2*lineScore, g.score)
	assert.Zero(t, g.erasedPixels(fieldRows-1))
	assert.Equal(t, 2, g.lines)
}

func Test_spawn_blocked_by_removing_line(t *testing.T) {
	g := newTestGame(1) // O, rows 0-1
	for col := range fieldCols {
		g.field[1][col] = 3
	}
	g.clearing = []int{1}

	g.next = 1
	g.spawn()

	assert.False(t, g.over, "the block fits after the line is removed")
	assert.Empty(t, g.clearing)
}

func Test_settle_floating_square(t *testing.T) {
	g := newTestGame(1)
	g.field[10][9] = 5 // nothing under it
	g.field[fieldRows-1][0] = 3
	g.field[fieldRows-2][0] = 3

	g.settle()

	assert.Zero(t, g.field[10][9])
	assert.Equal(t, uint8(5), g.field[fieldRows-1][9], "the square fell to the floor")
	assert.Equal(t, uint8(3), g.field[fieldRows-2][0], "the stack didn't move")
}

func Test_settle_keeps_connected(t *testing.T) {
	g := newTestGame(1)
	g.field[fieldRows-1][0] = 3
	g.field[fieldRows-2][0] = 3
	g.field[fieldRows-2][1] = 3 // overhang, connected by the side
	g.field[fieldRows-3][1] = 4 // on the overhang
	want := g.field

	g.settle()

	assert.Equal(t, want, g.field)
}

func Test_settle_group_falls_as_whole(t *testing.T) {
	g := newTestGame(1)
	g.field[fieldRows-1][3] = 3
	g.field[fieldRows-2][3] = 3 // the stack is under the left square of the group only
	g.field[10][3], g.field[10][4] = 5, 6

	g.settle()

	assert.Equal(t, uint8(5), g.field[fieldRows-3][3], "the group stopped on the stack")
	assert.Equal(t, uint8(6), g.field[fieldRows-3][4], "the right square is in the air, but it's connected")
	assert.Zero(t, g.field[fieldRows-1][4])
	assert.Zero(t, g.field[10][3])
}

func Test_settle_several_groups(t *testing.T) {
	g := newTestGame(1)
	g.field[5][0] = 4  // the upper group
	g.field[10][0] = 5 // the lower group, the upper one lands on it

	g.settle()

	assert.Equal(t, uint8(5), g.field[fieldRows-1][0])
	assert.Equal(t, uint8(4), g.field[fieldRows-2][0])
	assert.Zero(t, g.field[5][0])
	assert.Zero(t, g.field[10][0])
}

// setupHangingSquare - the bottom row is full except the column 9, above it is the row that the I block completes,
// and there is a square above that row in the column 9, so after removing of the line the square hangs over the hole
func setupHangingSquare(g *game) {
	for col := range fieldCols {
		if col != 9 {
			g.field[fieldRows-1][col] = 3
		}
		if col < 4 || col >= 8 {
			g.field[fieldRows-2][col] = 3
		}
	}
	g.field[fieldRows-3][9] = 5
}

func Test_clear_line_drops_hanging_square(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	setupHangingSquare(g)
	g.hardDrop()
	require.Equal(t, []int{fieldRows - 2}, g.clearing)

	g.finishClearing()

	assert.Equal(t, uint8(5), g.field[fieldRows-1][9], "the square fell into the hole")
	assert.Zero(t, g.field[fieldRows-2][9])
	assert.Zero(t, g.field[fieldRows-3][9])
}

func Test_fallen_square_completes_line(t *testing.T) {
	g := newTestGame(0) // I
	g.next = 1
	setupHangingSquare(g)
	g.hardDrop()
	require.Equal(t, lineScore, g.score)

	// the animation of the first line is over, the square falls and completes the bottom line
	for range squareH {
		g.animate(clearStepTime)
		require.Equal(t, []int{fieldRows - 2}, g.clearing)
	}
	assert.True(t, g.animate(clearStepTime))

	assert.Equal(t, []int{fieldRows - 1}, g.clearing, "the new line is being removed")
	assert.Equal(t, 2*lineScore, g.score)
	assert.Equal(t, 2, g.lines)
}

func Test_modeFlag(t *testing.T) {
	mode := modeFlag{tcg.Mode2x3}
	assert.Equal(t, "2x3", mode.String())

	require.NoError(t, mode.Set("2x4Braille"))
	assert.Equal(t, 4, mode.Height())
	require.NoError(t, mode.Set("2x3"))
	assert.Equal(t, 3, mode.Height())

	for _, in := range []string{"1x1", "1x2", "2x2", "1x1Simple", "3x2", ""} {
		assert.Error(t, mode.Set(in), "mode %q", in)
	}
	assert.Equal(t, 3, mode.Height(), "the mode is not changed after errors")
}

func Test_modes(t *testing.T) {
	t.Cleanup(func() { setMode(tcg.Mode2x3) })

	for _, mode := range []tcg.PixelMode{tcg.Mode2x3, tcg.Mode2x4Braille} {
		t.Run(mode.String(), func(t *testing.T) {
			setMode(mode)

			// the square is two characters wide and one character high
			assert.Equal(t, 2*mode.Width(), squareW)
			assert.Equal(t, mode.Height(), squareH)

			// the full line is erased by pixel rows of the squares
			g := newTestGame(0) // I
			g.next = 1
			fillRowExceptSlot(g, fieldRows-1)
			g.hardDrop()
			require.Equal(t, []int{fieldRows - 1}, g.clearing)
			for erased := 1; erased <= squareH; erased++ {
				require.True(t, g.animate(clearStepTime))
				assert.Equal(t, erased, g.erasedPixels(fieldRows-1))
			}
			assert.True(t, g.animate(clearStepTime))
			assert.Empty(t, g.clearing)

			// a block moved by one pixel takes a part of the next row of characters
			g = newTestGame(1) // O
			cx := g.cur.col * squareCharsW
			g.cur.y++
			assert.Equal(t, uint8(2), g.cellShape(cx, 2*squareCharsH))
			assert.Zero(t, g.cellShape(cx, 2*squareCharsH+1))
			assert.Zero(t, g.cellShape(cx+2*squareCharsW, 0))
		})
	}
}
