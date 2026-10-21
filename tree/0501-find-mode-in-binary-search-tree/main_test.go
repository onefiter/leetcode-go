package findbstmode

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) []int{
	"findMode":            findMode,
	"findModeRecursive":   findModeRecursive,
	"findModeEnumeration": findModeEnumeration,
}

// buildTree 按 LeetCode 层序数组构造树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	q := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := q[0]
		q = q[1:]
		if vals[i] != null {
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

// levelOrder 层序序列化（去掉末尾 null），用于确认 Morris 遍历后树结构已恢复。
func levelOrder(root *TreeNode) []int {
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, null)
			continue
		}
		res = append(res, node.Val)
		q = append(q, node.Left, node.Right)
	}
	for len(res) > 0 && res[len(res)-1] == null {
		res = res[:len(res)-1]
	}
	return res
}

func TestFindMode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{1, null, 2, 2}, []int{2}},
		{"示例2-单节点", []int{0}, []int{0}},
		{"全部不同-都是众数", []int{4, 2, 6, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"多个众数", []int{2, 1, 3, 1, null, 3}, []int{1, 3}},
		{"左斜链全相同", []int{5, 5, null, 5, null, 5}, []int{5}},
		{"右斜链", []int{1, null, 2, null, 2, null, 3}, []int{2}},
		{"含负数", []int{0, -1, 1, -1, null, null, 1}, []int{-1, 1}},
		{"根与子树相同值", []int{6, 2, 8, 0, 4, 7, 9, null, null, 2, 6}, []int{2, 6}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals)
				got := fn(root)
				slices.Sort(got)
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
				if after := levelOrder(root); !slices.Equal(after, levelOrder(buildTree(tt.vals))) {
					t.Errorf("%s modified tree: %v", name, after)
				}
			})
		}
	}
}
