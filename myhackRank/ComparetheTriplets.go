package myhackerRank

func ComparetheTriplets(x, y []int32) []int32 {
	r := []int32{0, 0}
	for i := 0; i < len(x); i++ {
		if x[i] > y[i] {
			r[0]++
		} else {
			r[1]++
		}
	}
	return r
}
