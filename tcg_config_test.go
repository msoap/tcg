package tcg

import "testing"

func TestParseSizeString(t *testing.T) {
	tests := []struct {
		in      string
		w       int
		h       int
		wantErr bool
	}{
		{
			in:      "",
			w:       0,
			h:       0,
			wantErr: true,
		},
		{
			in:      "x",
			w:       0,
			h:       0,
			wantErr: true,
		},
		{
			in:      "80 x 0",
			w:       0,
			h:       0,
			wantErr: true,
		},
		{
			in:      "-80 x -20",
			w:       0,
			h:       0,
			wantErr: true,
		},
		{
			in:      "80x25",
			w:       80,
			h:       25,
			wantErr: false,
		},
		{
			in:      "80 x 25",
			w:       80,
			h:       25,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, got1, err := ParseSizeString(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSizeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.w {
				t.Errorf("ParseSizeString() got = %v, want %v", got, tt.w)
			}
			if got1 != tt.h {
				t.Errorf("ParseSizeString() got1 = %v, want %v", got1, tt.h)
			}
		})
	}
}
