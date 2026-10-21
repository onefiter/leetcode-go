package searchbst

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) *TreeNode{
	"searchBST":          searchBST,
	"searchBSTRecursive": searchBSTRecursive,
}

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

// serialize 层序序列化回 LeetCode 数组（去掉末尾 null）。
func serialize(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
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

func TestSearchBST(t *testing.T) {
	example := []int{4, 2, 7, 1, 3}
	tests := []struct {
		name string
		vals []int
		val  int
		want []int // 返回子树的层序数组，空表示 nil
	}{
		{"示例1", example, 2, []int{2, 1, 3}},
		{"示例2-不存在", example, 5, []int{}},
		{"找到根", example, 4, []int{4, 2, 7, 1, 3}},
		{"找到叶子", example, 3, []int{3}},
		{"空树", []int{}, 1, []int{}},
		{"单节点-命中", []int{1}, 1, []int{1}},
		{"左斜链", []int{3, 2, null, 1}, 2, []int{2, 1}},
		{"右斜链", []int{1, null, 2, null, 3}, 2, []int{2, null, 3}},
		{"完全二叉树", []int{4, 2, 6, 1, 3, 5, 7}, 6, []int{6, 5, 7}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := serialize(fn(buildTree(tt.vals), tt.val)); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.val, got, tt.want)
				}
			})
		}
	}
}
