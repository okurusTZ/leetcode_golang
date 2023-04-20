package leetcode_golang

import (
	"fmt"
	"strconv"
)

func convertTemperature(celsius float64) []float64 {
	tmp1, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", celsius+273.15), 64)
	tmp2, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", celsius*1.80+32.00), 64)
	return []float64{tmp1, tmp2}
}
