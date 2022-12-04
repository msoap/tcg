package tcg

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
			if got := RenderAsStrings(MustNewBufferFromStrings(tt.img), tt.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenderAsStrings() = \n%s\n, want \n%s\n", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}

func TestRenderAsStringsWithChatMap(t *testing.T) {
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
			got := RenderAsStrings(MustNewBufferFromStrings(tt.img), *mode)
			assert.Equal(t, tt.want, got)
		})
	}
}
