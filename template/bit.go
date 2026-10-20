package template

// BIT 树状数组（下标从 1 开始）：单点更新、前缀和查询，均为 O(log n)。
type BIT struct {
	n    int
	tree []int
}

func NewBIT(n int) *BIT {
	return &BIT{n: n, tree: make([]int, n+1)}
}

// Add 第 i 个位置加上 delta。
func (b *BIT) Add(i, delta int) {
	for ; i <= b.n; i += i & -i {
		b.tree[i] += delta
	}
}

// Query 返回 [1, i] 的前缀和。
func (b *BIT) Query(i int) int {
	s := 0
	for ; i > 0; i -= i & -i {
		s += b.tree[i]
	}
	return s
}

// RangeQuery 返回 [l, r] 的区间和。
func (b *BIT) RangeQuery(l, r int) int {
	return b.Query(r) - b.Query(l-1)
}

// RangeBIT 支持区间加、区间求和的树状数组（下标从 1 开始）。
// 原理：维护差分数组 d 的两个 BIT：sum(d[i]) 与 sum(d[i]*i)。
type RangeBIT struct {
	b1, b2 *BIT
}

func NewRangeBIT(n int) *RangeBIT {
	return &RangeBIT{b1: NewBIT(n), b2: NewBIT(n)}
}

func (r *RangeBIT) add(i, delta int) {
	r.b1.Add(i, delta)
	r.b2.Add(i, delta*i)
}

// RangeAdd 区间 [l, r] 每个元素加上 delta。
func (r *RangeBIT) RangeAdd(l, rr, delta int) {
	r.add(l, delta)
	if rr+1 <= r.b1.n {
		r.add(rr+1, -delta)
	}
}

// prefix 返回 [1, i] 的前缀和。
func (r *RangeBIT) prefix(i int) int {
	return (i+1)*r.b1.Query(i) - r.b2.Query(i)
}

// RangeQuery 返回 [l, r] 的区间和。
func (r *RangeBIT) RangeQuery(l, rr int) int {
	return r.prefix(rr) - r.prefix(l-1)
}
