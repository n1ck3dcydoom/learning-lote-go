package codetop

import (
	"container/list"
)

func IsValid(s string) bool {
	// 栈
	stack := list.New()

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack.PushBack(c)
		} else {
			if stack.Len() == 0 {
				return false
			} else if c == ')' {
				pop := stack.Back()
				stack.Remove(pop)
				if pop.Value != '(' {
					return false
				}
			} else if c == ']' {
				pop := stack.Back()
				stack.Remove(pop)
				if pop.Value != '[' {
					return false
				}
			} else if c == '}' {
				pop := stack.Back()
				stack.Remove(pop)
				if pop.Value != '{' {
					return false
				}
			}
		}
	}
	return stack.Len() == 0
}
