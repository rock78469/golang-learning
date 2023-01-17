package router

import "go-learning/controller/hackrank"

func setHackRankRouter(testtype string, num int) {
	switch testtype {
	case "warmup":
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
		}
	case "implementation":
		switch num {
		case 1: // Grading Students
			hackrank.GradingStudents()
		case 2: // NumberLineJumps
			hackrank.NumberLineJumps()
		case 3: // BreakingRecords
			hackrank.BreakingRecords()
		}
	}
}
