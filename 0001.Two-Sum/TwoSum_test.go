package gotest

import (
	"reflect"
	"testing"
)

// func TestTwoSum(t *testing.T) {
// 	type args struct {
// 		nums   []int
// 		target int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		// TODO: Add test cases.
// 		{"1", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
// 		{"2", args{[]int{3, 2, 4}, 5}, []int{0, 1}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestTwoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		// {"2", args{[]int{3, 2, 4}, 5}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
