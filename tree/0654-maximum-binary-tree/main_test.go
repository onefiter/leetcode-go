package maxbinarytree

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

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

var solutions = map[string]func([]int) *TreeNode{
	"constructMaximumBinaryTree":          constructMaximumBinaryTree,
	"constructMaximumBinaryTreeRecursive": constructMaximumBinaryTreeRecursive,
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{3, 2, 1, 6, 0, 5}, []int{6, 3, 5, null, 2, 0, null, null, 1}},
		{"示例2", []int{3, 2, 1}, []int{3, null, 2, null, 1}},
		{"空数组", []int{}, []int{}},
		{"单元素", []int{7}, []int{7}},
		{"严格递增-左斜链", []int{1, 2, 3}, []int{3, 2, null, 1}},
		{"最大值在中间", []int{1, 3, 2}, []int{3, 1, 2}},
		{"含负数", []int{-1, -5, 0, -2}, []int{0, -1, -2, null, -5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := serialize(fn(tt.nums)); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}

// TestAgainstRecursive 用一批不重复数组让单调栈解与按定义递归解对拍。
func TestAgainstRecursive(t *testing.T) {
	seed := 7
	for n := 0; n <= 40; n++ {
		perm := make([]int, n)
		for i := range perm {
			perm[i] = i
		}
		for i := n - 1; i > 0; i-- { // 确定性的 Fisher-Yates 洗牌
			seed = (seed*1103515245 + 12345) % 2147483648
			j := seed % (i + 1)
			perm[i], perm[j] = perm[j], perm[i]
		}
		want := serialize(constructMaximumBinaryTreeRecursive(perm))
		if got := serialize(constructMaximumBinaryTree(perm)); !slices.Equal(got, want) {
			t.Fatalf("n=%d nums=%v: 单调栈 %v, 递归 %v", n, perm, got, want)
		}
	}
}
