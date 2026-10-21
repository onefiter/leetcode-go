package courseschedule

// 207. 课程表
// https://leetcode.cn/problems/course-schedule/
// 共有 numCourses 门课（0..numCourses-1），prerequisites[i] = [a, b] 表示学 a 之前必须先学 b。
// 判断是否可能完成所有课程。
//
// 核心思路：把课程看成有向图 b -> a，能学完所有课 ⇔ 图中无环。
// 拓扑排序（Kahn 入度法）能弹出全部节点即无环；也可以用 DFS 三色标记找回边。

// canFinish 解法一：Kahn 拓扑排序（BFS 入度法）⭐ 最优方案
// 时间复杂度 O(V+E)，空间复杂度 O(V+E)。
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indeg[p[0]]++
	}
	queue := []int{}
	for i, d := range indeg {
		if d == 0 {
			queue = append(queue, i)
		}
	}
	done := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		done++
		for _, v := range graph[u] {
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	// 有环时环上节点入度永远不为 0，无法全部弹出
	return done == numCourses
}

// canFinishDFS 解法二：DFS 三色标记（对照用）
// 0=未访问，1=递归栈中（灰），2=已完成（黑）。遇到灰色节点说明存在回边，即有环。
// 时间复杂度 O(V+E)，空间复杂度 O(V+E)。
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	color := make([]int, numCourses)
	var hasCycle func(u int) bool
	hasCycle = func(u int) bool {
		color[u] = 1
		for _, v := range graph[u] {
			if color[v] == 1 || (color[v] == 0 && hasCycle(v)) {
				return true
			}
		}
		color[u] = 2
		return false
	}
	for i := 0; i < numCourses; i++ {
		if color[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}
