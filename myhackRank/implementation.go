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
