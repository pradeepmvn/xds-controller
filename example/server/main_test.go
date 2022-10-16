package main

import (
	"fmt"
	"testing"
)

func Test_randNum(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"Length test", args{length: 5}, 12345},
		{"Length test", args{length: 6}, 132345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randNum(tt.args.length)
			t.Log("Matc: randNum() = " + fmt.Sprint(got) + ", want " + fmt.Sprint(tt.want))
			if getLength(got) != getLength(tt.want) {
				t.Errorf("randNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getLength(num int32) int32 {
	var c int32 = 0
	for num != 0 {
		num /= 10
		c += 1
	}
	return c
}
