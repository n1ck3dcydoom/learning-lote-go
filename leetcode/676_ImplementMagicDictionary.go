package leetcode

type MagicDictionary struct {
	words []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	n := len(dictionary)
	this.words = make([]string, n)
	for i := range dictionary {
		this.words[i] = dictionary[i]
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range this.words {
		// 长度不相等的字符串肯定不能通过替换一个字符得到
		if len(word) != len(searchWord) {
			continue
		}
		// 寻找不相等的字符个数,有且等于 1 个时才返回 true
		wordRunes := []rune(word)
		searchWordRunes := []rune(searchWord)

		count := 0
		for i := 0; i < len(wordRunes); i++ {
			// 提前剪枝
			if count > 1 {
				break
			}
			if wordRunes[i] != searchWordRunes[i] {
				count++
			}
		}
		if count == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
