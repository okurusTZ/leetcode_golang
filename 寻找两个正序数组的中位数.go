package leetcode_golang

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n))

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	odd := false // 奇数
	if (len(nums1)+len(nums2))%2 != 0 {
		odd = true
	}

	// 提前存一份，注意写在里面的话length是会变的
	length := (len(nums1) + len(nums2))

	idx := 0
	head := 0

	for idx <= length/2+1 {

		var hit = 0

		if len(nums1) > 0 && len(nums2) > 0 {
			if nums1[0] > nums2[0] {
				hit = nums2[0]
				nums2 = nums2[1:]
			} else {
				hit = nums1[0]
				nums1 = nums1[1:]
			}
		} else {
			if len(nums1) > 0 {
				hit = nums1[0]
				nums1 = nums1[1:]
			}
			if len(nums2) > 0 {
				hit = nums2[0]
				nums2 = nums2[1:]
			}
		}

		idx++

		if idx == length/2 && !odd {
			head += hit
		}
		if idx == length/2+1 {
			head += hit
		}
	}

	if odd {
		return float64(head)
	}
	return float64(head) / float64(2)
}
