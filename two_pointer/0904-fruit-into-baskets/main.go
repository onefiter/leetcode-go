package fruitbaskets

// 904. 水果成篮
// https://leetcode.cn/problems/fruit-into-baskets/
// 一排果树 fruits[i] 表示第 i 棵树的水果种类。你有两个篮子，每个篮子只能装一种水果（数量不限），
// 从任意一棵树开始向右连续采摘，遇到第三种水果就必须停下。返回最多能采摘多少个水果。
// 例如 fruits=[1,2,3,2,2] -> 4（采 [2,3,2,2]）。
//
// 一句话理解题意：求「最多包含两种不同数字」的最长连续子数组长度。
//
// 核心思路：可变长滑动窗口。右端点不断扩张，用哈希表记录窗口内每种水果的个数，
// 种类数超过 2 时收缩左端点，直到只剩 2 种。每次更新最大窗口长度。

// totalFruit 解法一：滑动窗口 + 计数哈希表 ⭐ 最优方案
// cnt 记录窗口内每种水果的数量，len(cnt) 就是种类数；某种水果数量减到 0 时从 map 删除。
// 时间复杂度 O(n)，空间复杂度 O(1)（map 中最多 3 个键）。
func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	ans, left := 0, 0
	for right, f := range fruits {
		cnt[f]++
		for len(cnt) > 2 {
			out := fruits[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// totalFruitEnumeration 解法二：枚举起点向右扩展（对照用）
// 固定起点 i，向右走并记录见过的种类，出现第三种时停下，记录长度。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func totalFruitEnumeration(fruits []int) int {
	ans := 0
	for i := range fruits {
		a, b, hasB := fruits[i], 0, false // 两个篮子里的种类，hasB 表示第二个篮子是否已用
		j := i
		for ; j < len(fruits); j++ {
			f := fruits[j]
			if f == a || (hasB && f == b) {
				continue
			}
			if !hasB {
				b, hasB = f, true
				continue
			}
			break // 出现第三种
		}
		ans = max(ans, j-i)
	}
	return ans
}
