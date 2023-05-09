package leetcode_golang

func countTime(time string) int {

	var minute = 1

	var hour = 1

	if time[1] == '?' && time[0] == '?' {
		hour = 24
	}
	if time[1] != '?' && time[0] == '?' {
		if time[1] > '3' {
			hour = 2
		} else {
			hour = 3
		}
	}

	if time[1] == '?' && time[0] != '?' {
		if time[0] == '2' {
			hour = 4
		} else {
			hour = 10
		}
	}

	// 0 1 2 3 4 5
	// 0 1 2 3 4 5 6 7 8 9
	if time[3] == '?' && time[4] == '?' {
		minute = 60
	}

	if time[3] != '?' && time[4] == '?' {
		minute = 10
	}

	if time[3] == '?' && time[4] != '?' {
		minute = 6
	}

	// fmt.Println(minute, hour)

	return minute * hour
}
