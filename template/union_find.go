package template

// UnionFind 并查集：路径压缩 + 按大小合并。
type UnionFind struct {
	parent []int
	size   []int
	Count  int // 连通分量个数
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size, Count: n}
}

func (uf *UnionFind) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

// Union 合并 x 和 y 所在集合，若原本已连通返回 false。
func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return false
	}
	if uf.size[rx] < uf.size[ry] {
		rx, ry = ry, rx
	}
	uf.parent[ry] = rx
	uf.size[rx] += uf.size[ry]
	uf.Count--
	return true
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Size 返回 x 所在集合的大小。
func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}
