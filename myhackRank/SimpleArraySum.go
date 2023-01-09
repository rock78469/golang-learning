package myhackerRank

func SimpleArraySum(n []int32) int32 {
	var sum int32
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}
