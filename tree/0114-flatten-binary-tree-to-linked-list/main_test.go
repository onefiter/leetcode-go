package flattentree

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode){
	"flatten":            flatten,
	"flattenRecursive":   flattenRecursive,
	"flattenEnumeration": flattenEnumeration,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		cur := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			cur.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			cur.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}

// rightChain 沿 Right 收集值；若遇到 Left 非 nil 返回 ok=false。
// 设置节点上限防止错误实现形成环导致死循环。
func rightChain(root *TreeNode) (vals []int, ok bool) {
	vals = []int{}
	for n, steps := root, 0; n != nil; n, steps = n.Right, steps+1 {
		if n.Left != nil || steps > 10000 {
			return vals, false
		}
		vals = append(vals, n.Val)
	}
	return vals, true
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int // 前序序列
	}{
		{"示例1", []int{1, 2, 5, 3, 4, null, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"示例2-空树", []int{}, []int{}},
		{"示例3-单节点", []int{0}, []int{0}},
		{"左斜链", []int{1, 2, null, 3, null, 4}, []int{1, 2, 3, 4}},
		{"右斜链", []int{1, null, 2, null, 3}, []int{1, 2, 3}},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 3, 6, 7}},
		{"之字形", []int{1, 2, null, null, 3, 4}, []int{1, 2, 3, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals) // 每个解法重新构造树
				fn(root)
				got, ok := rightChain(root)
				if !ok {
					t.Fatalf("%s(%v): 存在非 nil 的 Left 指针或成环", name, tt.vals)
				}
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
