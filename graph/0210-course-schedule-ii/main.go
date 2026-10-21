package courseschedule2

import "slices"

// 210. 课程表 II
// https://leetcode.cn/problems/course-schedule-ii/
// 共有 numCourses 门课，prerequisites[i] = [a, b] 表示学 a 之前必须先学 b。
// 返回一种可以学完所有课程的学习顺序；若不可能完成，返回空数组。
//
// 核心思路：求有向图 b -> a 的一个拓扑序。Kahn 算法按入度为 0 的顺序弹出节点即为拓扑序；
// DFS 后序（节点所有后继都完成后才记录）反转后也是拓扑序，同时用三色标记判环。

// findOrder 解法一：Kahn 拓扑排序 ⭐ 最优方案
// 时间复杂度 O(V+E)，空间复杂度 O(V+E)。
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indeg[p[0]]++
	}
	order := make([]int, 0, numCourses)
	for i, d := range indeg {
		if d == 0 {
			order = append(order, i)
		}
	}
	// order 本身兼作队列：head 之前是已弹出的节点
	for head := 0; head < len(order); head++ {
		for _, v := range graph[order[head]] {
			indeg[v]--
			if indeg[v] == 0 {
				order = append(order, v)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}

// findOrderDFS 解法二：DFS 三色标记 + 后序反转（对照用）
// 0=未访问，1=递归栈中，2=已完成。节点完成时追加到 post，最终反转 post。
// 时间复杂度 O(V+E)，空间复杂度 O(V+E)。
func findOrderDFS(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	color := make([]int, numCourses)
	post := make([]int, 0, numCourses)
	var dfs func(u int) bool // 返回 false 表示发现环
	dfs = func(u int) bool {
		color[u] = 1
		for _, v := range graph[u] {
			if color[v] == 1 || (color[v] == 0 && !dfs(v)) {
				return false
			}
		}
		color[u] = 2
		post = append(post, u)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if color[i] == 0 && !dfs(i) {
			return []int{}
		}
	}
	slices.Reverse(post)
	return post
}
