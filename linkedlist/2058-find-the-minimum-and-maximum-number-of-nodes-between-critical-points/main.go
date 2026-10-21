package criticalpoints

import "leetcode-go/template"

// 2058. 找出临界点之间的最小和最大距离
// https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
// 「临界点」是局部极大值或局部极小值：节点值严格大于（或严格小于）前后两个节点。
// 首尾节点不可能是临界点。返回 [任意两个不同临界点的最小距离, 最大距离]，
// 临界点少于两个时返回 [-1,-1]。
//
// 核心思路：临界点的下标是递增出现的，
// 最大距离 = 最后一个临界点 - 第一个临界点；
// 最小距离 = 相邻两个临界点距离的最小值。
// 一次遍历，记录第一个、上一个临界点的下标即可。

type ListNode = template.ListNode

// nodesBetweenCriticalPoints 解法一：一次遍历 ⭐ 最优方案
// 用 prev、cur、cur.Next 三个连续节点判断 cur 是否为临界点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func nodesBetweenCriticalPoints(head *ListNode) []int {
	first, last := -1, -1
	minDist := -1
	prev, cur := head, head.Next
	for i := 1; cur != nil && cur.Next != nil; i++ {
		next := cur.Next
		if (cur.Val > prev.Val && cur.Val > next.Val) || (cur.Val < prev.Val && cur.Val < next.Val) {
			if first == -1 {
				first = i
			} else if minDist == -1 || i-last < minDist {
				minDist = i - last
			}
			last = i
		}
		prev, cur = cur, next
	}
	if minDist == -1 {
		return []int{-1, -1}
	}
	return []int{minDist, last - first}
}

// nodesBetweenCriticalPointsEnumeration 解法二：收集所有临界点 + 两两枚举（对照用）
// 先转数组，找出所有临界点下标，再枚举每一对求最小/最大距离。
// 时间复杂度 O(n + k²)，k 为临界点个数；空间复杂度 O(n)。
func nodesBetweenCriticalPointsEnumeration(head *ListNode) []int {
	var vals []int
	for p := head; p != nil; p = p.Next {
		vals = append(vals, p.Val)
	}
	var idx []int
	for i := 1; i+1 < len(vals); i++ {
		if (vals[i] > vals[i-1] && vals[i] > vals[i+1]) || (vals[i] < vals[i-1] && vals[i] < vals[i+1]) {
			idx = append(idx, i)
		}
	}
	if len(idx) < 2 {
		return []int{-1, -1}
	}
	lo, hi := len(vals), 0
	for i := range idx {
		for j := i + 1; j < len(idx); j++ {
			d := idx[j] - idx[i]
			lo = min(lo, d)
			hi = max(hi, d)
		}
	}
	return []int{lo, hi}
}
