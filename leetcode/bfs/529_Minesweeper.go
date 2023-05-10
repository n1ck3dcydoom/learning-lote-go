package bfs

// import "container/list"

// func UpdateBoard(board [][]byte, click []int) [][]byte {
// 	if board[click[0]][click[1]] == 'M' {
// 		board[click[0]][click[1]] = 'X'
// 		return board
// 	}
// 	// 保存方向向量    上       下       左       右     左上      左下      右上     右下
// 	ways := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
// 	// 保存已经添加过队列的元素，避免重复搜索
// 	visited := make([][]int, 0)
// 	for i := 0; i < len(board); i++ {
// 		visited = append(visited, make([]int, len(board[0])))
// 	}
// 	visited[click[0]][click[1]] = 1
// 	// 头结点入队
// 	queue := list.New()
// 	queue.PushBack(click)
// 	// bfs 从 click 位置开始直接搜索
// 	for queue.Len() > 0 {
// 		// 当前层需要遍历的节点个数
// 		n := queue.Len()
// 		for i := 0; i < n; i++ {
// 			node := queue.Front().Value.([]int)
// 			// c 保存当前 cell 周围的地雷个数
// 			c, x, y := 0, node[0], node[1]
// 			queue.Remove(queue.Front())
// 			// 搜索当前 cell 周围所有方向
// 			for j := 0; j < len(ways); j++ {
// 				xx, yy := x+ways[j][0], y+ways[j][1]
// 				// 计算周围地雷的个数
// 				if arrive(board, xx, yy) {
// 					nextCell := board[xx][yy]
// 					if nextCell == 'M' {
// 						c++
// 					}
// 				}
// 			}
// 			// 如果 cell 周围有地雷，则停止揭露当前 cell 周围的格子
// 			if c > 0 {
// 				board[x][y] = byte(c + '0')
// 			} else {
// 				// 当前 cell 周围没有地雷，标记为 B 并且继续揭露周围的 E
// 				board[x][y] = 'B'
// 				for j := 0; j < len(ways); j++ {
// 					xx, yy := x+ways[j][0], y+ways[j][1]
// 					if arrive(board, xx, yy) && board[xx][yy] == 'E' && visited[xx][yy] == 0 {
// 						queue.PushBack([]int{xx, yy})
// 						// 标记当前 cell 已经被搜索
// 						visited[xx][yy] = 1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return board
// }

// func arrive(board [][]byte, x, y int) bool {
// 	m := len(board)
// 	n := len(board[0])
// 	if x < 0 || x >= m || y < 0 || y >= n {
// 		return false
// 	}
// 	return true
// }
