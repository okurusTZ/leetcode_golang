package slide_window

func longestOnes(nums []int, k int) int {

	left := 0
	ans := 0

	countZero := 0

	for right := 0; right < len(nums); right++ {

		if nums[right] == 0 {
			countZero++
		}

		for countZero > k {
			if nums[left] == 0 {
				countZero--
			}
			left++
		}

		ans = max(ans, right-left+1)
	}
	return ans
}

func main() {

}
