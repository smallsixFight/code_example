package main

import (
	"log"
	"math/rand"
	"time"
)

func durstenfeld(src []int) []int {
	now := time.Now().UnixNano()
	log.Println("start", now)
	defer func() {
		log.Println("duration", time.Now().UnixNano()-now)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(src); i > 0; i-- {
		idx := r.Intn(i)
		src[idx], src[i-1] = src[i-1], src[idx]
	}
	return src
}

func fisherYatesWithSlice(src []int) []int {
	now := time.Now().UnixNano()
	log.Println("start", now)
	defer func() {
		log.Println("duration", time.Now().UnixNano()-now)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, 0, len(src))
	for len(src) > 0 {
		idx := r.Intn(len(src))
		res = append(res, src[idx])
		src = append(src[:idx], src[idx+1:]...)
		//fmt.Printf("%p\n", src)
	}
	return res
}

func fisherYates(src []int) []int {
	now := time.Now().UnixNano()
	log.Println("start", now)
	defer func() {
		log.Println("duration", time.Now().UnixNano()-now)
	}()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, 0, len(src))
	for len(src) > 0 {
		idx := r.Intn(len(src))
		res = append(res, src[idx])
		temp := make([]int, 0, len(src)-1)
		for i := 0; i < len(src); i++ {
			if i != idx {
				temp = append(temp, src[i])
			}
		}
		//fmt.Printf("%p ", src)
		src = temp
		//fmt.Printf("%p\n", src)
	}
	return res
}

func main() {
	log.Println("Fisher-Yates shuffle:")
	fisherYates(getArr())
	log.Println("==================================================")
	log.Println("Fisher-Yates shuffle with slice:")
	fisherYatesWithSlice(getArr())
	log.Println("==================================================")
	log.Println("Durstenfeld shuffle:")
	durstenfeld(getArr())
}

func getArr() []int {
	arr := make([]int, 100000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}
