package tcg

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_allocateBuffer(t *testing.T) {
	type args struct {
		w int
		h int
	}

	tests := []struct {
		name       string
		args       args
		wantLen    int
		wantSubLen int
	}{
		{
			name:       "1x1 pixels",
			args:       args{w: 1, h: 1},
			wantLen:    1,
			wantSubLen: 1,
		},
		{
			name:       "2x1 pixels",
			args:       args{w: 2, h: 1},
			wantLen:    1,
			wantSubLen: 1,
		},
		{
			name:       "8x1 pixels",
			args:       args{w: 8, h: 1},
			wantLen:    1,
			wantSubLen: 1,
		},
		{
			name:       "9x1 pixels",
			args:       args{w: 9, h: 1},
			wantLen:    1,
			wantSubLen: 2,
		},
		{
			name:       "1x2 pixels",
			args:       args{w: 1, h: 2},
			wantLen:    2,
			wantSubLen: 1,
		},
		{
			name:       "8x8 pixels",
			args:       args{w: 8, h: 8},
			wantLen:    8,
			wantSubLen: 1,
		},
		{
			name:       "16x16 pixels",
			args:       args{w: 16, h: 16},
			wantLen:    16,
			wantSubLen: 2,
		},
		{
			name:       "17x16 pixels",
			args:       args{w: 17, h: 16},
			wantLen:    16,
			wantSubLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allocateBuffer(tt.args.w, tt.args.h)
			require.Len(t, got, tt.wantLen, "length of buffer broken")
			for i, row := range got {
				require.Len(t, row, tt.wantSubLen, "length of buffer row (%d) is broken", i)
			}
		})
	}
}

// parse binary string to byte
func mustBin(str string) []byte {
	str = regexp.MustCompile(`_|\s+`).ReplaceAllString(str, "")
	str = regexp.MustCompile(`\.`).ReplaceAllString(str, "0")
	str = regexp.MustCompile(`\*`).ReplaceAllString(str, "1")
	rem := len(str) % 8
	if rem > 0 {
		str += strings.Repeat("0", 8-rem)
	}

	result := []byte{}
	for i := 0; i < len(str); i += 8 {
		byteStr := str[i : i+8]
		b, err := strconv.ParseUint(byteStr, 2, 8)
		if err != nil {
			panic(err)
		}
		result = append(result, byte(b))
	}

	return result
}

func TestBuffer_Set(t *testing.T) {
	type fields struct {
		Width  int
		Height int
		buffer [][]byte
	}
	type args struct {
		x     int
		y     int
		color int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]byte
	}{
		{
			name:   "0x0 black pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("0000_0000")}},
			args:   args{x: 0, y: 0, color: Black},
			want:   [][]byte{mustBin("1000_0000")},
		},
		{
			name:   "0x0 white pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("1111_1111")}},
			args:   args{x: 0, y: 0, color: White},
			want:   [][]byte{mustBin("0111_1111")},
		},
		{
			name:   "1x0 black pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("0000_0000")}},
			args:   args{x: 1, y: 0, color: Black},
			want:   [][]byte{mustBin("0100_0000")},
		},
		{
			name:   "1x0 white pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("1111_1111")}},
			args:   args{x: 1, y: 0, color: White},
			want:   [][]byte{mustBin("1011_1111")},
		},
		{
			name:   "0x1 black pixel",
			fields: fields{Width: 8, Height: 2, buffer: [][]byte{mustBin("0000_0000"), mustBin("0000_0000")}},
			args:   args{x: 0, y: 1, color: Black},
			want:   [][]byte{mustBin("0000_0000"), mustBin("1000_0000")},
		},
		{
			name:   "0x1 white pixel",
			fields: fields{Width: 8, Height: 2, buffer: [][]byte{mustBin("1111_1111"), mustBin("1111_1111")}},
			args:   args{x: 0, y: 1, color: White},
			want:   [][]byte{mustBin("1111_1111"), mustBin("0111_1111")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Buffer{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
				buffer: tt.fields.buffer,
			}
			b.Set(tt.args.x, tt.args.y, tt.args.color)
			require.Equal(t, tt.want, b.buffer)
		})
	}
}

