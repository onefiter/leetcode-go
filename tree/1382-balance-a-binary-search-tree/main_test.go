package balancebst

import (
	"slices"
	"testing"
)

var solutions = map[string]func(*TreeNode) *TreeNode{
	"balanceBST":            balanceBST,
	"balanceBSTEnumeration": balanceBSTEnumeration,
}

// buildBST 按顺序把 vals 插入一棵空 BST（不做平衡），用于构造输入。
func buildBST(vals []int) *TreeNode {
	var root *TreeNode
	for _, v := range vals {
		node := &TreeNode{Val: v}
		if root == nil {
			root = node
			continue
		}
		cur := root
		for {
			if v < cur.Val {
				if cur.Left == nil {
					cur.Left = node
					break
				}
				cur = cur.Left
			} else {
				if cur.Right == nil {
					cur.Right = node
					break
				}
				cur = cur.Right
			}
		}
	}
	return root
}

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

// isBST 用上下界递归判断 BST 性质（允许重复值，左 <= 根 <= 右）。
func isBST(root *TreeNode, lo, hi *int) bool {
	if root == nil {
		return true
	}
	if (lo != nil && root.Val < *lo) || (hi != nil && root.Val > *hi) {
		return false
	}
	return isBST(root.Left, lo, &root.Val) && isBST(root.Right, &root.Val, hi)
}

func TestBalanceBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 插入顺序
	}{
		{"示例1-右斜链", []int{1, 2, 3, 4}},
		{"示例2-已平衡", []int{2, 1, 3}},
		{"单个节点", []int{1}},
		{"空树", []int{}},
		{"左斜链", []int{7, 6, 5, 4, 3, 2, 1}},
		{"之字形", []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}},
		{"长链", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}},
		{"含重复值", []int{2, 2, 2, 3, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法一会复用并改写原节点，每次重新构造输入。
				root := buildBST(tt.vals)
				want := inorderVals(root)
				got := fn(root)
				if seq := inorderVals(got); !slices.Equal(seq, want) {
					t.Errorf("%s: 中序序列 = %v, want %v", name, seq, want)
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
