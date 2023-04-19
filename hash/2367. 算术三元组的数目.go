package hashmap

func arithmeticTriplets(nums []int, diff int) int {
	// 严格递增 -> 没有相同的数字

	hashMap := make(map[int]int)
	ans := 0

	for idx, num := range nums {
		hashMap[num] = idx + 1
		if hashMap[num-diff] > 0 && hashMap[num-diff*2] > 0 {
			ans++
		}
	}

	return ans
}

// 这里有一个tricky的地方，就是时间复杂度其实是O(n)
func arithmeticTriplets2(nums []int, diff int) int {
	// 三指针
	var (
		i   = 0
		j   = 1
		res = 0
	)

	// 对于i，j来说，虽然有两层循环，但是只会从0-n/n-1/n-2 所以实际上是O(n)

	for k := 0; k < len(nums); k++ {
		for i < len(nums) && nums[i] < nums[k]+diff {
			i++
		}
		if i >= len(nums) || nums[i] != nums[k]+diff {
			continue
		}

		j = i

		for j < len(nums) && nums[j] < nums[k]+2*diff {
			j++
		}
		if j >= len(nums) || nums[j] != nums[k]+2*diff {
			continue
		}
		res++
	}

	return res
}
