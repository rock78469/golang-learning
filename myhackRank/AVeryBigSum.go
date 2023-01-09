package myhackerRank

func AVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, i := range ar {
		sum += i
	}
	return sum
}
