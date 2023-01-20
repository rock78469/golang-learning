package router

func SetRouter(site, testType string, num int) {
	switch site {
	case "hackrank":
		setHackRankRouter(testType, num)
	case "leetcode":
		switch num {
		case 383:
		case 412:
		case 1480:
		}
	}
}
