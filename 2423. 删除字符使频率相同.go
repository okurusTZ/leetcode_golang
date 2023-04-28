package leetcode_golang

func equalFrequency(word string) bool {

next:
	for i := 0; i < len(word); i++ {
		hashMap := make([]int, 26)
		for j, w := range word {
			if j != i {
				hashMap[w-'a']++
			}
		}

		var c = 0

		// fmt.Println(hashMap)

		for _, time := range hashMap {
			if c == 0 {
				c = time
			} else if c != time && time != 0 {
				continue next
			}
		}
		return true
	}
	return false
}
