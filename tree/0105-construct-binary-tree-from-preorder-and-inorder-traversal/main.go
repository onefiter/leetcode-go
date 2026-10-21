package buildtreepreinorder

import "leetcode-go/template"

// 105. 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 给定无重复元素的前序 preorder 和中序 inorder，构造并返回这棵二叉树。
//
// 核心思路：前序的第一个元素就是根；在中序里找到根的位置 k，
// 中序 k 左边是左子树、右边是右子树，左子树大小 = k - 中序起点，
// 据此把前序也切成「根 | 左子树 | 右子树」，递归构造。

type TreeNode = template.TreeNode

// buildTree 解法一：递归 + 哈希表定位根 ⭐ 最优方案
// 先把中序值 → 下标存进哈希表，O(1) 找到根在中序中的位置，避免每层线性查找。
// build(preL, inL, inR) 表示用前序从 preL 开始、中序 [inL, inR] 的区间建树。
// 时间复杂度 O(n)，空间复杂度 O(n)（哈希表 + 递归栈）。
func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}
	var build func(preL, inL, inR int) *TreeNode
	build = func(preL, inL, inR int) *TreeNode {
		if inL > inR {
			return nil
		}
		rootVal := preorder[preL]
		k := pos[rootVal]
		leftSize := k - inL
		return &TreeNode{
			Val:   rootVal,
			Left:  build(preL+1, inL, k-1),
			Right: build(preL+1+leftSize, k+1, inR),
		}
	}
	return build(0, 0, len(inorder)-1)
}

// buildTreeStack 解法二：迭代 + 栈
// 按前序依次处理节点，栈里存「还没确定右孩子」的祖先链。
// 用指针 in 指向中序当前位置：
//   - 若栈顶值 != inorder[in]，说明还在往左走，新节点是栈顶的左孩子；
//   - 否则不断弹栈并右移 in，直到不相等，最后弹出的节点就是新节点的父亲，新节点是它的右孩子。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func buildTreeStack(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	in := 0
	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{Val: preorder[i]}
		top := stack[len(stack)-1]
		if top.Val != inorder[in] {
			top.Left = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[in] {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				in++
			}
			top.Right = node
		}
		stack = append(stack, node)
	}
	return root
}
