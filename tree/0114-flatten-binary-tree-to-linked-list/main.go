package flattentree

import "leetcode-go/template"

// 114. 二叉树展开为链表
// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 把二叉树原地展开成一条「单链表」：顺序与前序遍历相同，
// 用 Right 指针指向下一个节点，Left 指针全部置为 nil。
//
// 核心思路：对当前节点 cur，如果有左子树，那么左子树前序遍历的「最后一个节点」
// （左子树最右下的节点）就是右子树的前驱。把右子树接到这个前驱的 Right 上，
// 再把左子树整体挪到右边，然后 cur 沿 Right 继续往下走。

type TreeNode = template.TreeNode

// flatten 解法一：原地寻找前驱（类 Morris）⭐ 最优方案
// 对每个有左子树的节点，找到左子树最右节点 pre，令 pre.Right = cur.Right，
// cur.Right = cur.Left，cur.Left = nil，然后处理 cur.Right。
// 每个节点最多被「找前驱」的过程访问常数次。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func flatten(root *TreeNode) {
	for cur := root; cur != nil; cur = cur.Right {
		if cur.Left == nil {
			continue
		}
		pre := cur.Left
		for pre.Right != nil {
			pre = pre.Right
		}
		pre.Right = cur.Right
		cur.Right = cur.Left
		cur.Left = nil
	}
}

// flattenRecursive 解法二：反前序递归（右 -> 左 -> 根）
// 按「右、左、根」的顺序遍历，正好是前序的逆序；
// 用 prev 记录上一个处理的节点，当前节点的 Right 指向 prev 即可从尾到头串起链表。
// 时间复杂度 O(n)，空间复杂度 O(h)。
func flattenRecursive(root *TreeNode) {
	var prev *TreeNode
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Right)
		dfs(n.Left)
		n.Right = prev
		n.Left = nil
		prev = n
	}
	dfs(root)
}

// flattenEnumeration 解法三：前序收集节点后重连（对照用）
// 先迭代前序遍历把所有节点按顺序放进切片，再依次把相邻节点串起来。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func flattenEnumeration(root *TreeNode) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nodes = append(nodes, n)
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}
	}
	for i := 0; i < len(nodes); i++ {
		nodes[i].Left = nil
		if i+1 < len(nodes) {
			nodes[i].Right = nodes[i+1]
		} else {
			nodes[i].Right = nil
		}
	}
}
