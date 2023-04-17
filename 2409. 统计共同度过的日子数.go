package main

import (
	"strconv"
	"strings"
)

// arriveAlice = "08-15", leaveAlice = "08-18", arriveBob = "08-16", leaveBob = "08-19"

var date = [13]int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	m1, d1, m2, d2 := strings.Split(arriveAlice, "-")[0], strings.Split(arriveAlice, "-")[1], strings.Split(leaveAlice, "-")[0], strings.Split(leaveAlice, "-")[1]
	m3, d3, m4, d4 := strings.Split(arriveBob, "-")[0], strings.Split(arriveBob, "-")[1], strings.Split(leaveBob, "-")[0], strings.Split(leaveBob, "-")[1]

	var dd1, dd2, dd3, dd4 int = 0, 0, 0, 0

	m11, _ := strconv.Atoi(m1)
	d11, _ := strconv.Atoi(d1)
	m22, _ := strconv.Atoi(m2)
	d22, _ := strconv.Atoi(d2)
	m33, _ := strconv.Atoi(m3)
	d33, _ := strconv.Atoi(d3)
	m44, _ := strconv.Atoi(m4)
	d44, _ := strconv.Atoi(d4)

	// 这里统计的时候可以用前缀和优化 -> 但我懒得改了
	for i := 1; i < len(date); i++ {
		if i < m11 {
			dd1 += date[i]
		}

		if i < m22 {
			dd2 += date[i]
		}

		if i < m33 {
			dd3 += date[i]
		}

		if i < m44 {
			dd4 += date[i]
		}

		if i == m11 {
			dd1 += d11
		}

		if i == m22 {
			dd2 += d22
		}

		if i == m33 {
			dd3 += d33
		}

		if i == m44 {
			dd4 += d44
		}
	}

	// 离开的天数里最小的 - 到达的天数里最大的
	return max(0, min(dd2, dd4)-max(dd1, dd3)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
