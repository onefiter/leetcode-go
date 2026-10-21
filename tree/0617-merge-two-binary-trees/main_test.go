package mergetrees

import (
	"math"
	"slices"
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

// serialize 层序序列化回 LeetCode 数组（去掉末尾 null），用于比较结构和值。
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

var solutions = map[string]func(*TreeNode, *TreeNode) *TreeNode{
	"mergeTrees":    mergeTrees,
	"mergeTreesBFS": mergeTreesBFS,
}

func TestMergeTrees(t *testing.T) {
	tests := []struct {
		name         string
		root1, root2 []int
		want         []int
	}{
		{"示例1", []int{1, 3, 2, 5}, []int{2, 1, 3, null, 4, null, 7}, []int{3, 4, 5, 5, 4, null, 7}},
		{"示例2", []int{1}, []int{1, 2}, []int{2, 2}},
		{"两棵空树", []int{}, []int{}, []int{}},
		{"左空", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"右空", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"结构完全相同", []int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},
		{"一左斜一右斜", []int{1, 2, null, 3}, []int{1, null, 2, null, 3}, []int{2, 2, 2, 3, null, null, 3}},
		{"含负数", []int{-1, 2}, []int{1, -2, 3}, []int{0, 0, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := serialize(fn(buildTree(tt.root1), buildTree(tt.root2))); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.root1, tt.root2, got, tt.want)
				}
			})
		}
	}
}
