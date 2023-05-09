package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	s := scanner.Text()
// 	rs := []rune(s)

// 	letter := 0
// 	space := 0
// 	digit := 0
// 	other := 0

// 	for i := range rs {
// 		if 'a' <= rs[i] && rs[i] <= 'z' || 'A' <= rs[i] && rs[i] <= 'Z' {
// 			letter++
// 		} else if rs[i] == ' ' {
// 			space++
// 		} else if '0' <= rs[i] && rs[i] <= '9' {
// 			digit++
// 		} else {
// 			other++
// 		}
// 	}

// 	fmt.Printf("%d\n%d\n%d\n%d", letter, space, digit, other)
// }
