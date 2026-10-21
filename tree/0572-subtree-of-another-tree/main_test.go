package subtree

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

var solutions = map[string]func(*TreeNode, *TreeNode) bool{
	"isSubtree":            isSubtree,
	"isSubtreeEnumeration": isSubtreeEnumeration,
}

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name          string
		root, subRoot []int
		want          bool
	}{
		{"示例1", []int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{"示例2-子树下面多了节点", []int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{4, 1, 2}, false},
		{"整棵树相同", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"子树是叶子", []int{1, 2, 3}, []int{3}, true},
		{"值不存在", []int{1, 2, 3}, []int{4}, false},
		{"数字边界-12不应匹配2", []int{12}, []int{2}, false},
		{"数字边界-2不应匹配12", []int{2}, []int{12}, false},
		{"结构不同", []int{1, 2}, []int{1, null, 2}, false},
		{"深层匹配", []int{1, null, 2, null, 3, 4, 5}, []int{3, 4, 5}, true},
		{"只匹配到中间不算", []int{3, 4, 5, 1}, []int{3, 4, 5}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.root), buildTree(tt.subRoot)); got != tt.want {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.root, tt.subRoot, got, tt.want)
				}
			})
		}
	}
}
