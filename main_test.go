package firstHomework

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		slice []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "usual slice, rotate 0",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
				n:     0,
			},
			want: []int{2, 3, 5, 7, 11, 13},
		},
		{
			name: "zero slice, rotate 100",
			args: args{
				slice: []int{},
				n:     100,
			},
			want: []int{},
		},
		{
			name: "usual slice, rotate -100",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
				n:     -100,
			},
			want: []int{2, 3, 5, 7, 11, 13},
		},
		{
			name: "usual slice, rotate 1",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
				n:     1,
			},
			want: []int{13, 2, 3, 5, 7, 11},
		},
		{
			name: "usual slice, rotate 3",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
				n:     3,
			},
			want: []int{7, 11, 13, 2, 3, 5},
		},
		{
			name: "usual slice, rotate 1 * 3 len(slice)",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
				n:     1 + 3*6,
			},
			want: []int{13, 2, 3, 5, 7, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_revert(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "zero slice",
			args: args{
				slice: []int{},
			},
			want: []int{},
		},
		{
			name: "mirror slice",
			args: args{
				slice: []int{1, 2, 3, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 3, 2, 1},
		},
		{
			name: "usual slice with even n",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13},
			},
			want: []int{13, 11, 7, 5, 3, 2},
		},
		{
			name: "usual slice with odd n",
			args: args{
				slice: []int{2, 3, 5, 7, 11, 13, 95},
			},
			want: []int{95, 13, 11, 7, 5, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revert(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revert() = %v, want %v", got, tt.want)
			}
		})
	}
}
