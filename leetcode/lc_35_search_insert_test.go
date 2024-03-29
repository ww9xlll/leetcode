package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"4", args{[]int{1}, 2}, 1},
		{"4", args{[]int{1}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
