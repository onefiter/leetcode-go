package clonegraph

import (
	"reflect"
	"testing"
)

// buildGraph 按 LeetCode 邻接表构造图（节点值从 1 开始），返回节点 1。
func buildGraph(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adj)+1)
	for i := 1; i <= len(adj); i++ {
		nodes[i] = &Node{Val: i}
	}
	for i, nbs := range adj {
		for _, v := range nbs {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[v])
		}
	}
	return nodes[1]
}

// collect BFS 收集从 node 出发可达的全部节点。
func collect(node *Node) []*Node {
	if node == nil {
		return nil
	}
	seen := map[*Node]bool{node: true}
	order := []*Node{node}
	for i := 0; i < len(order); i++ {
		for _, nb := range order[i].Neighbors {
			if !seen[nb] {
				seen[nb] = true
				order = append(order, nb)
			}
		}
	}
	return order
}

// toAdj 把图序列化为邻接表（按节点值索引），用于比较结构。
func toAdj(node *Node) [][]int {
	nodes := collect(node)
	adj := make([][]int, len(nodes))
	for _, n := range nodes {
		row := []int{}
		for _, nb := range n.Neighbors {
			row = append(row, nb.Val)
		}
		adj[n.Val-1] = row
	}
	return adj
}

var solutions = map[string]func(*Node) *Node{
	"cloneGraph":    cloneGraph,
	"cloneGraphBFS": cloneGraphBFS,
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		adj  [][]int
	}{
		{"示例1-四节点环", [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		{"示例2-单节点无邻居", [][]int{{}}},
		{"示例3-空图", [][]int{}},
		{"两节点", [][]int{{2}, {1}}},
		{"三角形", [][]int{{2, 3}, {1, 3}, {1, 2}}},
		{"星形", [][]int{{2, 3, 4, 5}, {1}, {1}, {1}, {1}}},
		{"链", [][]int{{2}, {1, 3}, {2, 4}, {3}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				orig := buildGraph(tt.adj)
				got := fn(orig)
				if orig == nil {
					if got != nil {
						t.Fatalf("%s(nil) = %v, want nil", name, got)
					}
					return
				}
				if got == nil {
					t.Fatalf("%s(%v) = nil", name, tt.adj)
				}
				// 结构一致
				if gotAdj := toAdj(got); !reflect.DeepEqual(gotAdj, tt.adj) {
					t.Errorf("%s(%v) 结构 = %v, want %v", name, tt.adj, gotAdj, tt.adj)
				}
				// 深拷贝：克隆图中不能出现任何原图节点指针
				origSet := map[*Node]bool{}
				for _, n := range collect(orig) {
					origSet[n] = true
				}
				for _, n := range collect(got) {
					if origSet[n] {
						t.Errorf("%s(%v) 节点 %d 与原图共享指针", name, tt.adj, n.Val)
					}
				}
				// 原图不能被修改
				if origAdj := toAdj(orig); !reflect.DeepEqual(origAdj, tt.adj) {
					t.Errorf("%s 修改了原图: %v", name, origAdj)
				}
			})
		}
	}
}
