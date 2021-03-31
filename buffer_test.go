package tcg

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

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

func TestBuffer_PutPixel(t *testing.T) {
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

func TestBuffer_GetPixel(t *testing.T) {
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
