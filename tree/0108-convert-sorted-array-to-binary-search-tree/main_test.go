package sortedarraytobst

import (
	"slices"
	"testing"
)

// inorderVals 返回中序遍历序列。
func inorderVals(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		res = append(res, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

// height 返回树高；若某处不平衡返回 -1。
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := height(root.Left), height(root.Right)
	if l == -1 || r == -1 || l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}

// isBST 用上下界递归判断严格 BST 性质。
func isBST(root *TreeNode, lo, hi *int) bool {
	if root == nil {
		return true
	}
	if (lo != nil && root.Val <= *lo) || (hi != nil && root.Val >= *hi) {
		return false
	}
	return isBST(root.Left, lo, &root.Val) && isBST(root.Right, &root.Val, hi)
}

var solutions = map[string]func([]int) *TreeNode{
	"sortedArrayToBST":      sortedArrayToBST,
	"sortedArrayToBSTQueue": sortedArrayToBSTQueue,
}

// 答案不唯一：校验结果是平衡 BST，且中序遍历等于输入。
func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"示例1", []int{-10, -3, 0, 5, 9}},
		{"示例2", []int{1, 3}},
		{"空数组", []int{}},
		{"单个元素", []int{7}},
		{"三个元素", []int{1, 2, 3}},
		{"完全二叉树规模", []int{1, 2, 3, 4, 5, 6, 7}},
		{"偶数个", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"含负数", []int{-100, -50, -1, 0, 1, 50, 100, 200}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.nums)
				if seq := inorderVals(got); !slices.Equal(seq, tt.nums) {
					t.Errorf("%s: 中序序列 = %v, want %v", name, seq, tt.nums)
				}
				if !isBST(got, nil, nil) {
					t.Errorf("%s: 结果不是 BST", name)
				}
				if height(got) == -1 {
					t.Errorf("%s: 结果不平衡", name)
				}
			})
		}
	}
}
