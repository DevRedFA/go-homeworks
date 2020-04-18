package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty slice",
			args: args{slice: []int{}},
			want: []int{},
		},
		{
			name: "one element",
			args: args{slice: []int{1}},
			want: []int{1},
		},
		{
			name: "already sorted slice",
			args: args{slice: []int{1, 2, 3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "reverse sorted slice",
			args: args{slice: []int{4, 3, 2, 1}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "negative numbers slice",
			args: args{slice: []int{-1, -4, -2, -3}},
			want: []int{-4, -3, -2, -1},
		},
		{
			name: "negative and positive numbers slice",
			args: args{slice: []int{-1, 4, -2, 3, 0}},
			want: []int{-2, -1, 0, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
