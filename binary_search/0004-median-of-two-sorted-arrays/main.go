package mediantwosorted

import "math"

// 4. 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/
// 给定两个从小到大排好序的数组 nums1、nums2，求合并后的中位数。
// 例如 [1,3] 和 [2] -> 2.0；[1,2] 和 [3,4] -> 2.5。
//
// 一句话理解题意：不真的合并，也要找到「合并后排在正中间」的那一个（或两个）数。
//
// 核心思路：中位数把所有数切成「左半边」和「右半边」两堆，
// 左半边每个数都 <= 右半边每个数。在短数组上二分切口位置 i，
// 长数组的切口 j 由「左半边总个数固定」直接算出，检查切口两侧是否交叉合法即可。
// 约定：两个数组都为空时返回 0。

// findMedianSortedArrays 解法一：二分划分 ⭐ 最优方案
// 在较短的数组 nums1 上二分切口 i（左边取 i 个），nums2 取 j = half-i 个，
// half = (m+n+1)/2 保证奇数时左半边多一个。
// 设 l1=nums1[i-1], r1=nums1[i], l2=nums2[j-1], r2=nums2[j]（越界用 ±∞），
// 合法切分要求 l1 <= r2 且 l2 <= r1：
//   - l1 > r2：nums1 左边拿多了，i 往左；
//   - l2 > r1：nums1 左边拿少了，i 往右。
//
// 合法时：奇数个取 max(l1,l2)，偶数个取 (max(l1,l2)+min(r1,r2))/2。
// 时间复杂度 O(log(min(m,n)))，空间复杂度 O(1)。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	if m+n == 0 {
		return 0
	}
	half := (m + n + 1) / 2
	lo, hi := 0, m
	for lo <= hi {
		i := lo + (hi-lo)/2
		j := half - i
		l1, r1 := math.MinInt, math.MaxInt
		if i > 0 {
			l1 = nums1[i-1]
		}
		if i < m {
			r1 = nums1[i]
		}
		l2, r2 := math.MinInt, math.MaxInt
		if j > 0 {
			l2 = nums2[j-1]
		}
		if j < n {
			r2 = nums2[j]
		}
		switch {
		case l1 > r2:
			hi = i - 1
		case l2 > r1:
			lo = i + 1
		default:
			leftMax := max(l1, l2)
			if (m+n)%2 == 1 {
				return float64(leftMax)
			}
			return float64(leftMax+min(r1, r2)) / 2
		}
	}
	return 0 // 输入有序时不会走到这里
}

// findMedianSortedArraysKth 解法二：第 k 小递归
// 中位数就是第 (m+n+1)/2 小和第 (m+n+2)/2 小的平均（奇数时两者相同）。
// 求第 k 小：两个数组各看第 k/2 个元素，较小那一侧的前 k/2 个一定不是第 k 小，
// 整段扔掉，k 减去扔掉的个数，递归直到 k==1 或某个数组为空。
// 时间复杂度 O(log(m+n))，空间复杂度 O(log(m+n))（递归栈）。
func findMedianSortedArraysKth(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total == 0 {
		return 0
	}
	a := kth(nums1, nums2, (total+1)/2)
	b := kth(nums1, nums2, (total+2)/2)
	return float64(a+b) / 2
}

// kth 返回 a、b 合并后的第 k 小（k 从 1 开始）。
func kth(a, b []int, k int) int {
	if len(a) > len(b) {
		return kth(b, a, k)
	}
	if len(a) == 0 {
		return b[k-1]
	}
	if k == 1 {
		return min(a[0], b[0])
	}
	i := min(len(a), k/2)
	j := min(len(b), k/2)
	if a[i-1] < b[j-1] {
		return kth(a[i:], b, k-i)
	}
	return kth(a, b[j:], k-j)
}

// findMedianSortedArraysEnumeration 解法三：归并合并（对照用）
// 像归并排序那样把两个数组合并成一个有序数组，直接取中间。
// 时间复杂度 O(m+n)，空间复杂度 O(m+n)。
func findMedianSortedArraysEnumeration(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)
	n := len(merged)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return float64(merged[n/2])
	}
	return float64(merged[n/2-1]+merged[n/2]) / 2
}
