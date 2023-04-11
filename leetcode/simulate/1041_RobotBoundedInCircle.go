package simulate

func IsRobotBounded(instructions string) bool {
	// 模拟走完所有指令
	// 1. 如果最后回到原点(0,0)，则陷入死循环
	// 2. 最后没有回到原点，此时判断方向
	// 2.1 方向朝北(x,y)，下一次执行指令后来到（2x,2y)仍然朝北，不会陷入死循环
	// 2.2 方向朝南(-x,-y)，下一次执行指令后反向回到(0,0)，陷入死循环
	// 2.3 方向朝东或者朝西，这样第 1 次和第 3 次执行都互为相反；第 2 次和第 3 次执行互为相反，陷入死循环

	// 类似于网格中的 bfs 或者 dfs 搜索算法，记录 4 个搜索向量
	// 为了方便左转右转，这里采用顺时针顺序记录
	// --------------上，------右，------下，------左
	ways := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 记录当前的方向，实际记录的是 ways 数组里面的下标
	wayIndex := 0
	// 记录坐标
	x, y := 0, 0
	// 遍历指令集
	for _, i := range instructions {
		if i == 'G' {
			x += ways[wayIndex][0]
			y += ways[wayIndex][1]
		} else if i == 'R' {
			// 右转
			wayIndex++
			// 防止数组越界
			wayIndex %= 4
		} else if i == 'L' {
			// 左转
			// 防止变为负数，先加上数组长度后再减一
			wayIndex = wayIndex + 4 - 1
			wayIndex %= 4
		}
	}
	// 回到原点或者方向不朝北，都会导致陷入死循环
	return (x == 0 && y == 0) || (wayIndex != 0)
}
