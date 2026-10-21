package evaldivision

// 399. 除法求值
// https://leetcode.cn/problems/evaluate-division/
// 给定方程 equations[i] = [Ai, Bi] 与对应值 values[i]，表示 Ai / Bi = values[i]。
// 对每个查询 queries[j] = [Cj, Dj] 求 Cj / Dj；无法确定（含未出现的变量）时返回 -1.0。
//
// 核心思路：把变量看成节点，a / b = v 是一条带权边。
// 带权并查集：weight[x] 表示 x / parent[x]，路径压缩时把权值沿路径相乘，
// 查询 a / b 时若同根，答案为 weight[a] / weight[b]（都是相对根的比值）。
// 对照解法用 Floyd 求出所有变量两两之间的比值。

// calcEquation 解法一：带权并查集 ⭐ 最优方案
// 设变量数为 n、方程数 E、查询数 Q，时间复杂度 O((E+Q)*α(n))，空间复杂度 O(n)。
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	var parent []int
	var weight []float64 // weight[x] = x / parent[x]
	getID := func(s string) int {
		if i, ok := id[s]; ok {
			return i
		}
		id[s] = len(parent)
		parent = append(parent, len(parent))
		weight = append(weight, 1)
		return id[s]
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			root := find(parent[x])
			weight[x] *= weight[parent[x]] // 父节点已压缩到根，权值相乘得到 x / root
			parent[x] = root
		}
		return parent[x]
	}
	for i, eq := range equations {
		a, b := getID(eq[0]), getID(eq[1])
		ra, rb := find(a), find(b)
		if ra != rb {
			// 令 ra 挂到 rb 下：ra/rb = (a/b) * (b/rb) / (a/ra)
			parent[ra] = rb
			weight[ra] = values[i] * weight[b] / weight[a]
		}
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		a, okA := id[q[0]]
		b, okB := id[q[1]]
		if !okA || !okB || find(a) != find(b) {
			res[i] = -1
			continue
		}
		res[i] = weight[a] / weight[b]
	}
	return res
}

// calcEquationFloyd 解法二：Floyd 预处理全源比值（对照用）
// dist[i][j] = dist[i][k] * dist[k][j]，0 表示未知。适合查询非常多、变量很少的场景。
// 时间复杂度 O(n^3 + E + Q)，空间复杂度 O(n^2)。
func calcEquationFloyd(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, eq := range equations {
		for _, s := range eq {
			if _, ok := id[s]; !ok {
				id[s] = len(id)
			}
		}
	}
	n := len(id)
	dist := make([][]float64, n)
	for i := range dist {
		dist[i] = make([]float64, n)
		dist[i][i] = 1
	}
	for i, eq := range equations {
		a, b := id[eq[0]], id[eq[1]]
		dist[a][b] = values[i]
		dist[b][a] = 1 / values[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if dist[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if dist[i][j] == 0 && dist[k][j] != 0 {
					dist[i][j] = dist[i][k] * dist[k][j]
				}
			}
		}
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		a, okA := id[q[0]]
		b, okB := id[q[1]]
		if !okA || !okB || dist[a][b] == 0 {
			res[i] = -1
			continue
		}
		res[i] = dist[a][b]
	}
	return res
}
