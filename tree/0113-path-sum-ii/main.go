package pathsumii

import "leetcode-go/template"

// 113. 路径总和 II
// https://leetcode.cn/problems/path-sum-ii/
// 找出所有「根节点到叶子节点」路径中，节点值之和等于 targetSum 的路径。
//
// 核心思路：回溯。DFS 时把当前节点加入路径，到叶子时检查剩余和是否为 0，
// 满足就把路径「拷贝一份」放进结果；返回上一层前把当前节点从路径中弹出。

type TreeNode = template.TreeNode

// pathSum 解法一：回溯 DFS ⭐ 最优方案
// 共用一条 path 切片，进入节点 append，离开节点截断，只在命中叶子时拷贝。
// 时间复杂度 O(n^2)（最坏每个叶子都要拷贝长度 O(n) 的路径），空间复杂度 O(h)（不计结果）。
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, rest int) {
		if n == nil {
			return
		}
		rest -= n.Val
		path = append(path, n.Val)
		if n.Left == nil && n.Right == nil && rest == 0 {
			res = append(res, append([]int(nil), path...))
		}
		dfs(n.Left, rest)
		dfs(n.Right, rest)
		path = path[:len(path)-1] // 回溯
	}
	dfs(root, targetSum)
	return res
}

// pathSumBFS 解法二：BFS + 父指针哈希还原路径
// 层序遍历时记录每个节点的父节点和根到它的路径和；
// 遇到满足条件的叶子，沿父指针回到根，反转得到路径。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func pathSumBFS(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	parent := map[*TreeNode]*TreeNode{}
	sum := map[*TreeNode]int{root: root.Val}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.Left == nil && n.Right == nil {
			if sum[n] == targetSum {
				res = append(res, buildPath(n, parent))
			}
			continue
		}
		for _, child := range []*TreeNode{n.Left, n.Right} {
			if child != nil {
				parent[child] = n
				sum[child] = sum[n] + child.Val
				queue = append(queue, child)
			}
		}
	}
	return res
}

// buildPath 从叶子沿父指针回到根，返回根到叶子的值序列。
func buildPath(leaf *TreeNode, parent map[*TreeNode]*TreeNode) []int {
	path := []int{}
	for n := leaf; n != nil; n = parent[n] {
		path = append(path, n.Val)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
