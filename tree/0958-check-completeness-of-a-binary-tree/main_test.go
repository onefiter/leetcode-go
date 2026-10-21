package completetree

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
	"isCompleteTree":      isCompleteTree,
	"isCompleteTreeIndex": isCompleteTreeIndex,
}

func TestIsCompleteTree(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want bool
	}{
		{"示例1", []int{1, 2, 3, 4, 5, 6}, true},
		{"示例2-中间缺口", []int{1, 2, 3, 4, 5, null, 7}, false},
		{"空树", []int{}, true},
		{"单节点", []int{1}, true},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"只有左孩子", []int{1, 2}, true},
		{"只有右孩子", []int{1, null, 2}, false},
		{"左斜链深度3", []int{1, 2, null, 3}, false},
		{"最后一层靠左", []int{1, 2, 3, 4}, true},
		{"最后一层右侧缺左", []int{1, 2, 3, null, 5}, false},
		{"上一层未满", []int{1, 2, 3, 4, 5, null, null, 8}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.root)); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.root, got, tt.want)
				}
			})
		}
	}
}
