package zigzaglevelorder

import (
	"math"
	"reflect"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) [][]int{
	"zigzagLevelOrderBFS": zigzagLevelOrderBFS,
	"zigzagLevelOrderDFS": zigzagLevelOrderDFS,
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

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want [][]int
	}{
		{"示例1", []int{3, 9, 20, null, null, 15, 7}, [][]int{{3}, {20, 9}, {15, 7}}},
		{"空树", []int{}, [][]int{}},
		{"单节点", []int{1}, [][]int{{1}}},
		{"左斜链", []int{3, 2, null, 1}, [][]int{{3}, {2}, {1}}},
		{"右斜链", []int{1, null, 2, null, 3}, [][]int{{1}, {2}, {3}}},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}},
		{"四层", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, [][]int{{1}, {3, 2}, {4, 5, 6, 7}, {11, 10, 9, 8}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.tree)); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.tree, got, tt.want)
				}
			})
		}
	}
}
