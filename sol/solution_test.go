package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	digits := []int{4, 3, 2, 1}
	for idx := 0; idx < b.N; idx++ {
		plusOne(digits)
	}
}
func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "digits = [1,2,3]",
			args: args{digits: []int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
		{
			name: " digits = [4,3,2,1]",
			args: args{digits: []int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: " digits = [9]",
			args: args{digits: []int{9}},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
