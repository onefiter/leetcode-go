package findclosest

import (
	"slices"
	"sort"
)

// 658. 找到 K 个最接近的元素
// https://leetcode.cn/problems/find-k-closest-elements/
// 给定有序数组 arr、整数 k 和 x，找出与 x 最接近的 k 个数，结果按升序返回。
// 「更接近」：|a-x| < |b-x|，或距离相等时 a < b（更小的优先）。
// 例如 arr=[1,2,3,4,5], k=4, x=3 -> [1,2,3,4]。
//
// 一句话理解题意：答案一定是 arr 中一段长度为 k 的连续区间，找这个区间的左端点。
//
// 核心思路：二分窗口左端点 left ∈ [0, n-k]。比较窗口 [mid, mid+k) 两侧的
// arr[mid] 和 arr[mid+k]：若 x-arr[mid] > arr[mid+k]-x，说明右边的更近，窗口应右移。

// findClosestElements 解法一：二分窗口左端点 ⭐ 最优方案
// 在 [0, n-k] 上二分：x-arr[mid] > arr[mid+k]-x 时左端点太靠左，lo = mid+1；
// 否则（包括相等，按规则取更小的）hi = mid。
// 用差值而非绝对值比较，可以正确处理 x 在窗口两侧之外的情况。
// 时间复杂度 O(log(n-k) + k)，空间复杂度 O(1)（不计返回值）。
func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-k
	for lo < hi {
		mid := lo + (hi-lo)/2
		if x-arr[mid] > arr[mid+k]-x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return slices.Clone(arr[lo : lo+k])
}

// findClosestElementsTwoPointers 解法二：双指针删除
// 从整个数组出发，每次比较两端，删掉离 x 更远的那个（相等时删右端，保留更小的），
// 直到剩下 k 个。
// 时间复杂度 O(n-k)，空间复杂度 O(1)（不计返回值）。
func findClosestElementsTwoPointers(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-1
	for r-l+1 > k {
		if x-arr[l] > arr[r]-x {
			l++
		} else {
			r--
		}
	}
	return slices.Clone(arr[l : r+1])
}

// findClosestElementsEnumeration 解法三：按距离排序（对照用）
// 复制一份按（距离, 值）稳定排序，取前 k 个，再升序排列。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func findClosestElementsEnumeration(arr []int, k int, x int) []int {
	a := slices.Clone(arr)
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	sort.SliceStable(a, func(i, j int) bool {
		di, dj := abs(a[i]-x), abs(a[j]-x)
		if di != dj {
			return di < dj
		}
		return a[i] < a[j]
	})
	res := a[:k]
	slices.Sort(res)
	return res
}
