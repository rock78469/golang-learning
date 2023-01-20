package myhackerRank

// CamelCase 如果ascii code 的位置小於 a的ascii code ，則判斷為大寫
func CamelCase(s string) int32 {
	r := int32(1)
	for _, x := range s {
		if x < 97 {
			r++
		}

	}
	return r
}

func StrongPassword(n int32, password string) int32 {
	r := int32(0)
	l := len(password)
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-+"
	// lower_case := "abcdefghijklmnopqrstuvwxyz"
	// upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// special_characters := "!@#$%^&*()-+"

	for _, r1 := range str {
		for _, r2 := range password {
			if r1 == r2 {
				break
			}
		}
	}
	for l < 6 {
		r += 1
	}
	return r
}
