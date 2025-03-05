package main

import "fmt"

// Excersise 1: Solve Me First (HackerRank)
func Solution(a int64, b int64) int64 {
	return a + b
}

//
func SolveMeFirst() {
	var a, b int64
	fmt.Println("Enter the a number:")
	fmt.Scanf("%d\n", &a)
	fmt.Println("Enter the b number:")
	fmt.Scanf("%d", &b)
	result := Solution(a, b)
	fmt.Println("The value is :", result)
}
