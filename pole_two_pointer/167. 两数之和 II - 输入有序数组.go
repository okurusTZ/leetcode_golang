package pole_two_pointer

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
		if sum == target {
			return []int{left + 1, right + 1}
		}
	}
	return []int{left + 1, right + 1}
}
