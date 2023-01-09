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

func ComparetheTriplets() {
	userA := []int32{2, 5, 7}
	userB := []int32{3, 1, 8}
	r := m.ComparetheTriplets(userA, userB)
	fmt.Println(r)
}

// Diagonal Difference
func DiagonalDifference() {
	cc := [][]int32{{1, 4, 8, 8}, {6, 2, 9, 2}, {7, 4, 2, 6}, {6, 7, 4, 3}}
	r := m.DiagonalDifference(cc)
	fmt.Println(r)
}
