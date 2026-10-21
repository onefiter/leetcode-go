package buildtreepreinorder

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

// serialize 层序序列化为 LeetCode 格式切片（空节点为 null，去掉末尾 null）。
func serialize(root *TreeNode) []int {
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, null)
			continue
		}
		res = append(res, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(res) > 0 && res[len(res)-1] == null {
		res = res[:len(res)-1]
	}
	return res
}

var solutions = map[string]func([]int, []int) *TreeNode{
	"buildTree":      buildTree,
	"buildTreeStack": buildTreeStack,
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name              string
		preorder, inorder []int
		want              []int // 层序
	}{
		{"示例1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, null, null, 15, 7}},
		{"示例2-单节点", []int{-1}, []int{-1}, []int{-1}},
		{"空树", []int{}, []int{}, []int{}},
		{"左斜链", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}, []int{1, 2, null, 3, null, 4}},
		{"右斜链", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []int{1, null, 2, null, 3, null, 4}},
		{"完全二叉树", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"之字形", []int{1, 2, 3, 4}, []int{2, 4, 3, 1}, []int{1, 2, null, null, 3, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := serialize(fn(tt.preorder, tt.inorder)); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.preorder, tt.inorder, got, tt.want)
				}
			})
		}
	}
}
