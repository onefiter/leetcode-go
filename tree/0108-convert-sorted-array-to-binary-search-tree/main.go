package sortedarraytobst

import "leetcode-go/template"

// 108. 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
// 给一个升序整数数组 nums，把它转换成一棵「高度平衡」的二叉搜索树（答案不唯一）。
//
// 核心思路：BST 的中序遍历就是升序数组。每次取区间中点作为根，
// 左半边建左子树、右半边建右子树，左右规模最多差 1，自然平衡。

type TreeNode = template.TreeNode

// sortedArrayToBST 解法一：递归取中点 ⭐ 最优方案
// build(l, r) 用 nums[l..r] 建树：mid = (l+r)/2 作根，递归建左右子树。
// 时间复杂度 O(n)，空间复杂度 O(log n)（递归栈，树是平衡的）。
func sortedArrayToBST(nums []int) *TreeNode {
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		return &TreeNode{
			Val:   nums[mid],
			Left:  build(l, mid-1),
			Right: build(mid+1, r),
		}
	}
	return build(0, len(nums)-1)
}

// sortedArrayToBSTQueue 解法二：迭代 BFS
// 队列里存「待填值的节点 + 它负责的区间 [l, r]」：
// 出队后把 nums[mid] 填进节点，若左/右半区间非空就新建孩子并入队。
// 时间复杂度 O(n)，空间复杂度 O(n)（队列）。
func sortedArrayToBSTQueue(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	type item struct {
		node *TreeNode
		l, r int
	}
	root := &TreeNode{}
	queue := []item{{root, 0, len(nums) - 1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		mid := cur.l + (cur.r-cur.l)/2
		cur.node.Val = nums[mid]
		if cur.l <= mid-1 {
			cur.node.Left = &TreeNode{}
			queue = append(queue, item{cur.node.Left, cur.l, mid - 1})
		}
		if mid+1 <= cur.r {
			cur.node.Right = &TreeNode{}
			queue = append(queue, item{cur.node.Right, mid + 1, cur.r})
		}
	}
	return root
}
