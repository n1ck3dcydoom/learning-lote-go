package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	a := 0.3
// 	// n, _ := fmt.Scan(&a)
// 	// if n == 0 {
// 	// fmt.Println("error")
// 	// }
// 	// 直接二分查找
// 	l := 0.0
// 	r := math.Max(1.0, a)
// 	if a < 0 {
// 		// 调整负数区间
// 		l = math.Min(-1.0, a)
// 		r = 0
// 	}

// 	m := 0.0
// 	for math.Abs(r-l) >= 0.001 {
// 		m = l + (r-l)/2
// 		if m*m*m > a {
// 			r = m
// 		} else if m*m*m < a {
// 			l = m
// 		}
// 	}
// 	fmt.Printf("%.1f", m)
// }
