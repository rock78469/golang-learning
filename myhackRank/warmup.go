package myhackerRank

import (
	"fmt"
	"strconv"
	"strings"
)

// A Very Big Sum
func AVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, i := range ar {
		sum += i
	}
	return sum
}

// ComparetheTriplets
func ComparetheTriplets(a, b []int32) []int32 {
	r := []int32{0, 0}
	for i := 0; i < len(a); i++ {
		switch {
		case a[i] > b[i]:
			r[0]++
		case a[i] < b[i]:
			r[1]++
		case a[i] == b[i]:
		}
	}
	return r
}

// Simple Array Sum
func SimpleArraySum(n []int32) int32 {
	var sum int32
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

// Diagonal Difference
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

// Birthday Cake Candles
func BirthdayCakeCandles(candles []int32) int32 {
	var x int32
	var y int32

	for i := 0; i < len(candles); i++ {
		for j := i; j < len(candles); j++ {
			if candles[i] == candles[j] {
				fmt.Printf("A= %v ,B= %v \n", candles[i], candles[j])
				x++
			}
		}

		if y < x {
			y = x
		}
		x = 0

	}

	return y
}

func TimeConversion(s string) string {
	// var r string
	var str1 string
	var str2 []string
	if strings.Contains(s, "PM") {
		str1 = strings.Replace(s, "PM", "", -1)
		str2 = strings.Split(str1, ":")
		h, m, s := str2[0], str2[1], str2[2]
		newH, _ := strconv.Atoi(h)
		if newH < 12 {
			newH += 12
		}
		h = strconv.Itoa(newH)
		if len(h) < 2 {
			h = "0" + h
		}
		return fmt.Sprintf("%v:%v:%v", h, m, s)
	}
	str1 = strings.Replace(s, "AM", "", -1)
	str2 = strings.Split(str1, ":")
	h, m, s := str2[0], str2[1], str2[2]
	newH, _ := strconv.Atoi(h)
	if newH >= 12 {
		newH -= 12
	}
	h = strconv.Itoa(newH)
	if len(h) < 2 {
		h = "0" + h
	}
	return fmt.Sprintf("%v:%v:%v", h, m, s)

}
