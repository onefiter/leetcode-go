package populatenextii

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*Node) *Node{
	"connect":    connect,
	"connectBFS": connectBFS,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *Node {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &Node{Val: vals[0]}
	queue := []*Node{root}
	for i := 1; i < len(vals); {
		cur := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			cur.Left = &Node{Val: vals[i]}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			cur.Right = &Node{Val: vals[i]}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}

// serializeByNext 按 LeetCode 输出格式序列化：每层沿 Next 走，层末追加 '#'（用 null 表示）。
// 每层的起点通过层序 BFS 找到，从而独立于 Next 指针校验整层。
func serializeByNext(root *Node) []int {
	res := []int{}
	level := []*Node{}
	if root != nil {
		level = append(level, root)
	}
	for len(level) > 0 {
		for n := level[0]; n != nil; n = n.Next {
			res = append(res, n.Val)
		}
		res = append(res, null)
		next := []*Node{}
		for _, n := range level {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		level = next
	}
	return res
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5, null, 7}, []int{1, null, 2, 3, null, 4, 5, 7, null}},
		{"示例2-空树", []int{}, []int{}},
		{"单节点", []int{1}, []int{1, null}},
		{"左斜链", []int{1, 2, null, 3}, []int{1, null, 2, null, 3, null}},
		{"右斜链", []int{1, null, 2, null, 3}, []int{1, null, 2, null, 3, null}},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, null, 2, 3, null, 4, 5, 6, 7, null}},
		{"跨越空缺连接", []int{1, 2, 3, 4, null, null, 5}, []int{1, null, 2, 3, null, 4, 5, null}},
		{"中间整段缺失", []int{1, 2, 3, 4, null, null, 5, 6, null, null, 7},
			[]int{1, null, 2, 3, null, 4, 5, null, 6, 7, null}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				root := buildTree(tt.vals)
				got := serializeByNext(fn(root))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
