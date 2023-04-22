package prefix_sum

type NumArray struct {
	Arr       []int
	PrefixSum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			sum[i] = num
			continue
		}
		sum[i] = sum[i-1] + num
	}
	return NumArray{
		Arr:       nums,
		PrefixSum: sum,
	}
}

func (arr *NumArray) SumRange(left int, right int) int {
	return arr.PrefixSum[right] - arr.PrefixSum[left] + arr.Arr[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
