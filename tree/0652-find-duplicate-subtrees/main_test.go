package duplicatesubtrees

import (
	"fmt"
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

// serialize 层序序列化回 LeetCode 数组（去掉末尾 null）。
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

// normalize 把返回的每个子树根序列化成字符串后排序，消除返回顺序的影响。
func normalize(roots []*TreeNode) []string {
	out := make([]string, 0, len(roots))
	for _, r := range roots {
		out = append(out, fmt.Sprint(serialize(r)))
	}
	slices.Sort(out)
	return out
}

var solutions = map[string]func(*TreeNode) []*TreeNode{
	"findDuplicateSubtrees":          findDuplicateSubtrees,
	"findDuplicateSubtreesSerialize": findDuplicateSubtreesSerialize,
}

func TestFindDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want [][]int // 每类重复子树的层序数组
	}{
		{"示例1", []int{1, 2, 3, 4, null, 2, 4, null, null, 4}, [][]int{{2, 4}, {4}}},
		{"示例2", []int{2, 1, 1}, [][]int{{1}}},
		{"示例3", []int{2, 2, 2, 3, null, 3, null}, [][]int{{2, 3}, {3}}},
		{"空树", []int{}, nil},
		{"单节点", []int{1}, nil},
		{"无重复", []int{1, 2, 3}, nil},
		{"值相同结构不同不算重复", []int{0, 0, 0, 0, null, null, 0}, [][]int{{0}}},
		{"出现三次只返回一次", []int{1, 2, 2, 2}, [][]int{{2}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				want := make([]string, 0, len(tt.want))
				for _, w := range tt.want {
					want = append(want, fmt.Sprint(w))
				}
				slices.Sort(want)
				if got := normalize(fn(buildTree(tt.root))); !slices.Equal(got, want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.root, got, want)
				}
			})
		}
	}
}
