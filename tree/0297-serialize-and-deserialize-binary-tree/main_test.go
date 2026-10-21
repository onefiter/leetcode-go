package serializebt

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

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

// levelOrder 把树按层序序列化回切片（去掉末尾的 null）。
func levelOrder(root *TreeNode) []int {
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, null)
			continue
		}
		res = append(res, node.Val)
		q = append(q, node.Left, node.Right)
	}
	for len(res) > 0 && res[len(res)-1] == null {
		res = res[:len(res)-1]
	}
	return res
}

var solutions = map[string]func(*TreeNode) *TreeNode{
	"Codec": func(root *TreeNode) *TreeNode {
		c := Constructor()
		return c.deserialize(c.serialize(root))
	},
	"CodecBFS": func(root *TreeNode) *TreeNode {
		c := ConstructorBFS()
		return c.deserialize(c.serialize(root))
	},
}

func TestCodecRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"示例1", []int{1, 2, 3, null, null, 4, 5}},
		{"空树", []int{}},
		{"单节点", []int{1}},
		{"含负数", []int{-1, -20, 300, null, -4000}},
		{"左斜链", []int{1, 2, null, 3, null, 4}},
		{"右斜链", []int{1, null, 2, null, 3, null, 4}},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}},
		{"多位数与 0", []int{12, 2, 0, null, 120}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := levelOrder(fn(buildTree(tt.vals)))
				want := levelOrder(buildTree(tt.vals))
				if !slices.Equal(got, want) {
					t.Errorf("%s round trip %v = %v, want %v", name, tt.vals, got, want)
				}
			})
		}
	}
}

func TestCodecFormat(t *testing.T) {
	c := Constructor()
	if got := c.serialize(buildTree([]int{1, 2, 3, null, null, 4, 5})); got != "1,2,#,#,3,4,#,#,5,#,#" {
		t.Errorf("Codec.serialize = %q", got)
	}
	b := ConstructorBFS()
	if got := b.serialize(buildTree([]int{1, 2, 3, null, null, 4, 5})); got != "1,2,3,#,#,4,5,#,#,#,#" {
		t.Errorf("CodecBFS.serialize = %q", got)
	}
}
