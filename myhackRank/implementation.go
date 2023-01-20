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

func SubararyDivision(s []int32, d, m int32) int32 {
	var count int32
	for i := 0; i <= len(s)-int(m); i++ {
		for x := 1; x < int(m); x++ {
			s[i] += s[i+x]
		}
		if s[i] == d {
			count++
		}
	}

	return count
}
