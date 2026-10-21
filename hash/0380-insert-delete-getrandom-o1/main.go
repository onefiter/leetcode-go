package randomset

import "math/rand"

// 380. O(1) 时间插入、删除和获取随机元素
// https://leetcode.cn/problems/insert-delete-getrandom-o1/
// 实现 RandomizedSet：Insert(val) 不存在时插入并返回 true；Remove(val) 存在时删除并返回 true；
// GetRandom() 等概率随机返回一个现有元素（调用时保证集合非空）。所有操作平均 O(1)。
//
// 核心思路：变长数组 + 哈希表。
//   - 数组存元素，支持 O(1) 按下标随机访问；
//   - 哈希表 val -> 下标，支持 O(1) 判断存在与定位；
//   - 删除时把末尾元素搬到被删位置并更新其下标，再截掉末尾，避免数组中间删除的 O(n) 搬移。

// RandomizedSet 解法一：数组 + 哈希表，删除时与末尾交换 ⭐ 最优方案
// 时间复杂度：Insert / Remove / GetRandom 均为 O(1)（平均）；空间复杂度 O(n)。
type RandomizedSet struct {
	nums    []int
	indices map[int]int // val -> 在 nums 中的下标
}

func Constructor() RandomizedSet {
	return RandomizedSet{indices: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.indices[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[idx] = this.nums[last] // 末尾元素搬到被删位置
	this.indices[this.nums[idx]] = idx
	this.nums = this.nums[:last]
	delete(this.indices, val) // 必须在更新下标之后删，兼容 val 恰好是末尾元素的情况
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

// RandomizedSetMap 解法二：只用哈希表，随机时遍历 map（对照用）
// 先随机一个序号 k，再遍历 map 走到第 k 个元素返回（Go 的 map 遍历顺序本身不保证均匀随机）。
// 时间复杂度：Insert / Remove O(1)，GetRandom O(n)；空间复杂度 O(n)。
type RandomizedSetMap struct {
	set map[int]struct{}
}

func NewRandomizedSetMap() RandomizedSetMap {
	return RandomizedSetMap{set: map[int]struct{}{}}
}

func (this *RandomizedSetMap) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = struct{}{}
	return true
}

func (this *RandomizedSetMap) Remove(val int) bool {
	if _, ok := this.set[val]; !ok {
		return false
	}
	delete(this.set, val)
	return true
}

func (this *RandomizedSetMap) GetRandom() int {
	k := rand.Intn(len(this.set))
	for v := range this.set {
		if k == 0 {
			return v
		}
		k--
	}
	panic("unreachable")
}
