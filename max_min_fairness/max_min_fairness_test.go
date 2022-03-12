package main

import (
	"testing"
)

func Test_weightMaxMinFairness(t *testing.T) {
	t.Log(weightMaxMinFairness(16, map[string][]float64{"1": {4, 2.5}, "2": {2, 4},
		"3": {10, 0.5}, "4": {7, 1}}))
}

func Test_normalized(t *testing.T) {
	t.Log(normalized([]float64{2.5, 4, 0.5, 1}))       // [5, 8, 1, 2]
	t.Log(normalized([]float64{2.5, 4, 0.5, 1, 1.25})) // [10, 16, 2, 4, 5]
}

func Test_simplify(t *testing.T) {
	arr := []int64{6, 12, 18, 24, 36}
	simplify(arr)
	t.Log(arr)
	arr = []int64{1, 12, 24, 36}
	simplify(arr)
	t.Log(arr)
	arr = []int64{45, 9, 63, 27}
	simplify(arr)
	t.Log(arr)
	arr = []int64{43, 12, 3, 4}
	simplify(arr)
	t.Log(arr)
	arr = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	simplify(arr)
	t.Log(arr)
}

func Test_gcd(t *testing.T) {
	t.Log(gcd(60, 30))
}
