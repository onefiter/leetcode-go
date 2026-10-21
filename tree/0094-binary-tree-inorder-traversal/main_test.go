package inordertraversal

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) []int{
	"inorderTraversalMorris":    inorderTraversalMorris,
	"inorderTraversalRecursive": inorderTraversalRecursive,
	"inorderTraversalIterative": inorderTraversalIterative,
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

// serialize 把树按层序序列化成 LeetCode 格式（去掉末尾 null），用于检查树结构是否被恢复。
func serialize(root *TreeNode) []int {
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

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want []int
	}{
		{"示例1", []int{1, null, 2, 3}, []int{1, 3, 2}},
		{"空树", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"左斜链", []int{3, 2, null, 1}, []int{1, 2, 3}},
		{"右斜链", []int{1, null, 2, null, 3}, []int{1, 2, 3}},
		{"完全二叉树", []int{4, 2, 6, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"不完全", []int{1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9}, []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.tree)
				if got := fn(root); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.tree, got, tt.want)
				}
				// Morris 会临时修改树，结束后必须恢复原结构。
				if got := serialize(root); !slices.Equal(got, tt.tree) {
					t.Errorf("%s 修改了树结构: %v, want %v", name, got, tt.tree)
				}
			})
		}
	}
}
