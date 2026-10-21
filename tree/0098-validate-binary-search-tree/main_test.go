package validatebst

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) bool{
	"isValidBST":        isValidBST,
	"isValidBSTInorder": isValidBSTInorder,
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

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"示例1", []int{2, 1, 3}, true},
		{"示例2", []int{5, 1, 4, null, null, 3, 6}, false},
		{"空树", []int{}, true},
		{"单节点", []int{1}, true},
		{"单节点-MinInt32", []int{math.MinInt32}, true},
		{"单节点-MaxInt32", []int{math.MaxInt32}, true},
		{"边界值合法", []int{0, math.MinInt32, math.MaxInt32}, true},
		{"MaxInt32在左侧", []int{math.MaxInt32, math.MaxInt32}, false},
		{"重复值", []int{2, 2, 2}, false},
		{"左斜链合法", []int{3, 2, null, 1}, true},
		{"右斜链合法", []int{1, null, 2, null, 3}, true},
		{"右斜链非法", []int{1, null, 3, null, 2, null, 4}, false},
		{"完全二叉树合法", []int{4, 2, 6, 1, 3, 5, 7}, true},
		{"跨层违规", []int{5, 4, 6, null, null, 3, 7}, false},
		{"左子树深处违规", []int{10, 5, 15, 3, 11}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
