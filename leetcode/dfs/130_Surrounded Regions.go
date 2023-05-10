package dfs

// func Solve(board [][]byte) {
// 	// 所有不被包围的 O 都直接或者间接地与边界相连
// 	// 从所有与边界直接相连的 O 开始搜索，将所有与之直接或间接相连的 O 标记
// 	// 搜索完成后，将所有没有被标记的 O 修改为 X

// 	m := len(board)
// 	n := len(board[0])
// 	// 搜索边界上的每个 O
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
// 				dfs(board, i, j)
// 			}
// 		}
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			// 将所有没有被标记的位置修改为 X
// 			if board[i][j] != 'A' {
// 				board[i][j] = 'X'
// 			} else {
// 				board[i][j] = 'O'
// 			}
// 		}
// 	}
// }

// func dfs(board [][]byte, i, j int) {
// 	// 到达边界，到达 X，已经访问过
// 	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == 'X' || board[i][j] == 'A' {
// 		return
// 	}
// 	// 将当前位置标记为可达边界
// 	board[i][j] = 'A'
// 	// 搜索上下左右四个方向
// 	dfs(board, i-1, j)
// 	dfs(board, i+1, j)
// 	dfs(board, i, j-1)
// 	dfs(board, i, j+1)
// }
