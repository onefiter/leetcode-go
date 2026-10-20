package template

// Dirs4 上下左右四个方向。
var Dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// GridBFS 在网格中从 (sr, sc) 出发做 BFS，返回到每个格子的最短步数（不可达为 -1）。
// passable 判断格子是否可通行。
func GridBFS(m, n, sr, sc int, passable func(r, c int) bool) [][]int {
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	if !passable(sr, sc) {
		return dist
	}
	dist[sr][sc] = 0
	q := [][2]int{{sr, sc}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, d := range Dirs4 {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || dist[r][c] != -1 || !passable(r, c) {
				continue
			}
			dist[r][c] = dist[cur[0]][cur[1]] + 1
			q = append(q, [2]int{r, c})
		}
	}
	return dist
}

// GraphBFS 在邻接表图中从 start 出发做 BFS，返回到每个点的最短距离（不可达为 -1）。
func GraphBFS(graph [][]int, start int) []int {
	dist := make([]int, len(graph))
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 0
	q := []int{start}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range graph[u] {
			if dist[v] == -1 {
				dist[v] = dist[u] + 1
				q = append(q, v)
			}
		}
	}
	return dist
}
