package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	ip := ""
// 	num := 0
// 	n, _ := fmt.Scan(&ip, &num)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	// 将 ip 当做 256 进制处理
// 	segs := strings.Split(ip, ".")
// 	ip2numRes := int64(0)
// 	carry := 3
// 	for i := range segs {
// 		num, _ := strconv.ParseInt(segs[i], 10, 64)
// 		ip2numRes += int64(math.Pow(float64(256), float64(carry)) * float64(num))
// 		carry--
// 	}
// 	fmt.Println(ip2numRes)

// 	// 将 10 进制数转换为 256 进制的 ip
// 	num2ipRes := make([]string, 4)
// 	// 短除法是逆序的
// 	for i := 3; i >= 0; i-- {
// 		num2ipRes[i] = strconv.FormatInt(int64(num%256), 10)
// 		num /= 256
// 	}
// 	fmt.Println(strings.Join(num2ipRes, "."))
// }

// // func main() {
// // 	ip := ""
// // 	num := 0
// // 	n, _ := fmt.Scan(&ip, &num)
// // 	if n == 0 {
// // 		fmt.Println("error")
// // 	}

// // 	// ip to num
// // 	segs := strings.Split(ip, ".")
// // 	nums := make([]rune, 0)
// // 	for i := 0; i < len(segs); i++ {
// // 		nums = append(nums, segto2(segs[i])...)
// // 	}
// // 	ans, _ := strconv.ParseInt(string(nums), 2, 64)
// // 	fmt.Println(ans)

// // 	// num to ip
// // 	res := make([]string, 4)
// // 	count := 0
// // 	index := 3
// // 	tmp := make([]rune, 0)
// // 	for num != 0 {
// // 		if count == 8 {
// // 			count = 0
// // 			// 逆向
// // 			rtmp := make([]rune, 0)
// // 			for i := len(tmp) - 1; i >= 0; i-- {
// // 				rtmp = append(rtmp, tmp[i])
// // 			}
// // 			ans, _ := strconv.ParseInt(string(rtmp), 2, 64)
// // 			res[index] = strconv.FormatInt(ans, 10)
// // 			index--
// // 			tmp = make([]rune, 0)
// // 		} else {
// // 			c := num % 2
// // 			num /= 2
// // 			tmp = append(tmp, rune(c)+'0')
// // 			count++
// // 		}
// // 	}
// // 	if len(tmp) != 0 {
// // 		// 逆向
// // 		rtmp := make([]rune, 0)
// // 		for i := len(tmp) - 1; i >= 0; i-- {
// // 			rtmp = append(rtmp, tmp[i])
// // 		}
// // 		ans, _ := strconv.ParseInt(string(rtmp), 2, 64)
// // 		res[index] = strconv.FormatInt(ans, 10)
// // 	}
// // 	fmt.Println(strings.Join(res, "."))
// // }

// // func segto2(seg string) (ans []rune) {
// // 	n, _ := strconv.Atoi(seg)
// // 	tmp := make([]rune, 0)
// // 	for n != 0 {
// // 		c := n % 2
// // 		tmp = append(tmp, rune(c)+'0')
// // 		n /= 2
// // 	}
// // 	// 补 0
// // 	diff := 8 - len(tmp)
// // 	for i := 0; i < diff; i++ {
// // 		tmp = append(tmp, '0')
// // 	}
// // 	// 逆向
// // 	for i := len(tmp) - 1; i >= 0; i-- {
// // 		ans = append(ans, tmp[i])
// // 	}
// // 	return ans
// // }
