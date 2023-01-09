package router

import "go-learning/controller/hackrank"

func SetRouter(testType string, num int) {
	switch testType {
	case "hackrank":
		switch num {
		case 1: // Simple Array Sum
			hackrank.SimpleArraySum()
		case 4: // Diagonal Difference
			hackrank.DiagonalDifference()
		}
	case "leetcode":
		switch num {
		case 383:
		case 412:
		case 1480:
		}

	}
}
