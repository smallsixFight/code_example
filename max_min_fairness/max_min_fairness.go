package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 最大最小公平分配算法
func main() {
	demand := map[string]int64{"1": 10, "2": 2, "3": 5, "4": 3}
	fmt.Println(maxMinFairness(8, demand))
}

func weightMaxMinFairness(resource int64, demand map[string][]int64) map[string]int64 {
	arr := make([]string, 0, len(demand))
	for k := range demand {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return demand[arr[i]][0] < demand[arr[j]][0]
	})

	res := make(map[string]int64)

	return res
}

func normalized(weights []float64) []int64 {
	maxLen := 1
	for i := range weights {
		if arr := strings.Split(strconv.FormatFloat(weights[i], 'f', 10, 64), ","); len(arr) > 1 {
			if len(arr[1]) > maxLen {
				maxLen = len(arr[1])
			}
		}
	}
	for i := 1; i < maxLen; i++ {
		maxLen *= 10
	}
	res := make([]int64, len(weights))
	for i := range weights {
		res[i] = int64(weights[i] * float64(maxLen))
	}
	simplify(res)
	return res
}

// todo wait for implement
func simplify(arr []int64) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	val := gcd(arr[0], arr[1])
	for i := 1; i < len(arr); i++ {
		tmp := gcd(arr[i-1], arr[i])
		if tmp > val {
			val, tmp = tmp, val
		}
		if val%tmp != 0 {
			val *= tmp
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] /= val
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int64) int64 {
	var tmp int64
	for {
		tmp = a % b
		if tmp > 0 {
			a = b
			b = tmp
		} else {
			return b
		}
	}
}

func maxMinFairness(resource int64, demand map[string]int64) map[string]int64 {
	arr := make([]string, 0, len(demand))
	for k := range demand {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return demand[arr[i]] < demand[arr[j]]
	})
	singleResource, extra := resource/int64(len(demand)), int64(0)
	idx := 0
	for ; idx < len(arr); idx++ {
		if singleResource+extra < demand[arr[idx]] {
			break
		}
		extra = singleResource + extra - demand[arr[idx]]
	}
	for ; idx < len(arr); idx++ {
		demand[arr[idx]] = singleResource + extra
		extra = 0
	}
	return demand
}
