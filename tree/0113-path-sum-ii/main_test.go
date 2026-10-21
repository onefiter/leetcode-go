package pathsumii

import (
	"math"
	"slices"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) [][]int{
	"pathSum":    pathSum,
	"pathSumBFS": pathSumBFS,
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

// normalize 对路径列表按字典序排序，消除不同解法输出顺序的差异。
func normalize(paths [][]int) [][]int {
	out := slices.Clone(paths)
	slices.SortFunc(out, slices.Compare[[]int])
	return out
}

func TestPathSum(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		target int
		want   [][]int
	}{
		{"示例1", []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}, 22, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{"示例2", []int{1, 2, 3}, 5, [][]int{}},
		{"示例3", []int{1, 2}, 0, [][]int{}},
		{"空树", []int{}, 0, [][]int{}},
		{"单节点", []int{7}, 7, [][]int{{7}}},
		{"左斜链", []int{1, 2, null, 3}, 6, [][]int{{1, 2, 3}}},
		{"右斜链", []int{1, null, 2, null, 3}, 6, [][]int{{1, 2, 3}}},
		{"完全二叉树-多条", []int{1, 2, 2, 3, 3, 3, 3}, 6, [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}}},
		{"负数", []int{-2, null, -3}, -5, [][]int{{-2, -3}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(buildTree(tt.vals), tt.target))
				want := normalize(tt.want)
				if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.target, got, want)
				}
			})
		}
	}
}

// TestPathSumDFSOrder 解法一按 DFS 左先输出，顺序应与题目示例一致。
func TestPathSumDFSOrder(t *testing.T) {
	root := buildTree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1})
	got := pathSum(root, 22)
	want := [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Errorf("pathSum order = %v, want %v", got, want)
	}
}
