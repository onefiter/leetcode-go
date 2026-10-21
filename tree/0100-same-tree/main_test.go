package sametree

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

var solutions = map[string]func(*TreeNode, *TreeNode) bool{
	"isSameTree":    isSameTree,
	"isSameTreeBFS": isSameTreeBFS,
}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p, q []int
		want bool
	}{
		{"示例1", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"示例2-结构不同", []int{1, 2}, []int{1, null, 2}, false},
		{"示例3-值不同", []int{1, 2, 1}, []int{1, 1, 2}, false},
		{"两棵空树", []int{}, []int{}, true},
		{"一空一非空", []int{}, []int{1}, false},
		{"单节点相同", []int{1}, []int{1}, true},
		{"单节点不同", []int{1}, []int{2}, false},
		{"左斜链相同", []int{1, 2, null, 3}, []int{1, 2, null, 3}, true},
		{"右斜链相同", []int{1, null, 2, null, 3}, []int{1, null, 2, null, 3}, true},
		{"完全二叉树相同", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"深层值不同", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 8}, false},
		{"深层多一个节点", []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.p), buildTree(tt.q)); got != tt.want {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.p, tt.q, got, tt.want)
				}
			})
		}
	}
}
