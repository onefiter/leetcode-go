package firstunique

// 1429. 第一个唯一数字
// https://leetcode.cn/problems/first-unique-number/
// 给一个整数队列，实现：
//   - Constructor(nums)：用 nums 初始化队列；
//   - ShowFirstUnique()：返回队列里第一个「只出现一次」的数，没有返回 -1；
//   - Add(value)：把 value 加到队尾。
//
// 核心思路：哈希表记每个数出现次数；另用一个队列按到达顺序存「首次出现」的数。
// 查询时从队头开始，把出现次数 > 1 的数弹掉（懒删除），队头就是答案。
// 一个数一旦重复就永远不会再唯一，所以弹掉是安全的。

// FirstUnique 解法一：哈希计数 + 队列懒删除 ⭐ 最优方案
// 每个数最多入队、出队一次。
// 时间复杂度 Constructor O(n)，Add O(1)，ShowFirstUnique 均摊 O(1)；空间复杂度 O(n)。
type FirstUnique struct {
	cnt map[int]int
	q   []int
}

func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{cnt: make(map[int]int)}
	for _, x := range nums {
		fu.Add(x)
	}
	return fu
}

func (this *FirstUnique) ShowFirstUnique() int {
	// 懒删除：队头已经重复了就丢掉。
	for len(this.q) > 0 && this.cnt[this.q[0]] > 1 {
		this.q = this.q[1:]
	}
	if len(this.q) == 0 {
		return -1
	}
	return this.q[0]
}

func (this *FirstUnique) Add(value int) {
	this.cnt[value]++
	if this.cnt[value] == 1 {
		this.q = append(this.q, value) // 只有首次出现才入队
	}
}
