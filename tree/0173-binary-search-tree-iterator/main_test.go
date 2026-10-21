package bstiterator

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

type iterator interface {
	Next() int
	HasNext() bool
}

var solutions = map[string]func(*TreeNode) iterator{
	"BSTIterator": func(root *TreeNode) iterator {
		it := Constructor(root)
		return &it
	},
	"BSTIteratorEnumeration": func(root *TreeNode) iterator {
		it := ConstructorEnumeration(root)
		return &it
	},
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

// TestBSTIteratorExample 按官方示例的操作序列逐步校验，包括交替调用 HasNext。
func TestBSTIteratorExample(t *testing.T) {
	ops := []string{"next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"}
	want := []int{3, 7, 1, 9, 1, 15, 1, 20, 0} // hasNext: 1=true, 0=false
	for name, newIt := range solutions {
		it := newIt(buildTree([]int{7, 3, 15, null, null, 9, 20}))
		for i, op := range ops {
			var got int
			if op == "next" {
				got = it.Next()
			} else if it.HasNext() {
				got = 1
			}
			if got != want[i] {
				t.Errorf("%s: 第 %d 步 %s = %d, want %d", name, i, op, got, want[i])
			}
		}
	}
}

func TestBSTIteratorDrain(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"空树", []int{}, []int{}},
		{"单节点", []int{1}, []int{1}},
		{"左斜链", []int{3, 2, null, 1}, []int{1, 2, 3}},
		{"右斜链", []int{1, null, 2, null, 3}, []int{1, 2, 3}},
		{"完全二叉树", []int{4, 2, 6, 1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"之字形", []int{5, 1, null, null, 4, 2, null, null, 3}, []int{1, 2, 3, 4, 5}},
	}
	for name, newIt := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				it := newIt(buildTree(tt.vals))
				got := []int{}
				for it.HasNext() {
					got = append(got, it.Next())
				}
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
