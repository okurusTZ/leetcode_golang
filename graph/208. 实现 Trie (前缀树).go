package graph

type Trie struct {
	Children [26]*Trie // 代表26个字母
	IsEnd    bool      // 是否已经完整
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for _, s := range word {
		idx := s - 'a'
		if t.Children[idx] == nil {
			t.Children[idx] = &Trie{}
		}
		t = t.Children[idx] // 移动下去
	}
	t.IsEnd = true // 该字符串处理完成了，表示已经完整
}

func (t *Trie) Search(word string) bool {
	for _, s := range word {
		idx := s - 'a'
		if t.Children[idx] == nil {
			return false
		}
		t = t.Children[idx]
	}
	return t.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, s := range prefix {
		idx := s - 'a'
		if t.Children[idx] == nil {
			return false
		}
		t = t.Children[idx]
	}
	return true
}
