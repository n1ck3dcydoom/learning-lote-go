package main

// import (
// 	"fmt"
// )

// func main() {
// 	N, k := 0, 0
// 	n, _ := fmt.Scan(&N, &k)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	nums := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		_, _ = fmt.Scan(&nums[i])
// 	}
// 	// 我他妈听到排序就直接手写快排
// 	quickSort(nums, 0, len(nums)-1)
// 	for i := 0; i < k; i++ {
// 		if i < k-1 {
// 			fmt.Printf("%d ", nums[i])
// 		} else {
// 			fmt.Print(nums[i])
// 		}
// 	}
// }

// func quickSort(nums []int, l, r int) {
// 	if l > r {
// 		return
// 	}
// 	// 第一个元素作为哨兵
// 	o := l
// 	ll := l
// 	rr := r
// 	for ll < rr {
// 		// 哨兵在左边，先从右往左查找第一个小于哨兵的数
// 		for ll < rr && nums[rr] >= nums[o] {
// 			rr--
// 		}
// 		// 哨兵在左边，再从左往右查找第一个大于哨兵的数
// 		for ll < rr && nums[ll] <= nums[o] {
// 			ll++
// 		}
// 		if rr > ll {
// 			// 交换 ll 和 rr
// 			nums[ll], nums[rr] = nums[rr], nums[ll]
// 		}
// 	}
// 	// 交换哨兵
// 	nums[o], nums[ll] = nums[ll], nums[o]
// 	// 继续递归排序
// 	quickSort(nums, l, ll-1)
// 	quickSort(nums, ll+1, r)
// }
