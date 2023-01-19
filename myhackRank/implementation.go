package myhackerRank

func GradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] > 37 {
			if (grades[i]+1)%5 == 0 {
				grades[i] = grades[i] + 1
			}
			if (grades[i]+2)%5 == 0 {
				grades[i] = grades[i] + 2
			}
		}
	}
	return grades
}

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	var s string
	if v1 < v2 {
		return "NO"
	}
	for {
		x1 += v1
		x2 += v2
		if x1 > x2 {
			s = "NO"
			break
		}
		if x1 == x2 {
			s = "YES"
			break
		}
	}
	return s
}

func BreakingRecords(scores []int32) []int32 {
	// Write your code here
	r := []int32{0, 0}
	highscores := scores[0]
	lowscores := scores[0]
	for i := 0; i < len(scores); i++ {
		if highscores < scores[i] {
			highscores = scores[i]
			r[0] += 1
		}
		if lowscores > scores[i] {
			lowscores = scores[i]
			r[1] += 1
		}
	}
	return r
}

// func SubararyDivision(s []int32, d, m int32) int32 {
// 	i := int(m)
// 	for i <= len(s) {

// 		for j := int32(1); j <= m-1; j++ {
// 			fmt.Printf("%v+%v\n", s[i], s[i+j])
// 			s[i] += s[i+j]
// 		}
// 		fmt.Printf("\n")
// 		if s[i] == d {
// 			x += 1
// 		}
// 		if x == m {
// 			break
// 		}

// 		i++
// 	}
// 	return x
// }

func SubararyDivision2(s []int32, d, m int32) int32 {
	i := 1
	j := int(m)
	count := int32(0)
	for j <= len(s) {
		i = 0
		for x := int32(0); x <= m-1; x++ {
			i += int(s[x])
		}
		if i == int(d) {
			count++
		}
		j++
	}

	return count
}
