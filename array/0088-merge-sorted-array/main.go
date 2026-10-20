package mergesortedarray

import "slices"

// merge 88. 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/
// 逆向双指针：从后往前填充 nums1，时间 O(m+n)，空间 O(1)。
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// mergeEnumeration 解法二：穷举/朴素做法（对照用）
// 把 nums2 直接接到 nums1 的有效部分后面，再整体排序。
// 逻辑最直白，但没利用两数组已有序这一点。
// 时间复杂度 O((m+n)log(m+n))，空间复杂度 O(1)（原地排序）。
func mergeEnumeration(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	slices.Sort(nums1[:m+n])
}
