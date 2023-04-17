package simulate

import (
	"learning-lote-go/leetcode/base"
	"strconv"
	"strings"
)

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	// 非闰年，一年 365 天，直接暴力枚举每一天
	monthDayArray := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// 通过前缀和数组快速得到 MM-DD 是一年当中的第几天
	// preSum[MM]+DD
	preSum := make([]int, 12+1)
	preSum[0] = 0
	for i := 1; i <= 12; i++ {
		preSum[i] = preSum[i-1] + monthDayArray[i-1]
	}

	arriveAliceSplits := strings.Split(arriveAlice, "-")
	arriveAliceMonth, _ := strconv.Atoi(arriveAliceSplits[0])
	arriveAliceDay, _ := strconv.Atoi(arriveAliceSplits[1])
	arriveAliceAt := preSum[arriveAliceMonth-1] + arriveAliceDay

	leaveAliceSplits := strings.Split(leaveAlice, "-")
	leaveAliceMonth, _ := strconv.Atoi(leaveAliceSplits[0])
	leaveAliceDay, _ := strconv.Atoi(leaveAliceSplits[1])
	leaveAliceAt := preSum[leaveAliceMonth-1] + leaveAliceDay

	arriveBobSplits := strings.Split(arriveBob, "-")
	arriveBobMonth, _ := strconv.Atoi(arriveBobSplits[0])
	arriveBobDay, _ := strconv.Atoi(arriveBobSplits[1])
	arriveBobAt := preSum[arriveBobMonth-1] + arriveBobDay

	leaveBobSplits := strings.Split(leaveBob, "-")
	leaveBobMonth, _ := strconv.Atoi(leaveBobSplits[0])
	leaveBobDay, _ := strconv.Atoi(leaveBobSplits[1])
	leaveBobAt := preSum[leaveBobMonth-1] + leaveBobDay

	// 计算 [aaa,lat] 和 [aba,lba] 两个区间的交集
	if leaveAliceAt < arriveBobAt || leaveBobAt < arriveAliceAt {
		return 0
	}
	return base.MinInt(leaveAliceAt, leaveBobAt) - base.MaxInt(arriveAliceAt, arriveBobAt) + 1
}
