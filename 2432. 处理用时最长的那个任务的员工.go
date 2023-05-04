package leetcode_golang

func hardestWorker(n int, logs [][]int) int {
	var (
		ans     = 0
		maxTime = 0
	)

	for i, log := range logs {
		id, time := log[0], log[1]
		if i == 0 {
			ans = id
			maxTime = time
		} else {
			var spendTime = time - logs[i-1][1]
			if spendTime > maxTime {
				ans = id
				maxTime = spendTime
			}
			if spendTime == maxTime && id < ans {
				ans = id
			}
		}
	}

	return ans
}
