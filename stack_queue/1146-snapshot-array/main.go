package snapshotarray

import "sort"

// 1146. 快照数组
// https://leetcode.cn/problems/snapshot-array/
// 实现一个长度为 length 的「可拍快照」数组（初始全 0）：
//   - Set(index, val)：设置当前值；
//   - Snap()：拍一次快照，返回快照编号（从 0 开始，即已拍次数 - 1）；
//   - Get(index, snapId)：返回第 snapId 次快照时 index 处的值。
//
// 核心思路：每次 Snap 都复制整个数组太浪费。只记录「变化」：
// 每个下标维护一个按 snapId 递增的 (snapId, val) 列表，Set 时追加（同一 snapId 就覆盖）。
// Get 时在该下标的列表里二分找「最后一个 snapId <= 查询 snapId」的记录，没有就是 0。

type record struct {
	snapID int
	val    int
}

// SnapshotArray 解法一：每个下标存修改历史 + 二分查找 ⭐ 最优方案
// 时间复杂度 Set/Snap O(1)，Get O(log m)（m 为该下标的修改次数）；
// 空间复杂度 O(length + Set 调用次数)。
type SnapshotArray struct {
	history [][]record
	snapID  int // 当前（尚未拍下的）快照编号
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{history: make([][]record, length)}
}

func (this *SnapshotArray) Set(index int, val int) {
	h := this.history[index]
	// 同一个快照周期内多次 Set，只保留最后一次。
	if n := len(h); n > 0 && h[n-1].snapID == this.snapID {
		h[n-1].val = val
		return
	}
	this.history[index] = append(h, record{this.snapID, val})
}

func (this *SnapshotArray) Snap() int {
	this.snapID++
	return this.snapID - 1
}

func (this *SnapshotArray) Get(index int, snapId int) int {
	h := this.history[index]
	// 找第一个 snapID > snapId 的位置，它前一个就是答案。
	i := sort.Search(len(h), func(i int) bool { return h[i].snapID > snapId })
	if i == 0 {
		return 0 // 该快照之前从未被设置过
	}
	return h[i-1].val
}
