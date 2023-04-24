package slide_window

func checkInclusion(s1 string, s2 string) bool {

	hashMap := make([]int, 26)
	n := len(s1)
	for _, s := range s1 {
		hashMap[s-'a']++
	}

	var (
		window = make([]int, 26)
	)

	if len(s2) < len(s1) {
		return false
	}

	// init
	for i := 0; i < n; i++ {
		window[s2[i]-'a']++
		if SameSlice(hashMap, window) {
			return true
		}
	}

	for i := 1; i+n-1 < len(s2); i++ {
		window[s2[i-1]-'a']--
		window[s2[i+n-1]-'a']++
		if SameSlice(hashMap, window) {
			return true
		}

	}
	return false
}

func SameSlice(a, b []int) bool {
	for i, num := range a {
		if num == 0 {
			continue
		}
		if b[i] != num {
			return false
		}
	}
	return true
}