func TestBuffer_At(t *testing.T) {
	type fields struct {
		Width  int
		Height int
		buffer [][]byte
	}
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "black pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("1000_0000")}},
			args:   args{x: 0, y: 0},
			want:   1,
		},
		{
			name:   "next pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{mustBin("1011_0000")}},
			args:   args{x: 1, y: 0},
			want:   0,
		},
		{
			name:   "white pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{{0x0F}}},
			args:   args{x: 0, y: 0},
			want:   0,
		},
		{
			name:   "black 4 pixel",
			fields: fields{Width: 8, Height: 1, buffer: [][]byte{{0x0F}}},
			args:   args{x: 4, y: 0},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buffer{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
				buffer: tt.fields.buffer,
			}
			if got := b.At(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Buffer.GetPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_widthInBytes(t *testing.T) {
	tests := []struct {
		name string
		w    int
		want int
	}{
		{w: 0, want: 0},
		{w: 1, want: 1},
		{w: 4, want: 1},
		{w: 7, want: 1},
		{w: 8, want: 1},
		{w: 9, want: 2},
		{w: 12, want: 2},
		{w: 15, want: 2},
		{w: 16, want: 2},
		{w: 17, want: 3},
		{w: 256, want: 32},
		{w: 257, want: 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthInBytes(tt.w); got != tt.want {
				t.Errorf("widthInBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewBufferFromImage(t *testing.T) {
	reader, err := os.Open("testdata/plus.png")
	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		t.Fatalf("failed to decode image: %s", err)
	}

	b := NewBufferFromImage(img)

	expected := MustNewBufferFromStrings([]string{
		"..........",
		"..........",
		"....**....",
		"....**....",
		"..******..",
		"..******..",
		"....**....",
		"....**....",
		"..........",
		"..........",
	})
	assertEqBuffers(t, b, expected)
}

func assertEqBuffers(t *testing.T, got, expected Buffer) {
	if got.Width != expected.Width {
		t.Errorf("buffer width of got (%d) != expected (%d)", got.Width, expected.Width)
		return
	}
	if got.Height != expected.Height {
		t.Errorf("buffer height of got (%d) != expected (%d)", got.Height, expected.Height)
		return
	}

	if !expected.IsEqual(got) {
		gotStrings, expectedStrings := got.Strings(), expected.Strings()
		msg := fmt.Sprintf("buffers isn't equal:\n%-*s | %-*s\n", got.Width, "got", expected.Width, "expected")
		for y := 0; y < got.Height; y++ {
			msg += fmt.Sprintf("%-*s | %-*s\n", got.Width, gotStrings[y], expected.Width, expectedStrings[y])
		}
		assert.True(t, false, msg)
	}
}

func TestBuffer_Cut(t *testing.T) {
	src := MustNewBufferFromStrings([]string{
		"..........",
		"...*****..",
		"..**...**.",
		".**.....**",
		".*.......*",
		".*.......*",
		".*.......*",
		".**.....**",
		"..**...**.",
		"...*****..",
	})
	new := src.Cut(2, 0, 6, 5)
	expected := MustNewBufferFromStrings([]string{
		"......",
		".*****",
		"**...*",
		"*.....",
		"......",
	})

	assertEqBuffers(t, new, expected)
}

func TestRenderAsStrings(t *testing.T) {
	tests := []struct {
		name string
		img  []string
		mode PixelMode
		want []string
	}{
		{
			name: "1x1",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			mode: Mode1x1,
			want: []string{
				"  ██  ",
				" █  █ ",
				" ████ ",
			},
		},
		{
			name: "1x2",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			mode: Mode1x2,
			want: []string{
				" ▄▀▀▄ ",
				" ▀▀▀▀ ",
			},
		},
		{
			name: "2x2",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			mode: Mode2x2,
			want: []string{
				"▗▀▖",
				"▝▀▘",
			},
		},
		{
			name: "2x3",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			mode: Mode2x3,
			want: []string{"🬦🬰🬓"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustNewBufferFromStrings(tt.img).RenderAsStrings(tt.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenderAsStrings() = \n%s\n, want \n%s\n", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}

func TestRenderAsStringsWithNewPixelMode(t *testing.T) {
	tests := []struct {
		name          string
		img           []string
		width, height int
		cmap          []rune
		want          []string
	}{
		{
			name: "1x1",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			width:  1,
			height: 1,
			cmap:   []rune{'0', '1'},
			want: []string{
				"001100",
				"010010",
				"011110",
			},
		},
		{
			name: "1x2",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			width:  1,
			height: 2,
			cmap:   []rune{'a', 'b', 'c', 'd'},
			want: []string{
				"abccba",
				"acccca",
			},
		},
		{
			name: "2x2",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			width:  2,
			height: 2,
			cmap:   []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'},
			want: []string{
				"1c2",
				"4c8",
			},
		},
		{
			name: "2x3",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			width:  2,
			height: 3,
			cmap: []rune{
				'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
				'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
				'w', 'x', 'y', 'z', '.', ',', '<', '>', '?', '/', '`', '~', ':', ';', '{', '}',
				'[', ']', '!', '@', '£', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=',
			},
			want: []string{"5@a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mode, err := NewPixelMode(tt.width, tt.height, tt.cmap)
			require.NoError(t, err, "%s: unexpected error", tt.name)
			got := MustNewBufferFromStrings(tt.img).RenderAsStrings(*mode)
			assert.Equal(t, tt.want, got)
		})
	}
}
