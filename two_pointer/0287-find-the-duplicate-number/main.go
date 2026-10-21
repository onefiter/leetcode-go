package findduplicate

// 287. 寻找重复数
// https://leetcode.cn/problems/find-the-duplicate-number/
// 长度为 n+1 的数组 nums，元素都在 [1, n] 范围内，恰好有一个数重复（可能重复多次），
// 在「不修改数组」且只用 O(1) 额外空间的前提下找出这个重复数。
// 例如 [1,3,4,2,2] -> 2，[3,1,3,4,2] -> 3。
//
// 核心思路：把下标 i 看成链表节点，i -> nums[i] 看成 next 指针。
// 由于有重复值，必然有两个节点指向同一个节点，链表中就会出现环，环的入口就是重复数。

// findDuplicate 解法一：快慢指针 Floyd 判环 ⭐ 最优方案
// 第一阶段：slow 走一步、fast 走两步，直到相遇（一定在环内）；
// 第二阶段：把 slow 放回起点 0，两者都每次走一步，再次相遇处就是环入口 = 重复数。
// 时间复杂度 O(n)，空间复杂度 O(1)，不修改数组。
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// findDuplicateBinarySearch 解法二：二分答案
// 对值域 [1, n] 二分：统计 nums 中 <= mid 的元素个数 cnt。
// 如果没有重复落在 [1, mid]，cnt 最多为 mid；cnt > mid 说明重复数在 [1, mid]，否则在 (mid, n]。
// 时间复杂度 O(n log n)，空间复杂度 O(1)，不修改数组。
func findDuplicateBinarySearch(nums []int) int {
	lo, hi := 1, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		cnt := 0
		for _, x := range nums {
			if x <= mid {
				cnt++
			}
		}
		if cnt > mid {
			hi = mid // 重复数在 [lo, mid]
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// findDuplicateEnumeration 解法三：哈希集合（对照用，不满足 O(1) 空间要求）
// 遍历时记录出现过的数，第一次遇到已出现的数就是答案。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func findDuplicateEnumeration(nums []int) int {
	seen := make(map[int]bool, len(nums))
	for _, x := range nums {
		if seen[x] {
			return x
		}
		seen[x] = true
	}
	return -1
}
