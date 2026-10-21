package buildtreeinpost

import "leetcode-go/template"

// 106. 从中序与后序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 给定无重复元素的中序 inorder 和后序 postorder，构造并返回这棵二叉树。
//
// 核心思路：后序的最后一个元素是根；在中序里找到根的位置 k，
// 中序 k 左边是左子树、右边是右子树。后序是「左 | 右 | 根」，
// 从后往前取根时先遇到右子树的根，所以先建右子树、再建左子树。

type TreeNode = template.TreeNode

// buildTree 解法一：递归 + 哈希表定位根 ⭐ 最优方案
// 哈希表存中序值 → 下标；post 指针从后往前依次取根，
// build(inL, inR) 用中序区间 [inL, inR] 建树，先右后左。
// 时间复杂度 O(n)，空间复杂度 O(n)（哈希表 + 递归栈）。
func buildTree(inorder []int, postorder []int) *TreeNode {
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}
	post := len(postorder) - 1
	var build func(inL, inR int) *TreeNode
	build = func(inL, inR int) *TreeNode {
		if inL > inR {
			return nil
		}
		root := &TreeNode{Val: postorder[post]}
		post--
		k := pos[root.Val]
		root.Right = build(k+1, inR) // 先右后左
		root.Left = build(inL, k-1)
		return root
	}
	return build(0, len(inorder)-1)
}

// buildTreeStack 解法二：迭代 + 栈
// 和 105 的迭代法镜像：后序倒着看是「根 右 左」，中序倒着看是「右 根 左」。
// 从后往前处理后序，in 指针从中序末尾往前走：
//   - 栈顶值 != inorder[in]：还在往右走，新节点是栈顶的右孩子；
//   - 否则不断弹栈并左移 in，最后弹出的节点是新节点的父亲，新节点是它的左孩子。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func buildTreeStack(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[n-1]}
	stack := []*TreeNode{root}
	in := n - 1
	for i := n - 2; i >= 0; i-- {
		node := &TreeNode{Val: postorder[i]}
		top := stack[len(stack)-1]
		if top.Val != inorder[in] {
			top.Right = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[in] {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				in--
			}
			top.Left = node
		}
		stack = append(stack, node)
	}
	return root
}
