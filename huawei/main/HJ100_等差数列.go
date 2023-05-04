package main

import "fmt"

func main() {
	N := 0
	n, _ := fmt.Scan(&N)
	if n == 0 {
		fmt.Println("error")
	}
	// 首项为 2，公差为 3，求前 n 项和
	a1 := 2
	d := 3
	// 等差数列公式
	// sum=(首项+末项)*项数÷2
	// 项数=(末项-首项)÷公差+1
	// an=a1+(n-1)*d

	// 先求第 n 项
	an := a1 + (N-1)*d
	// 接着求项数
	n = (an-a1)/d + 1
	// 最后求和
	sum := (a1 + an) * n / 2
	fmt.Println(sum)
}
