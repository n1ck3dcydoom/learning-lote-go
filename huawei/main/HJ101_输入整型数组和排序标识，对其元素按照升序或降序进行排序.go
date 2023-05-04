package main

// import "fmt"

// func main() {
// 	N := 0
// 	n, _ := fmt.Scan(&N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	nums := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		_, _ = fmt.Scan(&nums[i])
// 	}

// 	// 0 升序 1 降序
// 	sort := 0
// 	_, _ = fmt.Scan(&sort)

// 	quickSort(nums, 0, len(nums)-1, sort)

// 	for i := range nums {
// 		fmt.Printf("%d ", nums[i])
// 	}
// }

// // 来个手写快排吧 ˊ_>ˋ
// func quickSort(nums []int, l, r, sort int) {
// 	if l >= r {
// 		return
// 	}
// 	// l 作为哨兵
// 	o := l
// 	ql := l
// 	qr := r
// 	for ql < qr {
// 		if sort == 0 {
// 			// 从后往前查找第一个小于哨兵的元素
// 			for qr > ql && nums[qr] >= nums[o] {
// 				qr--
// 			}
// 			// 从前往后查找第一个大于哨兵的元素
// 			for qr > ql && nums[ql] <= nums[o] {
// 				ql++
// 			}
// 		} else {
// 			// 从后往前查找第一个大于哨兵的元素
// 			for qr > ql && nums[qr] <= nums[o] {
// 				qr--
// 			}
// 			// 从前往后查找第一个小于哨兵的元素
// 			for qr > ql && nums[ql] >= nums[o] {
// 				ql++
// 			}
// 		}
// 		// 交换 l 和 r 的位置
// 		nums[qr], nums[ql] = nums[ql], nums[qr]
// 	}
// 	// 交换哨兵位置
// 	nums[o], nums[ql] = nums[ql], nums[o]
// 	// 进入下一趟排序
// 	quickSort(nums, l, ql-1, sort)
// 	quickSort(nums, ql+1, r, sort)
// }
