package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := 0
	n, _ := fmt.Scan(&N)
	if n == 0 {
		fmt.Println("null")
	} else {
		ans := make([]int, 0)
		records := make([]string, N)
		// hash 保存每位员工的打卡记录，由于需要按照输入顺序输出，所以这里保存打卡记录的下标 i
		hash := make(map[string][]int)
		for i := 0; i < N; i++ {
			_, _ = fmt.Scan(&records[i])
			r := strings.Split(records[i], ",")
			who := r[0]
			d1 := r[3]
			d2 := r[4]
			if d1 == d2 {
				hash[who] = append(hash[who], i)
			} else {
				ans = append(ans, i)
			}
		}

		// 对于那些正常的打卡记录，需要对每个人检查连续两条记录是否间隔小于 60 分钟且打卡距离小于超过 5 km
		// 2 工号，时间，距离，登记设备，打卡设备
		// 100000,10,1,ABCD,ABCD
		// 100000,80,10,ABCE,ABCD
		for _, v := range hash {
			whoRecords := v
			// 检查两个打卡记录的时候，可能重复添加，再使用一个 map 去重
			set := make(map[int]int)
			for i := 0; i < len(whoRecords); i++ {
				r1 := strings.Split(records[whoRecords[i]], ",")
				t1, _ := strconv.Atoi(r1[1])
				p1, _ := strconv.Atoi(r1[2])
				for j := i + 1; j < len(whoRecords); j++ {
					r2 := strings.Split(records[whoRecords[j]], ",")
					t2, _ := strconv.Atoi(r2[1])
					p2, _ := strconv.Atoi(r2[2])
					// 打卡距离可能 p2 > p1 也可能 p1 > p2
					if abs(t2-t1) < 60 && abs(p2-p1) > 5 {
						// 由于 hash 里面保存的每个人的打卡记录不是按照输入顺序排序
						// 这里的 ans 不要直接保存打卡记录，而是记录打卡记录的下标
						if _, ok := set[whoRecords[i]]; !ok {
							ans = append(ans, whoRecords[i])
							set[whoRecords[i]] = whoRecords[i]
						}
						if _, ok := set[whoRecords[j]]; !ok {
							ans = append(ans, whoRecords[j])
							set[whoRecords[j]] = whoRecords[j]
						}
					}
				}
			}
		}

		if len(ans) == 0 {
			fmt.Println("null")
		} else {
			// 通过 ans 记录的下标，排序后检索 records 数组再输出
			sort.Slice(ans, func(i, j int) bool {
				return ans[i] < ans[j]
			})
			realAns := make([]string, 0)
			for i := 0; i < len(ans); i++ {
				realAns = append(realAns, records[ans[i]])
			}
			fmt.Print(strings.Join(realAns, ";"))
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
