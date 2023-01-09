package myhackerRank

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
