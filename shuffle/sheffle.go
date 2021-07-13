package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fisherYates(src []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		res = append(res, r.Intn(10))
	}
	return res
}

func main() {
	fmt.Println(fisherYates())
}
