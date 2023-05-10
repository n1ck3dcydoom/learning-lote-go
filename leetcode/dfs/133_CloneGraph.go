package dfs

// import (
// 	"learning-lote-go/codetop/common"
// )

// /**
//  * Definition for a Node.
//  * type Node struct {
//  *     Val int
//  *     Neighbors []*Node
//  * }
//  */

// func CloneGraph(node *common.Node) *common.Node {
// 	return dfs(node, make(map[*common.Node]*common.Node))
// }

// func dfs(node *common.Node, visited map[*common.Node]*common.Node) *common.Node {
// 	if node == nil {
// 		return nil
// 	}
// 	// 检查是否已经访问过
// 	if _, ok := visited[node]; ok {
// 		// 返回已经访问过的克隆节点
// 		return visited[node]
// 	}
// 	// 当前节点还没有被访问过，克隆并加入 visited
// 	newNode := &common.Node{
// 		Val:       node.Val,
// 		Neighbors: make([]*common.Node, 0),
// 	}
// 	// 源节点 node 对应的克隆节点已经访问过
// 	visited[node] = newNode
// 	// 重建新节点的邻居节点
// 	for _, n := range node.Neighbors {
// 		newNode.Neighbors = append(newNode.Neighbors, dfs(n, visited))
// 	}
// 	// 返回克隆后的新节点
// 	return newNode
// }
