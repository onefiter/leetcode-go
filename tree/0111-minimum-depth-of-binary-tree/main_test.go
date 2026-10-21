package mindepthtree

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) int{
	"minDepthBFS": minDepthBFS,
	"minDepthDFS": minDepthDFS,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	q := []*TreeNode{root}
	for i := 1; len(q) > 0 && i < len(vals); {
		node := q[0]
		q = q[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			q = append(q, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			q = append(q, node.Right)
		}
		i++
	}
	return root
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{"示例1", []int{3, 9, 20, null, null, 15, 7}, 2},
		{"示例2-右斜链", []int{2, null, 3, null, 4, null, 5, null, 6}, 5},
		{"空树", []int{}, 0},
		{"单节点", []int{1}, 1},
		{"只有左孩子", []int{1, 2}, 2},
		{"左斜链", []int{1, 2, null, 3, null, 4}, 4},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"一侧单链另一侧叶子", []int{1, 2, 3, 4, null, null, null, 5}, 2},
		{"根只有一侧且深层有叶子", []int{1, null, 2, 3, 4, null, null, 5}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.tree)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.tree, got, tt.want)
				}
			})
		}
	}
}
