package turtle

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var reComment = regexp.MustCompile(`#.*$`)
var reSpaces = regexp.MustCompile(`\s+`)
var reCommands = regexp.MustCompile(`([a-zA-Z]+)(?:(-?\d+)(?:,(-?\d+))?)?`)

// DrawScript - draw by script
// U12 G 1,-1 # up 12 times, and goto to (1,-1)
func (t *Turtle) DrawScript(script string) *Turtle {
	scanner := bufio.NewScanner(strings.NewReader(script))
	for scanner.Scan() {
		line := scanner.Text()
		line = reSpaces.ReplaceAllString(reComment.ReplaceAllString(line, ""), "")
		for _, cmd := range reCommands.FindAllStringSubmatch(line, -1) {
			if len(cmd) != 4 {
				continue
			}

			name := cmd[1]
			p1, _ := strconv.Atoi(cmd[2])
			p2, _ := strconv.Atoi(cmd[3])

			switch strings.ToUpper(name) {
			case "S":
				t.Set()
			case "C":
				t.SetColor(p1)
			case "N":
				t.Raise()
			case "Y":
				t.Put()
			case "U":
				t.Up(p1)
			case "D":
				t.Down(p1)
			case "R":
				t.Right(p1)
			case "L":
				t.Left(p1)
			case "UR":
				t.UpRight(p1)
			case "UL":
				t.UpLeft(p1)
			case "DR":
				t.DownRight(p1)
			case "DL":
				t.DownLeft(p1)
			case "G":
				t.GoTo(p1, p2)
			case "GA":
				t.GoToAbs(p1, p2)
			case "LT":
				t.LineTo(p1, p2)
			}
		}
	}

	return t
}
