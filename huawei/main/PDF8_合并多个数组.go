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
// 	inputs := make([]string, N)
// 	nums := make([][]int, 0)
// 	total := 0
// 	for i := 0; i < N; i++ {
// 		_, _ = fmt.Scan(&inputs[i])
// 		inputSplit := strings.Split(inputs[i], ",")
// 		// 冗余最后一个位置记录当前数组遍历的下标
// 		num := make([]int, len(inputSplit)+1)
// 		for j := 0; j < len(inputSplit); j++ {
// 			num[j], _ = strconv.Atoi(inputSplit[j])
// 		}
// 		nums = append(nums, num)
// 		total += len(num)
// 	}
// 	fmt.Println(len(nums))

// 	ans := make([]int, 0)
// 	count := 0
// 	// 因为 nums 每个 num 都多冗余了一个元素，计数时需要去掉冗余元素
// 	for count < total-len(nums) {
// 		for i := range nums {
// 			num := nums[i]
// 			// 判断当前数组是否完成了遍历
// 			if num[len(num)-1] == len(num)-1 {
// 				continue
// 			}
// 			// 取当前位置开始的 N 个元素，不够 N 个元素取完后停止遍历
// 			for k := 0; k < N && num[len(num)-1] < len(num)-1; k++ {
// 				ans = append(ans, num[num[len(num)-1]])
// 				num[len(num)-1]++
// 				count++
// 			}
// 		}
// 	}
// 	ansString := make([]string, 0)
// 	for i := 0; i < len(ans); i++ {
// 		ansString = append(ansString, strconv.FormatInt(int64(ans[i]), 10))
// 	}
// 	fmt.Println(strings.Join(ansString, ","))
// }
