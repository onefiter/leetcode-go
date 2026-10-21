package foursumii

// 454. 四数相加 II
// https://leetcode.cn/problems/4sum-ii/
// 给定四个长度均为 n 的整数数组 nums1, nums2, nums3, nums4，
// 统计有多少个四元组 (i, j, k, l) 满足
//   nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0。
//
// 核心思路：把四个数组「两两分组」，用哈希表把 O(n^4) 枚举降到 O(n^2)。
//   - 先枚举 nums1、nums2 的所有两两之和，用哈希表记录每个和出现的次数；
//   - 再枚举 nums3、nums4 的所有两两之和 s，查哈希表里有多少个 -s，
//     这些就能和当前组合凑成 0，累加其出现次数即可。

// fourSumCount 解法一：分组 + 哈希表 ⭐ 最优方案
// 把 nums1[i]+nums2[j] 的所有结果计数存入 map；
// 再遍历 nums3[k]+nums4[l]，查 map 中 -(nums3[k]+nums4[l]) 的次数并累加。
// 时间复杂度 O(n^2)：两次两重循环。
// 空间复杂度 O(n^2)：哈希表最多存 n^2 个不同的和。
func fourSumCount(nums1, nums2, nums3, nums4 []int) int {
	sumCount := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			sumCount[a+b]++
		}
	}
	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			count += sumCount[-(c + d)]
		}
	}
	return count
}

// fourSumCountBrute 解法二：四重循环暴力枚举（对照用）
// 直接枚举所有四元组并判断和是否为 0，数据量稍大就会超时。
// 时间复杂度 O(n^4)，空间复杂度 O(1)。
func fourSumCountBrute(nums1, nums2, nums3, nums4 []int) int {
	count := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			for _, c := range nums3 {
				for _, d := range nums4 {
					if a+b+c+d == 0 {
						count++
					}
				}
			}
		}
	}
	return count
}
