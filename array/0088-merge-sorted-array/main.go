package mergesortedarray

// merge 88. 合并两个有序数组
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
