package insertbst

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) *TreeNode{
	"insertIntoBST":          insertIntoBST,
	"insertIntoBSTRecursive": insertIntoBSTRecursive,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
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

// isBST 用 int64 上下界判断严格 BST。
func isBST(root *TreeNode, lo, hi int64) bool {
	if root == nil {
		return true
	}
	v := int64(root.Val)
	if v <= lo || v >= hi {
		return false
	}
	return isBST(root.Left, lo, v) && isBST(root.Right, v, hi)
}

func TestInsertIntoBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		val  int
	}{
		{"示例1", []int{4, 2, 7, 1, 3}, 5},
		{"示例2", []int{40, 20, 60, 10, 30, 50, 70}, 25},
		{"示例3", []int{4, 2, 7, 1, 3, null, null, null, null, null, null}, 5},
		{"空树", []int{}, 5},
		{"单节点-插左", []int{5}, 3},
		{"单节点-插右", []int{5}, 8},
		{"左斜链-插最小", []int{3, 2, null, 1}, 0},
		{"右斜链-插中间", []int{1, null, 3, null, 5}, 4},
		{"完全二叉树-插最大", []int{4, 2, 6, 1, 3, 5, 7}, 8},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals) // 原地修改，每次重新构造
				want := append(inorderVals(root), tt.val)
				slices.Sort(want)
				got := fn(root, tt.val)
				if !isBST(got, math.MinInt64, math.MaxInt64) {
					t.Errorf("%s: 结果不是 BST", name)
				}
				if seq := inorderVals(got); !slices.Equal(seq, want) {
					t.Errorf("%s(%v, %d): 中序 = %v, want %v", name, tt.vals, tt.val, seq, want)
				}
			})
		}
	}
}
