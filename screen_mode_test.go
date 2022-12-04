package tcg

import "testing"

func Test_checkCharMapping(t *testing.T) {
	type args struct {
		width  int
		height int
		cm     []rune
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "all zero",
			args:    args{width: 0, height: 0, cm: nil},
			wantErr: true,
		},
		{
			name:    "empty list",
			args:    args{width: 1, height: 1, cm: nil},
			wantErr: true,
		},
		{
			name:    "empty list2",
			args:    args{width: 1, height: 1, cm: []rune{}},
			wantErr: true,
		},
		{
			name:    "1x1 not valid",
			args:    args{width: 1, height: 1, cm: []rune{32}},
			wantErr: true,
		},
		{
			name:    "1x1 valid",
			args:    args{width: 1, height: 1, cm: []rune{32, 33}},
			wantErr: false,
		},
		{
			name:    "1x2 not valid",
			args:    args{width: 1, height: 2, cm: []rune{32}},
			wantErr: true,
		},
		{
			name:    "1x2 valid",
			args:    args{width: 1, height: 2, cm: []rune{32, 33, 34, 35}},
			wantErr: false,
		},
		{
			name:    "2x2 not valid",
			args:    args{width: 2, height: 2, cm: []rune{32, 33, 34, 35, 32, 33, 34, 35, 32, 33, 34, 35, 32, 33, 34}},
			wantErr: true,
		},
		{
			name:    "2x2 valid",
			args:    args{width: 2, height: 2, cm: []rune{32, 33, 34, 35, 32, 33, 34, 35, 32, 33, 34, 35, 32, 33, 34, 35}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCharMapping(tt.args.width, tt.args.height, tt.args.cm); (err != nil) != tt.wantErr {
				t.Errorf("checkCharMapping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
