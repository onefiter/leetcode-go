package copyrandomlist

import (
	"slices"
	"testing"
)

var solutions = map[string]func(*Node) *Node{
	"copyRandomList":            copyRandomList,
	"copyRandomListEnumeration": copyRandomListEnumeration,
}

// pair 表示 LeetCode 输入格式中的 [val, random_index]，random 为 -1 表示 nil。
type pair struct{ val, random int }

// buildRandomList 按 [val, random_index] 构造链表，返回表头和节点切片。
func buildRandomList(ps []pair) (*Node, []*Node) {
	nodes := make([]*Node, len(ps))
	for i, p := range ps {
		nodes[i] = &Node{Val: p.val}
	}
	for i, p := range ps {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		if p.random >= 0 {
			nodes[i].Random = nodes[p.random]
		}
	}
	if len(nodes) == 0 {
		return nil, nodes
	}
	return nodes[0], nodes
}

// toPairs 把链表还原成 [val, random_index] 形式，便于比较。
func toPairs(head *Node) []pair {
	idx := map[*Node]int{}
	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		idx[cur] = i
		i++
	}
	res := []pair{}
	for cur := head; cur != nil; cur = cur.Next {
		r := -1
		if cur.Random != nil {
			r = idx[cur.Random]
		}
		res = append(res, pair{cur.Val, r})
	}
	return res
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name string
		list []pair
	}{
		{"示例1", []pair{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}},
		{"示例2", []pair{{1, 1}, {2, 1}}},
		{"示例3", []pair{{3, -1}, {3, 0}, {3, -1}}},
		{"空链表", []pair{}},
		{"单节点无random", []pair{{5, -1}}},
		{"单节点指向自己", []pair{{5, 0}}},
		{"全部指向尾节点", []pair{{1, 2}, {2, 2}, {3, 2}}},
		{"全部指向头节点", []pair{{1, 0}, {1, 0}, {1, 0}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 每次重新构造输入：交织解法会临时修改原链表。
				head, orig := buildRandomList(tt.list)
				got := fn(head)
				if gp := toPairs(got); !slices.Equal(gp, tt.list) {
					t.Errorf("%s 拷贝结果 = %v, want %v", name, gp, tt.list)
				}
				// 原链表必须被恢复原样。
				if op := toPairs(head); !slices.Equal(op, tt.list) {
					t.Errorf("%s 破坏了原链表: %v, want %v", name, op, tt.list)
				}
				// 必须是深拷贝：新链表中不能出现任何原节点。
				origSet := map[*Node]bool{}
				for _, n := range orig {
					origSet[n] = true
				}
				for cur := got; cur != nil; cur = cur.Next {
					if origSet[cur] || (cur.Random != nil && origSet[cur.Random]) {
						t.Errorf("%s 结果引用了原链表节点", name)
						break
					}
				}
			})
		}
	}
}
