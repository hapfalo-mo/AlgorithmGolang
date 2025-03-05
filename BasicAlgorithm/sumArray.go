package main

import "fmt"

func SumArray(arr []int64) int64 {
	var sum int64 = 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func SumArraySolution() {
	arr := []int64{2, 5, 6, 4, 8, 9}
	fmt.Println(SumArray(arr))
}
