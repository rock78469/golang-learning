package myhackerRank

import "fmt"

func DiagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sumA int32
	var sumB int32
	x := int32(0)
	y := len(arr) - 1
	fmt.Println(y)
	for i := 0; i < len(arr); i++ {
		sumA += arr[i][x]
		sumB += arr[i][y]
		fmt.Printf("sumA: %v sumB: %v", arr[i][x], arr[i][y])
		fmt.Printf("\n")
		x++
		y--
	}
	if sumA < 0 {
		sumA = -sumA
	}
	if sumB < 0 {
		sumB = -sumB
	}

	if sumA > sumB {
		fmt.Printf("sumA %v - sumB %v = %v", sumA, sumB, sumA-sumB)
		fmt.Printf("\n")
		return sumA - sumB
	} else {
		fmt.Printf("sumB %v - sumA %v = %v", sumB, sumA, sumB-sumA)
		fmt.Printf("\n")
		return sumB - sumA
	}
}
