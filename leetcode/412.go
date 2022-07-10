package leetcode

import "strconv"

// func FizzBuzz(n int) []string {
// 	ans := []string{}
// 	for i := 1; i <= n; i++ {
// 		switch {
// 		case i%3 == 0 && i%5 == 0:
// 			ans = append(ans, "FizzBuzz")
// 		case i%3 == 0:
// 			ans = append(ans, "Fizz")
// 		case i%5 == 0:
// 			ans = append(ans, "Buzz")
// 		default:
// 			ans = append(ans, strconv.Itoa(i))
// 		}
// 	}
// 	return ans
// }

func FizzBuzz(n int) []string {
	ans := []string{}
	s := ""
	for i := 1; i <= n; i++ {
		s = ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if len(s) == 0 {
			s = strconv.Itoa(i)
		}
		ans = append(ans, s)
	}
	return ans
}
