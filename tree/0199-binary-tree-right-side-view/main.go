package rightsideview

import "leetcode-go/template"

// 199. 二叉树的右视图
// https://leetcode.cn/problems/binary-tree-right-side-view/
// 想象自己站在二叉树的右侧，从上到下返回能看到的节点值。
// 也就是每一层「最右边」的那个节点。
//
// 核心思路：
//   - BFS：按层遍历，取每层最后一个节点；
//   - DFS：按「根 -> 右 -> 左」遍历，每层第一次到达的节点就是最右的节点。
//     注意最右节点不一定在右子树里（右子树较浅时，左子树更深的部分也能被看到）。

type TreeNode = template.TreeNode

// rightSideViewBFS 解法一：BFS 取每层最后一个 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(w)（w 为最大层宽）。
func rightSideViewBFS(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		res = append(res, cur[len(cur)-1].Val)
		next := []*TreeNode{}
		for _, node := range cur {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
	}
	return res
}

// rightSideViewDFS 解法二：DFS 先右后左
// 深度 depth == len(res) 说明是第一次到达这一层，由于先走右子树，此节点就是该层最右的。
// 时间复杂度 O(n)，空间复杂度 O(h)（递归栈）。
func rightSideViewDFS(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}
