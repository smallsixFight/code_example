package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 最大最小公平分配算法
func main() {
	demand := map[string]int64{"1": 10, "2": 2, "3": 5, "4": 3}
	fmt.Println(maxMinFairness(8, demand))
}

/*
1. 对权重归一化处理；
2. 对需要进行分配的需求，以资源需求作为条件进行升序排序；
3. 根据权重进行资源分配，分配资源满足需求的放入结果集；
4. 对剩余的资源进行汇总，整理剩余需求的权重数据，并进行简化；
5. 重复步骤 3，直到剩余资源不能满足剩下任意一个需求时，直接按权重分配资源，并保存到结果集中。
*/
func weightMaxMinFairness(resource float64, demands map[string][]float64) map[string]float64 {
	idxMap, ws := make(map[string]int), make([]float64, 0, len(demands))
	arr := make([]string, 0, len(demands))
	for k := range demands {
		ws = append(ws, demands[k][1])
		arr = append(arr, k)
		idxMap[k] = len(ws) - 1
	}
	normalize := normalized(ws) // 归一化
	res := make(map[string]float64)
	canAllow := true
	for canAllow {
		canAllow = false
		sort.Slice(arr, func(i, j int) bool {
			return demands[arr[i]][0] > demands[arr[j]][0]
		})
		var totalWeight int64 = 0
		for i := range normalize {
			totalWeight += normalize[i]
		}
		singleResource := resource / float64(totalWeight)
		for i := len(arr) - 1; i >= 0; i-- {
			idx := idxMap[arr[i]]
			if v := singleResource * float64(normalize[idx]); v >= demands[arr[i]][0] {
				res[arr[i]] = demands[arr[i]][0]
				resource -= demands[arr[i]][0]
				arr = append(arr[:i], arr[i+1:]...)
				normalize = append(normalize[:idx], normalize[idx+1:]...)
				for k := range idxMap {
					if idxMap[k] > idx {
						idxMap[k]--
					}
				}
				canAllow = true
			}
		}
		if canAllow {
			simplify(normalize)
		}
	}
	// 若存在积压的需求，则根据权重的比例进行分配
	var totalWeight int64 = 0
	for i := range normalize {
		totalWeight += normalize[i]
	}
	singleResource := resource / float64(totalWeight)
	for i := len(arr) - 1; i >= 0; i-- {
		res[arr[i]] = math.Floor(singleResource*float64(normalize[idxMap[arr[i]]]) + 0.5)
	}
	return res
}

// 归一化
func normalized(weights []float64) []int64 {
	maxLen := 0
	for i := range weights {
		if arr := strings.Split(fmt.Sprint(weights[i]), "."); len(arr) > 1 {
			if len(arr[1]) > maxLen {
				maxLen = len(arr[1])
			}
		}
	}
	fa := 1
	for i := 0; i < maxLen; i++ {
		fa *= 10
	}
	res := make([]int64, len(weights))
	for i := range weights {
		res[i] = int64(weights[i] * float64(fa))
	}
	simplify(res)
	return res
}

// 求一组数字的最大公因数，然后进行简化
func simplify(arr []int64) {
	if len(arr) < 2 {
		return
	}
	temp := make([]int64, len(arr))
	copy(temp, arr)
	// 先从小到大排序，从头遍历每两位取最大公因数，然后与之前的得到公因数做比较，取最大公因数。
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	val := gcd(temp[0], temp[1])
	for i := 2; i < len(temp); i++ {
		val = gcd(val, gcd(temp[i-1], temp[i]))
	}
	for i := 0; i < len(temp); i++ {
		arr[i] /= val
	}
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
