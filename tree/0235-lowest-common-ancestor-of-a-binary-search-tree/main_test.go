package lcabst

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(root, p, q *TreeNode) *TreeNode{
	"lowestCommonAncestor":          lowestCommonAncestor,
	"lowestCommonAncestorRecursive": lowestCommonAncestorRecursive,
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

// findNode 按值查找节点指针（值唯一）。
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if n := findNode(root.Left, val); n != nil {
		return n
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	example := []int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}
	tests := []struct {
		name    string
		vals    []int
		p, q    int
		wantVal int
	}{
		{"示例1", example, 2, 8, 6},
		{"示例2-自身是祖先", example, 2, 4, 2},
		{"深层分叉", example, 3, 5, 4},
		{"跨根", example, 3, 7, 6},
		{"两节点", []int{2, 1}, 2, 1, 2},
		{"单节点-p等于q", []int{1}, 1, 1, 1},
		{"左斜链", []int{3, 2, null, 1}, 1, 2, 2},
		{"右斜链", []int{1, null, 2, null, 3}, 1, 3, 1},
		{"完全二叉树-叶子", []int{4, 2, 6, 1, 3, 5, 7}, 5, 7, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals)
				p, q := findNode(root, tt.p), findNode(root, tt.q)
				want := findNode(root, tt.wantVal)
				if got := fn(root, p, q); got != want {
					t.Errorf("%s(%v, %d, %d) = %v, want node %d", name, tt.vals, tt.p, tt.q, got, tt.wantVal)
				}
			})
		}
	}
}
