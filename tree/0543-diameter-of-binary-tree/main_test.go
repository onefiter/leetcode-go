package binarytreediameter

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) int{
	"diameterOfBinaryTree":            diameterOfBinaryTree,
	"diameterOfBinaryTreeIterative":   diameterOfBinaryTreeIterative,
	"diameterOfBinaryTreeEnumeration": diameterOfBinaryTreeEnumeration,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		cur := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			cur.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			cur.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 3},
		{"示例2", []int{1, 2}, 1},
		{"空树", []int{}, 0},
		{"单节点", []int{1}, 0},
		{"左斜链", []int{1, 2, null, 3, null, 4}, 3},
		{"右斜链", []int{1, null, 2, null, 3}, 2},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 4},
		{"不经过根", []int{1, 2, null, 3, 4, 5, null, null, 6, 7, null, null, 8}, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
