package main

import (
	"testing"
)

/**
 * @author  CoreyChen Zhang
 * @version  2021/4/13 14:51
 * @modified
 */

func TestAdd(t *testing.T) {
	tests := []struct{a, b, c int}{
		{0, 0, 0},
		{1, 2, 3},
		{5, 770, 775},
		//{math.MaxInt32, 1, math.MinInt32},
	}

	for _ , tt := range tests {
		if actual := add(tt.a, tt.b); actual != tt.c {
			t.Errorf("calculate add error(%d  + %d), get %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {

	answer := 2

	for i := 0; i < b.N; i++ {
		actual := add(1, 1)
		if  actual != answer {
			b.Errorf("calculate add error(%d  + %d), get %d; expected %d", 1, 1, actual, answer)
		}
	}

}