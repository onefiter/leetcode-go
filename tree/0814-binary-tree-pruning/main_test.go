package pruning

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
	"pruneTree":            pruneTree,
	"pruneTreeEnumeration": pruneTreeEnumeration,
}

func TestPruneTree(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want []int
	}{
		{"示例1", []int{1, null, 0, 0, 1}, []int{1, null, 0, null, 1}},
		{"示例2", []int{1, 0, 1, 0, 0, 0, 1}, []int{1, null, 1, null, 1}},
		{"示例3", []int{1, 1, 0, 1, 1, 0, 1, 0}, []int{1, 1, 0, 1, 1, null, 1}},
		{"空树", []int{}, []int{}},
		{"单个0被删光", []int{0}, []int{}},
		{"全是0", []int{0, 0, 0, 0}, []int{}},
		{"单个1保留", []int{1}, []int{1}},
		{"根是0但子树有1", []int{0, null, 0, null, 1}, []int{0, null, 0, null, 1}},
		{"全是1不变", []int{1, 1, 1}, []int{1, 1, 1}},
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
