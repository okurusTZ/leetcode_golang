package prefix_sum

func subarraySum(nums []int, k int) int {
	var (
		prefixSum = make([]int, len(nums))
		hashMap   = make(map[int]int)
		ans       = 0
	)

	hashMap[0] = 1 // 必须初始化的数值

	for i, num := range nums {
		if i == 0 {
			prefixSum[i] = num
		} else {
			prefixSum[i] = prefixSum[i-1] + num
		}

		if hashMap[prefixSum[i]-k] > 0 {
			ans += hashMap[prefixSum[i]-k]
		}
		hashMap[prefixSum[i]]++
	}

	return ans
}

// o(n^2)超时
func subarraySum2(nums []int, k int) int {
	var (
		prefixSum = make([]int, len(nums))
		ans       = 0
	)

	for i, num := range nums {
		if i == 0 {
			prefixSum[i] = num
		} else {
			prefixSum[i] = prefixSum[i-1] + num
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if prefixSum[j]-prefixSum[i]+nums[i] == k {
				ans++
			}
		}
	}
	return ans
}
