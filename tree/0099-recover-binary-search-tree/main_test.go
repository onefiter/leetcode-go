package recoverbst

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode){
	"recoverTree":          recoverTree,
	"recoverTreeRecursive": recoverTreeRecursive,
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

// inorderVals 返回中序遍历序列。
func inorderVals(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		res = append(res, n.Val)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int // 恢复后的层序数组
	}{
		{"示例1-相邻交换", []int{1, 3, null, null, 2}, []int{3, 1, null, null, 2}},
		{"示例2-不相邻交换", []int{3, 1, 4, null, null, 2}, []int{2, 1, 4, null, null, 3}},
		{"两节点", []int{1, 2}, []int{2, 1}},
		{"根与右孩子交换", []int{3, 1, 2}, []int{2, 1, 3}},
		{"左斜链", []int{1, 2, null, 3}, []int{3, 2, null, 1}},
		{"右斜链", []int{1, null, 3, null, 2}, []int{1, null, 2, null, 3}},
		{"完全二叉树-叶子交换", []int{4, 2, 6, 7, 3, 5, 1}, []int{4, 2, 6, 1, 3, 5, 7}},
		{"边界值", []int{math.MinInt32, math.MaxInt32}, []int{math.MaxInt32, math.MinInt32}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals) // 原地修改，每次重新构造
				fn(root)
				if got := serialize(root); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) => %v, want %v", name, tt.vals, got, tt.want)
				}
				if seq := inorderVals(root); !slices.IsSorted(seq) {
					t.Errorf("%s: 中序 %v 不是有序的", name, seq)
				}
			})
		}
	}
}
