package leetcode_golang

func isWinner(player1 []int, player2 []int) int {
	if getScore(player1) > getScore(player2) {
		return 1
	} else if getScore(player1) < getScore(player2) {
		return 2
	}
	return 0
}

func getScore(scores []int) int {
	sum := scores[0]
	if len(scores) == 1 {
		return sum
	}
	if len(scores) >= 2 {
		if sum == 10 {
			sum += 2 * scores[1]
		} else {
			sum += scores[1]
		}
		if len(scores) == 2 {
			return sum
		}
		for i := 2; i < len(scores); i++ {
			if scores[i-1] == 10 || scores[i-2] == 10 {
				sum += 2 * scores[i]
			} else {
				sum += scores[i]
			}
		}
	}
	return sum
}
