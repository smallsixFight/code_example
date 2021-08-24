package main

import (
	"fmt"
	"testing"
)

func Test_gcd(t *testing.T) {
	t.Log(gcd(60, 30))
}

func Test_simplify(t *testing.T) {
	arr := []int64{6, 12, 18, 24, 36}
	simplify(arr)
	fmt.Println(arr)
}
