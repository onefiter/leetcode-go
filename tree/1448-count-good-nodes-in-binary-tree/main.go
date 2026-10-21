package goodnodes

import "leetcode-go/template"

// 1448. 统计二叉树中好节点的数目
// https://leetcode.cn/problems/count-good-nodes-in-binary-tree/
// 从根到节点 X 的路径上，如果没有任何节点的值大于 X，X 就是「好节点」。求好节点个数。
//
// 核心思路：往下走时携带「根到当前路径上的最大值」mx，
// 当前节点值 >= mx 就是好节点，然后用 max(mx, 当前值) 继续往下传。

type TreeNode = template.TreeNode

// goodNodes 解法一：递归 DFS 携带路径最大值 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(h)。
func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(n *TreeNode, mx int) int {
		if n == nil {
			return 0
		}
		cnt := 0
		if n.Val >= mx {
			cnt = 1
			mx = n.Val
		}
		return cnt + dfs(n.Left, mx) + dfs(n.Right, mx)
	}
	if root == nil {
		return 0
	}
	return dfs(root, root.Val)
}

// goodNodesBFS 解法二：迭代 BFS，队列中同时存放节点和路径最大值
// 时间复杂度 O(n)，空间复杂度 O(w)，w 为最大宽度。
func goodNodesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type item struct {
		node *TreeNode
		mx   int
	}
	cnt := 0
	queue := []item{{root, root.Val}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		mx := cur.mx
		if cur.node.Val >= mx {
			cnt++
			mx = cur.node.Val
		}
		if cur.node.Left != nil {
			queue = append(queue, item{cur.node.Left, mx})
		}
		if cur.node.Right != nil {
			queue = append(queue, item{cur.node.Right, mx})
		}
	}
	return cnt
}

// goodNodesEnumeration 解法三：对每个节点记录完整路径再逐个检查（对照用）
// 回溯维护根到当前节点的路径，到每个节点时扫描整条路径判断是否有更大值。
// 时间复杂度 O(n*h)，空间复杂度 O(h)。
func goodNodesEnumeration(root *TreeNode) int {
	cnt := 0
	path := []int{}
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		good := true
		for _, v := range path {
			if v > n.Val {
				good = false
				break
			}
		}
		if good {
			cnt++
		}
		path = append(path, n.Val)
		dfs(n.Left)
		dfs(n.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return cnt
}
