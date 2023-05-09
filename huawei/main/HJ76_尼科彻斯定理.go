package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}
// 	ans := make([]string, N)
// 	// 找规律发现，展开后的第一项 a1 = N*N-(N-1)，项数为 N，公差为 2
// 	for i := 0; i < N; i++ {
// 		ans[i] = strconv.FormatInt(int64(N*N-(N-1)+2*i), 10)
// 	}
// 	fmt.Println(strings.Join(ans, "+"))
// }
