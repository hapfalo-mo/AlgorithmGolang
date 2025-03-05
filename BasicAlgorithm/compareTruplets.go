package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var flagA int32 = 0
	var flagB int32 = 0
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			flagA += 1
		} else if b[i] > a[i] {
			flagB += 1
		} else {
			flagA += 0
			flagB += 0
		}
	}
	result := []int32{flagA, flagB}
	return result
}

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}
