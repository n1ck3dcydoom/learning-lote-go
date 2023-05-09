package main

// import (
// 	"fmt"
// 	"strconv"
// 	"unicode"
// )

// func main() {
// 	// 1CNY=100fen
// 	// 1JPY=100sen
// 	// 1HKD=100cents
// 	// 1EUR=100eurocents
// 	// 1GBP=100pence

// 	// CNY	JPY  HKD	EUR	GBP
// 	// 100	1825 123	14	12

// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	hash := make(map[string]float64)
// 	hash["CNY"] = 100.0
// 	hash["fen"] = 1.0
// 	hash["JPY"] = 100.0 / 1825.0 * 100.0
// 	hash["sen"] = 100.0 / 1825.0
// 	hash["HKD"] = 100.0 / 123.0 * 100.0
// 	hash["cents"] = 100.0 / 123.0
// 	hash["EUR"] = 100.0 / 14.0 * 100.0
// 	hash["eurocents"] = 100.0 / 14.0
// 	hash["GBP"] = 100.0 / 12.0 * 100.0
// 	hash["pence"] = 100.0 / 12.0

// 	ms := make([]string, N)
// 	amount := 0
// 	currency := ""
// 	ans := 0.0
// 	for i := 0; i < N; i++ {
// 		_, _ = fmt.Scan(&ms[i])
// 		rms := []rune(ms[i])
// 		k := 0
// 		// 得到金额
// 		for ; k < len(rms); k++ {
// 			if !unicode.IsDigit(rms[k]) {
// 				tmp := rms[:k]
// 				amount, _ = strconv.Atoi(string(tmp))
// 				break
// 			}
// 		}
// 		// 得到币种
// 		currency = string(rms[k:])
// 		// 计算总价
// 		ans += hash[currency] * float64(amount)
// 	}
// 	fmt.Println(ans)
// }
