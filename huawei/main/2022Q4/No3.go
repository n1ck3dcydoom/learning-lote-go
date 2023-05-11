package main

// import (
// 	"container/list"
// 	"fmt"
// )

// func main() {
// 	M, N := 0, 0
// 	n, _ := fmt.Scan(&M, &N)
// 	if n == 0 {
// 		fmt.Println("error")
// 	}

// 	// 初始化网格
// 	// 6 5
// 	// 0 0 0 -1 0 0 0 0 0 0 0 0 -1 4 0 0 0 0 0 0 0 0 0 0 -1 0 0 0 0 0
// 	// 2 1

// 	// 6 5
// 	// 0 0 0 -1 0 0 0 0 0 0 0 0 -1 4 0 0 0 0 0 0 0 0 0 0 -1 0 0 0 0 0
// 	// 1 4

// 	// 找到头结点
// 	x0, y0 := 0, 0
// 	grid := make([][]int, 0)
// 	for i := 0; i < M; i++ {
// 		grid = append(grid, make([]int, N))
// 	}
// 	for i := 0; i < M; i++ {
// 		for j := 0; j < N; j++ {
// 			_, _ = fmt.Scan(&grid[i][j])
// 			if grid[i][j] > 0 {
// 				x0, y0 = i, j
// 			}
// 		}
// 	}

// 	for i := 0; i < M; i++ {
// 		for j := 0; j < N; j++ {
// 			fmt.Printf("%d  ", grid[i][j])
// 		}
// 		fmt.Println("")
// 	}

// 	x, y := 0, 0
// 	_, _ = fmt.Scan(&x, &y)

// 	queue := list.New()
// 	// 检索方向 上下左右
// 	ways := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 	// 头结点入队
// 	queue.PushBack([]int{x0, y0})
// 	// 记录已经搜索过的 cell
// 	visited := make([][]bool, 0)
// 	for i := 0; i < M; i++ {
// 		visited = append(visited, make([]bool, N))
// 	}
// 	// 直接 bfs 搜索
// 	found := false
// 	for queue.Len() > 0 {
// 		// 当前层需要搜索的节点个数
// 		n := queue.Len()
// 		// 处理 bfs 每一层的节点
// 		for i := 0; i < n; i++ {
// 			cell := queue.Front()
// 			queue.Remove(cell)
// 			pos := cell.Value.([]int)
// 			// 标记当前节点已经访问
// 			visited[pos[0]][pos[1]] = true
// 			// 检查当前节点的四周能否访问
// 			for j := 0; j < len(ways); j++ {
// 				xx, yy := pos[0]+ways[j][0], pos[1]+ways[j][1]
// 				// 检查是否能够达到新的节点
// 				if check(grid, xx, yy) && !visited[xx][yy] {
// 					// 能够达到新的节点，计算新节点的信号值，并且标记已经访问过
// 					grid[xx][yy] = grid[pos[0]][pos[1]] - 1
// 					if grid[xx][yy] == 0 {
// 						// 信号强度等于 0 就没必要搜索了
// 						break
// 					}
// 					// 已经搜索到目标节点
// 					if xx == x && yy == y {
// 						found = true
// 						break
// 					}
// 					visited[xx][yy] = true
// 					queue.PushBack([]int{xx, yy})
// 				}
// 			}
// 		}
// 		if found {
// 			break
// 		}
// 	}
// 	fmt.Print(grid[x][y])
// }

// func check(grid [][]int, x, y int) bool {
// 	m := len(grid)
// 	n := len(grid[0])
// 	return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != -1
// }
