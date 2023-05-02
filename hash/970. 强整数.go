package hashmap

func powerfulIntegers2(x int, y int, bound int) []int {
	var (
		xMap     = make(map[int]bool)
		yMap     = make(map[int]bool)
		xProduct = x
		yProduct = y
		ans      = make([]int, 0)
	)

	// init
	xMap[1] = true
	yMap[1] = true
	xMap[x] = true
	yMap[y] = true

	for i := 2; i <= bound && (xProduct <= bound || yProduct <= bound); i++ {

		if xProduct < bound {
			xProduct = xProduct * x
			xMap[xProduct] = true
		}
		if yProduct < bound {
			yProduct = yProduct * y
			yMap[yProduct] = true
		}
	}

	for i := 2; i <= bound; i++ {
		for x := range xMap {
			if i == x {
				continue
			}
			if yMap[i-x] {
				ans = append(ans, i)
				break
			}
		}
	}

	return ans
}

func powerfulIntegers(x, y, bound int) []int {

	var hashMap = make(map[int]bool)
	var ans = make([]int, 0)

	for a := 1; a <= bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			hashMap[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	for x := range hashMap {
		ans = append(ans, x)
	}
	return ans
}
