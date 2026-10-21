package clonegraph

// 133. 克隆图
// https://leetcode.cn/problems/clone-graph/
// 给你无向连通图中一个节点的引用，返回该图的深拷贝（克隆）。
// 每个节点包含值 Val 和邻居列表 Neighbors。
//
// 核心思路：用哈希表记录「原节点 -> 克隆节点」，既能避免重复克隆，
// 也能防止在环上无限递归。遍历方式 DFS / BFS 均可。

type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph 解法一：DFS + 哈希表 ⭐ 最优方案
// 时间复杂度 O(V+E)，空间复杂度 O(V)。
func cloneGraph(node *Node) *Node {
	cloned := map[*Node]*Node{}
	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if c, ok := cloned[n]; ok {
			return c
		}
		// 先登记再递归邻居，遇到环时直接返回已登记的克隆
		c := &Node{Val: n.Val, Neighbors: make([]*Node, 0, len(n.Neighbors))}
		cloned[n] = c
		for _, nb := range n.Neighbors {
			c.Neighbors = append(c.Neighbors, dfs(nb))
		}
		return c
	}
	return dfs(node)
}

// cloneGraphBFS 解法二：BFS + 哈希表（对照用）
// 入队时创建克隆节点，出队时补齐邻居关系。
// 时间复杂度 O(V+E)，空间复杂度 O(V)。
func cloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloned := map[*Node]*Node{node: {Val: node.Val}}
	queue := []*Node{node}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, nb := range cur.Neighbors {
			if _, ok := cloned[nb]; !ok {
				cloned[nb] = &Node{Val: nb.Val}
				queue = append(queue, nb)
			}
			cloned[cur].Neighbors = append(cloned[cur].Neighbors, cloned[nb])
		}
	}
	return cloned[node]
}
