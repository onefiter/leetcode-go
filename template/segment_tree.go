package template

// SegmentTree 带懒标记的线段树：区间加、区间求和（下标从 0 开始）。
type SegmentTree struct {
	n    int
	sum  []int
	lazy []int
}

func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	st := &SegmentTree{n: n, sum: make([]int, 4*n), lazy: make([]int, 4*n)}
	if n > 0 {
		st.build(nums, 1, 0, n-1)
	}
	return st
}

func (st *SegmentTree) build(nums []int, o, l, r int) {
	if l == r {
		st.sum[o] = nums[l]
		return
	}
	m := (l + r) / 2
	st.build(nums, o*2, l, m)
	st.build(nums, o*2+1, m+1, r)
	st.sum[o] = st.sum[o*2] + st.sum[o*2+1]
}

func (st *SegmentTree) apply(o, l, r, delta int) {
	st.sum[o] += delta * (r - l + 1)
	st.lazy[o] += delta
}

func (st *SegmentTree) pushDown(o, l, r int) {
	if st.lazy[o] != 0 {
		m := (l + r) / 2
		st.apply(o*2, l, m, st.lazy[o])
		st.apply(o*2+1, m+1, r, st.lazy[o])
		st.lazy[o] = 0
	}
}

// Update 区间 [ql, qr] 每个元素加上 delta。
func (st *SegmentTree) Update(ql, qr, delta int) {
	st.update(1, 0, st.n-1, ql, qr, delta)
}

func (st *SegmentTree) update(o, l, r, ql, qr, delta int) {
	if ql <= l && r <= qr {
		st.apply(o, l, r, delta)
		return
	}
	st.pushDown(o, l, r)
	m := (l + r) / 2
	if ql <= m {
		st.update(o*2, l, m, ql, qr, delta)
	}
	if qr > m {
		st.update(o*2+1, m+1, r, ql, qr, delta)
	}
	st.sum[o] = st.sum[o*2] + st.sum[o*2+1]
}

// Query 返回区间 [ql, qr] 的和。
func (st *SegmentTree) Query(ql, qr int) int {
	return st.query(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) query(o, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return st.sum[o]
	}
	st.pushDown(o, l, r)
	m := (l + r) / 2
	res := 0
	if ql <= m {
		res += st.query(o*2, l, m, ql, qr)
	}
	if qr > m {
		res += st.query(o*2+1, m+1, r, ql, qr)
	}
	return res
}
