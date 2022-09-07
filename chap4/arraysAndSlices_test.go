package arraysAndSlices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "숫자 1,2,3,4,5 를 입력하면 총 합 15를 반환",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "숫자 1,2,3,4 를 입력하면 총 합 10를 반환",
			args: args{
				numbers: []int{1, 2, 3, 4},
			},
			want: 10,
		},
		{
			name: "숫자 1,2,3 를 입력하면 총 합 6를 반환",
			args: args{
				numbers: []int{1, 2, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	type args struct {
		numbersToSum [][]int
	}
	tests := []struct {
		name     string
		args     args
		wantSums []int
	}{
		{
			name: "입력이 {1,2,3}, {1,2} 라면 {6, 3}을 반환",
			args: args{
				numbersToSum: [][]int{{1, 2, 3}, {1, 2}},
			},
			wantSums: []int{6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSums := SumAll(tt.args.numbersToSum...); !reflect.DeepEqual(gotSums, tt.wantSums) {
				t.Errorf("SumAll() = %v, want %v", gotSums, tt.wantSums)
			}
		})
	}
}
