package balancedtree

import (
	"math"
	"testing"
)

const null = math.MinInt64

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

var solutions = map[string]func(*TreeNode) bool{
	"isBalanced":            isBalanced,
	"isBalancedEnumeration": isBalancedEnumeration,
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"示例1", []int{3, 9, 20, null, null, 15, 7}, true},
		{"示例2", []int{1, 2, 2, 3, 3, null, null, 4, 4}, false},
		{"空树", []int{}, true},
		{"单节点", []int{1}, true},
		{"两个节点", []int{1, 2}, true},
		{"左斜链", []int{1, 2, null, 3}, false},
		{"右斜链", []int{1, null, 2, null, 3}, false},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"根平衡但子树不平衡", []int{1, 2, 2, 3, null, null, 3, 4, null, null, 4}, false},
		{"深度差恰好为1", []int{1, 2, 3, 4, 5, null, 6, 7}, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
