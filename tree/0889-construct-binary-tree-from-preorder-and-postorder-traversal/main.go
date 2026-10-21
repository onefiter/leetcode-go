package buildtreeprepost

import "leetcode-go/template"

// 889. 根据前序和后序遍历构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
// 给定无重复值的前序 preorder 和后序 postorder，返回任意一棵满足条件的二叉树。
//
// 核心思路：前序是「根 | 左 | 右」，后序是「左 | 右 | 根」。
// 前序第二个元素 preorder[1] 是左子树的根（假设左子树非空），
// 它在后序中的位置决定了左子树的大小：leftSize = pos(preorder[1]) - postL + 1。
// 只有一个孩子时无法区分左右，所以答案不唯一，这里统一当作左孩子。

type TreeNode = template.TreeNode

// constructFromPrePost 解法一：递归 + 哈希表定位 ⭐ 最优方案
// 哈希表存后序值 → 下标；build(preL, preR, postL) 用前序区间 [preL, preR]、
// 后序起点 postL 建树；用 preorder[preL+1] 在后序中的位置算出左子树大小，再分别递归。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	pos := make(map[int]int, len(postorder))
	for i, v := range postorder {
		pos[v] = i
	}
	var build func(preL, preR, postL int) *TreeNode
	build = func(preL, preR, postL int) *TreeNode {
		if preL > preR {
			return nil
		}
		root := &TreeNode{Val: preorder[preL]}
		if preL == preR {
			return root
		}
		leftSize := pos[preorder[preL+1]] - postL + 1
		root.Left = build(preL+1, preL+leftSize, postL)
		root.Right = build(preL+leftSize+1, preR, postL+leftSize)
		return root
	}
	return build(0, len(preorder)-1, 0)
}

// constructFromPrePostStack 解法二：迭代 + 栈
// 按前序依次创建节点，栈里是当前「子树还没构造完」的祖先链。
// 栈顶值等于 postorder[j] 说明该子树已经完整（后序里出现了它），弹出并 j++。
// 新节点挂到栈顶：左孩子空就挂左边，否则挂右边。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func constructFromPrePostStack(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	j := 0
	for i := 1; i < len(preorder); i++ {
		// 弹出已经构造完整的子树。
		for stack[len(stack)-1].Val == postorder[j] {
			stack = stack[:len(stack)-1]
			j++
		}
		node := &TreeNode{Val: preorder[i]}
		top := stack[len(stack)-1]
		if top.Left == nil {
			top.Left = node
		} else {
			top.Right = node
		}
		stack = append(stack, node)
	}
	return root
}
