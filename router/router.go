package router

import "go-learning/controller/hackrank"

func SetRouter(testType string, num int) {
	switch testType {
	case "hackrank":
		switch num {
		case 1: // Simple Array Sum
			hackrank.SimpleArraySum()
		case 2: // ComparetheTriplets
			hackrank.ComparetheTriplets()
		case 3: // A Very Big Sum
			hackrank.AVeryBigSum()
		case 4: // Diagonal Difference
			hackrank.DiagonalDifference()
		case 5: // Birthday Cake Candles
			hackrank.BirthdayCakeCandles()
		case 6: // Time Conversion
			hackrank.TimeConversion()
		case 7: // Grading Students
			hackrank.GradingStudents()
		}
	case "leetcode":
		switch num {
		case 383:
		case 412:
		case 1480:
		}

	}
}
