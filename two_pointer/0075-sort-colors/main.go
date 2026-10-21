package sortcolors

// 75. 颜色分类
// https://leetcode.cn/problems/sort-colors/
// 数组只含 0、1、2（红、白、蓝），原地排序使相同颜色相邻，并按 0、1、2 顺序排列。
// 不能使用库排序函数。例如 [2,0,2,1,1,0] -> [0,0,1,1,2,2]。
//
// 核心思路（荷兰国旗问题）：把数组划成三段 [0..p0) 全是 0，[p0..i) 全是 1，(p2..n-1] 全是 2，
// 中间 [i..p2] 是未处理区。扫描指针 i 每看到一个数就把它丢到对应区域。

// sortColors 解法一：荷兰国旗三指针 ⭐ 最优方案
// - nums[i]==0：和 p0 交换，p0++、i++（换过来的一定是 1 或就是自己，已处理过）；
// - nums[i]==1：i++；
// - nums[i]==2：和 p2 交换，p2--，i 不动（换过来的数还没检查过）。
// 时间复杂度 O(n)：一趟扫描；空间复杂度 O(1)。
func sortColors(nums []int) {
	p0, i, p2 := 0, 0, len(nums)-1
	for i <= p2 {
		switch nums[i] {
		case 0:
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		case 1:
			i++
		default:
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
	}
}

// sortColorsEnumeration 解法二：计数排序（对照用）
// 先数出 0、1、2 各有多少个，再按顺序回填。需要两趟扫描。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func sortColorsEnumeration(nums []int) {
	var cnt [3]int
	for _, x := range nums {
		cnt[x]++
	}
	idx := 0
	for c := range 3 {
		for range cnt[c] {
			nums[idx] = c
			idx++
		}
	}
}
