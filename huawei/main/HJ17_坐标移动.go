package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// func main() {
// 	inputs := ""
// 	n, _ := fmt.Scan(&inputs)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	steps := strings.Split(inputs, ";")
// 	x, y := 0, 0
// 	for _, s := range steps {
// 		if check(s) {
// 			s1 := s[0]
// 			s2, _ := strconv.Atoi(s[1:])
// 			if s1 == 'A' {
// 				x -= s2
// 			} else if s1 == 'D' {
// 				x += s2
// 			} else if s1 == 'W' {
// 				y += s2
// 			} else if s1 == 'S' {
// 				y -= s2
// 			}
// 		}
// 	}
// 	fmt.Printf("%d,%d", x, y)
// }

// func check(s string) bool {
// 	if len(s) != 3 && len(s) != 2 {
// 		return false
// 	}
// 	rs := []rune(s)

// 	s1 := rs[0]
// 	if s1 != 'A' && s1 != 'S' && s1 != 'D' && s1 != 'W' {
// 		return false
// 	}
// 	s2 := rs[1:]
// 	for i := 0; i < len(s2); i++ {
// 		if !unicode.IsDigit(s2[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }
