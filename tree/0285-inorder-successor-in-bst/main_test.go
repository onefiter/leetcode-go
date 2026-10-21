package inordersuccessorbst

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(root, p *TreeNode) *TreeNode{
	"inorderSuccessor":            inorderSuccessor,
	"inorderSuccessorEnumeration": inorderSuccessorEnumeration,
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

func TestInorderSuccessor(t *testing.T) {
	example2 := []int{5, 3, 6, 2, 4, null, null, 1}
	tests := []struct {
		name    string
		vals    []int
		p       int
		wantVal int // null 表示期望 nil
	}{
		{"示例1", []int{2, 1, 3}, 1, 2},
		{"示例2-无后继", example2, 6, null},
		{"后继是祖先", example2, 4, 5},
		{"后继在右子树", example2, 3, 4},
		{"叶子后继是父亲", example2, 1, 2},
		{"单节点", []int{1}, 1, null},
		{"左斜链", []int{3, 2, null, 1}, 1, 2},
		{"右斜链", []int{1, null, 2, null, 3}, 2, 3},
		{"完全二叉树-根", []int{4, 2, 6, 1, 3, 5, 7}, 4, 5},
		{"完全二叉树-左子树最右", []int{4, 2, 6, 1, 3, 5, 7}, 3, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals)
				p := findNode(root, tt.p)
				var want *TreeNode
				if tt.wantVal != null {
					want = findNode(root, tt.wantVal)
				}
				if got := fn(root, p); got != want {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.p, got, want)
				}
			})
		}
	}
}
