package main

import (
	"math"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sents := strings.Split(sentence, " ")
	res := make([]string, len(sents))
	for idx, s := range sents {
		minLen := math.MaxInt
		hit := false
		for _, dic := range dictionary {
			if strings.HasPrefix(s, dic) && len(dic) < minLen {
				minLen = len(dic)
				res[idx] = dic
				hit = true
			}
		}
		if !hit {
			res[idx] = s
		}
	}
	return strings.Join(res, " ")
}

func replaceWords2(dictionary []string, sentence string) string {
	type dictTree map[int32]dictTree
	root := dictTree{}

	for _, dict := range dictionary {
		cur := root
		for _, c := range dict {
			if cur[c] == nil {
				cur[c] = dictTree{}
			}
			cur = cur[c]
		}
		cur['#'] = dictTree{}
	}

	sentences := strings.Split(sentence, " ")

	for idx, s := range sentences {
		cur := root
		for i, c := range s {
			// 优先判断这个，说明匹配一个完整的前缀了
			if cur['#'] != nil {
				sentences[idx] = s[:i]
				break
			}
			// 第一个没有的时候，直接停止循环 -> 是匹配前缀，不是是否包含
			// 防止baba的case
			if cur[c] == nil {
				break
			}
			cur = cur[c]
		}
	}
	return strings.Join(sentences, " ")
}
