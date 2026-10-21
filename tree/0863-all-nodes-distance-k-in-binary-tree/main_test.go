package distancek

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

// find 按值找到节点指针（题目保证值唯一）。
func find(root *TreeNode, v int) *TreeNode {
	if root == nil || root.Val == v {
		return root
	}
	if n := find(root.Left, v); n != nil {
		return n
	}
	return find(root.Right, v)
}

var solutions = map[string]func(*TreeNode, *TreeNode, int) []int{
	"distanceK":            distanceK,
	"distanceKEnumeration": distanceKEnumeration,
}

func TestDistanceK(t *testing.T) {
	tests := []struct {
		name   string
		root   []int
		target int
		k      int
		want   []int
	}{
		{"示例1", []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}, 5, 2, []int{1, 4, 7}},
		{"示例2-单节点", []int{1}, 1, 3, []int{}},
		{"k=0 返回自己", []int{3, 5, 1}, 5, 0, []int{5}},
		{"往上再往下", []int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}, 6, 3, []int{4, 7, 1}},
		{"target 是根", []int{1, 2, 3, 4, 5, 6, 7}, 1, 2, []int{4, 5, 6, 7}},
		{"距离超出树高", []int{1, 2, 3}, 2, 5, []int{}},
		{"左斜链", []int{1, 2, null, 3, null, 4}, 4, 2, []int{2}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.root)
				got := fn(root, find(root, tt.target), tt.k)
				slices.Sort(got)
				want := slices.Clone(tt.want)
				slices.Sort(want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%v, target=%d, k=%d) = %v, want %v", name, tt.root, tt.target, tt.k, got, want)
				}
			})
		}
	}
}
