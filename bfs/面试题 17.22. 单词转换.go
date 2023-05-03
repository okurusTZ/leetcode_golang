package bfs

// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]

func findLadders(beginWord string, endWord string, wordList []string) []string {
	var (
		queue  = make([]string, 0)
		ans    = make([]string, 0)
		path   = make(map[string]string)
		inList = make(map[string]bool)
	)

	for _, item := range wordList {
		inList[item] = true
	}

	queue = append(queue, beginWord)

	for len(queue) > 0 {
		// 先把所有的都遍历了
		for k := 0; k < len(queue); k++ {
			currentWord := queue[0]
			// fmt.Println(currentWord, endWord, queue)
			if currentWord == endWord {
				ans = append([]string{endWord}, ans...)
				// 开始回溯
				s := endWord
				if s != beginWord {
					s = path[s] // 记录上一个
					ans = append([]string{s}, ans...)
				}
				return ans
			}

			for idx := range currentWord {
				for j := 'a'; j <= 'z'; j++ {
					// 下一个需要遍历的情况
					newWord := currentWord[:idx] + string(j) + currentWord[idx+1:]
					if inList[newWord] {
						path[newWord] = currentWord // 记录一下当前路径
						queue = append(queue, newWord)
						inList[newWord] = false // 防止重复访问
					}
				}
			}
			// fmt.Println(queue)
			queue = queue[1:]
		}
	}

	return ans
}
