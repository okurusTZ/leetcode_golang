package leetcode_golang

import "sort"

func minimumLines(stockPrices [][]int) int {

	if len(stockPrices) == 1 {
		// 边界，只有一个点，肯定是0
		return 0
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	pre_dx := 0
	pre_dy := 0
	// 第一个肯定会算作一条线

	sum := 0

	for idx := range stockPrices {
		if idx == 0 {
			continue
		}

		dx := stockPrices[idx][0] - stockPrices[idx-1][0]
		dy := stockPrices[idx][1] - stockPrices[idx-1][1]

		println(dx, dy)

		if dx*pre_dy != dy*pre_dx {
			sum += 1
		}

		pre_dx = dx
		pre_dy = dy
	}

	return sum
}
