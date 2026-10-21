package buildtreeprepost

import (
	"slices"
	"testing"
)

// preorderVals 返回前序遍历序列。
func preorderVals(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return res
}

// postorderVals 返回后序遍历序列。
func postorderVals(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		dfs(n.Right)
		res = append(res, n.Val)
	}
	dfs(root)
	return res
}

var solutions = map[string]func([]int, []int) *TreeNode{
	"constructFromPrePost":      constructFromPrePost,
	"constructFromPrePostStack": constructFromPrePostStack,
}

// 答案不唯一：校验结果树的前序、后序遍历与输入一致。
func TestConstructFromPrePost(t *testing.T) {
	tests := []struct {
		name      string
		pre, post []int
	}{
		{"示例1", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}},
		{"示例2-单节点", []int{1}, []int{1}},
		{"空树", []int{}, []int{}},
		{"链（左斜或右斜）", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"两个节点", []int{2, 1}, []int{1, 2}},
		{"不规则", []int{1, 2, 4, 3, 5, 6, 7}, []int{4, 2, 6, 7, 5, 3, 1}},
		{"左深右浅", []int{1, 2, 3, 4, 5, 6}, []int{4, 5, 3, 2, 6, 1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.pre, tt.post)
				if p := preorderVals(got); !slices.Equal(p, tt.pre) {
					t.Errorf("%s: 前序 = %v, want %v", name, p, tt.pre)
				}
				if p := postorderVals(got); !slices.Equal(p, tt.post) {
					t.Errorf("%s: 后序 = %v, want %v", name, p, tt.post)
				}
			})
		}
	}
}
