package serializebst

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
	"CodecInsert": func(root *TreeNode) *TreeNode {
		c := ConstructorInsert()
		return c.deserialize(c.serialize(root))
	},
}

func TestCodecRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"示例1", []int{2, 1, 3}},
		{"空树", []int{}},
		{"单节点", []int{1}},
		{"左斜链", []int{4, 3, null, 2, null, 1}},
		{"右斜链", []int{1, null, 2, null, 3, null, 4}},
		{"完全二叉树", []int{4, 2, 6, 1, 3, 5, 7}},
		{"含负数", []int{0, -5, 10, -8, -2, null, 20}},
		{"不平衡", []int{8, 3, 10, 1, 6, null, 14, null, null, 4, 7, 13}},
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

func TestCodecCompact(t *testing.T) {
	c := Constructor()
	if got := c.serialize(buildTree([]int{2, 1, 3})); got != "2,1,3" {
		t.Errorf("serialize = %q, want %q", got, "2,1,3")
	}
	if got := c.serialize(nil); got != "" {
		t.Errorf("serialize(nil) = %q, want empty", got)
	}
}
