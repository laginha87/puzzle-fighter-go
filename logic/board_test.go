package logic

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want Board
	}{
		{
			name: "Test",
			args: args{width: 2, height: 3},
			want: Board{
				grid: [][]*Block{{nil, nil}, {nil, nil}, {nil, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
