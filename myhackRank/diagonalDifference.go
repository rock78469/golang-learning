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
		fmt.Printf("sumA: arr[%v][%v] = %v", i, x, arr[i][x])
		fmt.Printf("\n")
		fmt.Printf("sumB: arr[%v][%v] = %v", i, y, arr[i][y])
		// fmt.Printf("sumA: %v sumB: %v", arr[i][x], arr[i][y])
		fmt.Printf("\n")
		fmt.Printf("\n")
		x++
		y--
	}
	if sumA-sumB < 0 {
		return -(sumA - sumB)
	}
	return (sumA - sumB)
}
