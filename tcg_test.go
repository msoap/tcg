//go:generate minimock -i github.com/gdamore/tcell/v2.Screen -o .

package tcg

import (
	"reflect"
	"strings"
	"testing"
)

func TestRenderAsStrings(t *testing.T) {
	tests := []struct {
		name string
		img  []string
		mode PixelsInChar
		want []string
	}{
		{
			name: "1",
			img: []string{
				"..**..",
				".*..*.",
				".****.",
			},
			mode: Mode2x3,
			want: []string{"🬦🬰🬓"},
		},
		{
			name: "2",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RenderAsStrings(MustNewBufferFromStrings(tt.img), tt.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenderAsStrings() = \n%s\n, want \n%s\n", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}
