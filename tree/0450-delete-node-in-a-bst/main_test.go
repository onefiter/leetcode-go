package deletebstnode

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) *TreeNode{
	"deleteNode":          deleteNode,
	"deleteNodeIterative": deleteNodeIterative,
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

// isBST 校验严格 BST（值互不相同）。
func isBST(root *TreeNode, lo, hi int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lo || root.Val >= hi {
		return false
	}
	return isBST(root.Left, lo, root.Val) && isBST(root.Right, root.Val, hi)
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		key  int
	}{
		{"示例1-删有两孩子节点", []int{5, 3, 6, 2, 4, null, 7}, 3},
		{"示例2-key不存在", []int{5, 3, 6, 2, 4, null, 7}, 0},
		{"示例3-空树", []int{}, 0},
		{"单节点-删除", []int{1}, 1},
		{"删根-两孩子", []int{5, 3, 6, 2, 4, null, 7}, 5},
		{"删叶子", []int{5, 3, 6, 2, 4, null, 7}, 7},
		{"删只有右孩子", []int{5, 3, 6, 2, 4, null, 7}, 6},
		{"左斜链删中间", []int{4, 3, null, 2, null, 1}, 3},
		{"右斜链删根", []int{1, null, 2, null, 3, null, 4}, 1},
		{"完全二叉树删根", []int{4, 2, 6, 1, 3, 5, 7}, 4},
		{"后继有右孩子", []int{5, 2, 9, null, null, 7, 10, 6, 8}, 5},
		{"含负数", []int{0, -5, 10, -8, -2, null, 20}, -5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				want := slices.DeleteFunc(inorderVals(buildTree(tt.vals)), func(v int) bool { return v == tt.key })
				got := fn(buildTree(tt.vals), tt.key)
				if !isBST(got, math.MinInt, math.MaxInt) {
					t.Errorf("%s(%v, %d) result is not a BST", name, tt.vals, tt.key)
				}
				if in := inorderVals(got); !slices.Equal(in, want) {
					t.Errorf("%s(%v, %d) inorder = %v, want %v", name, tt.vals, tt.key, in, want)
				}
			})
		}
	}
}
