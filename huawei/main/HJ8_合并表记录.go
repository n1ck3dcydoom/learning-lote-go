package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	hash := make(map[int]int)
// 	for i := 0; i < N; i++ {
// 		key, val := 0, 0
// 		_, _ = fmt.Scan(&key, &val)
// 		if v, ok := hash[key]; ok {
// 			hash[key] = val + v
// 		} else {
// 			hash[key] = val
// 		}
// 	}
// 	array := make([]Pair, 0)
// 	for k, v := range hash {
// 		array = append(array, Pair{K: k, V: v})
// 	}
// 	sort.Slice(array, func(i, j int) bool {
// 		return array[i].K < array[j].K
// 	})

// 	for i := range array {
// 		fmt.Printf("%d %d\n", array[i].K, array[i].V)
// 	}
// }

// type Pair struct {
// 	K int
// 	V int
// }
