package myhackerRank

func CamelCase(s string) int32 {
	r := int32(1)
	for _, x := range s {
		if x < 97 {
			r++
		}

	}
	return r
}
