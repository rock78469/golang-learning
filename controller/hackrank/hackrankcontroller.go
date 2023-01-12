package hackrank

import (
	"fmt"
	m "go-learning/myhackRank"
)

// Simple Array Sum
func SimpleArraySum() {
	n := []int32{1, 2, 3, 4, 5, 6, 7, 8, 8}
	r := m.SimpleArraySum(n)
	fmt.Println(r)

}

// AVery Big Sum
func AVeryBigSum() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	r := m.AVeryBigSum(ar)
	fmt.Println(r)
}

// Comparethe Triplets
func ComparetheTriplets() {
	userA := []int32{2, 5, 7}
	userB := []int32{3, 1, 8}
	r := m.ComparetheTriplets(userA, userB)
	fmt.Println(r)
}

// Diagonal Difference
func DiagonalDifference() {
	// cc := [][]int32{{1, 4, 8, 8}, {6, 2, 9, 2}, {7, 4, 2, 6}, {6, 7, 4, 3}}
	cc := [][]int32{
		{6, 6, 7, -10, 9, -3, 8, 9, -1},
		{9, 7, -10, 6, 4, 1, 6, 1, 1},
		{-1, -2, 4, -6, 1, -4, -6, 3, 9},
		{-8, 7, 6, -1, -6, -6, 6, -7},
		{-10, -4, 9, 1, -7, 8, -5, 3, -5},
		{-8, -3, -4, 2, -3, 7, -5, 1, -5},
		{-2, -7, -4, 8, 3, -1, 8, 2, 3},
		{-3, 4, 6, -7, -7, -8, -3, 9, -6},
		{-2, 0, 5, 4, 4, 4, -3, 3, 0},
	}
	r := m.DiagonalDifference(cc)
	fmt.Println(r)
}

func BirthdayCakeCandles() {
	candles := []int32{2, 1, 3, 45, 6, 1, 23, 2, 6, 23, 2, 1, 4, 5, 2, 2, 352, 3, 5, 2, 3, 3, 4, 2, 3, 4, 2, 3}
	r := m.BirthdayCakeCandles(candles)
	fmt.Println(r)
}

func TimeConversion() {
	time := "04:59:59AM"
	r := m.TimeConversion(time)
	fmt.Println(r)
}

func GradingStudents() {
	grades := []int32{23, 63, 78, 56, 89}
	r := m.GradingStudents(grades)
	fmt.Println(r)
}
