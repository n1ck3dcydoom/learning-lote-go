package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	line := scanner.Text()

// 	hash := make(map[rune]int)
// 	rline := []rune(line)
// 	for i := 0; i < len(rline); i++ {
// 		if _, ok := hash[rline[i]]; !ok {
// 			hash[rline[i]] = 1
// 		} else {
// 			hash[rline[i]] = hash[rline[i]] + 1
// 		}
// 	}

// 	pairs := make([]Pair, 0)
// 	for k, v := range hash {
// 		pairs = append(pairs, Pair{R: k, Count: v})
// 	}
// 	sort.Slice(pairs, func(i, j int) bool {
// 		if pairs[i].Count > pairs[j].Count {
// 			return true
// 		} else if pairs[i].Count < pairs[j].Count {
// 			return false
// 		} else {
// 			return pairs[i].R < pairs[j].R
// 		}
// 	})
// 	for _, p := range pairs {
// 		fmt.Printf("%c", p.R)
// 	}
// }

// type Pair struct {
// 	R     rune
// 	Count int
// }
