package invertbinarytree

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

var solutions = map[string]func(*TreeNode) *TreeNode{
	"invertTree":    invertTree,
	"invertTreeBFS": invertTreeBFS,
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want []int
	}{
		{"示例1", []int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"示例2", []int{2, 1, 3}, []int{2, 3, 1}},
		{"空树", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"左斜链变右斜链", []int{1, 2, null, 3}, []int{1, null, 2, null, 3}},
		{"右斜链变左斜链", []int{1, null, 2, null, 3}, []int{1, 2, null, 3}},
		{"只有一侧孩子", []int{1, 2, 3, null, 4}, []int{1, 3, 2, null, null, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := serialize(fn(buildTree(tt.root))); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.root, got, tt.want)
				}
			})
		}
	}
}

// TestInvertTwiceIsIdentity 翻转两次应还原为原树。
func TestInvertTwiceIsIdentity(t *testing.T) {
	vals := []int{5, 3, 8, 1, 4, 7, 9, null, 2, null, null, 6}
	for name, fn := range solutions {
		if got := serialize(fn(fn(buildTree(vals)))); !slices.Equal(got, vals) {
			t.Errorf("%s 翻转两次 = %v, want %v", name, got, vals)
		}
	}
}
