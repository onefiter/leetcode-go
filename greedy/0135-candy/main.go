package candy

// 135. 分发糖果
// https://leetcode.cn/problems/candy/
// n 个孩子站成一排，ratings[i] 为评分。每个孩子至少分 1 颗糖，
// 相邻两个孩子中评分更高的必须拿到更多糖果，求最少需要的糖果数。
//
// 核心思路：把「相邻」约束拆成左规则和右规则分别满足。
// 左→右扫：若 ratings[i] > ratings[i-1]，left[i] = left[i-1]+1，否则为 1；
// 右→左扫：若 ratings[i] > ratings[i+1]，right[i] = right[i+1]+1，否则为 1。
// 每个孩子取 max(left[i], right[i])：两条规则都满足，且各自都是满足该规则的最小值，
// 所以取最大值就是同时满足两条规则的最小值。

// candy 解法一：两次遍历贪心 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(n)。
func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	for i := range left {
		left[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	total, right := 0, 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		total += max(left[i], right)
	}
	return total
}

// candyBrute 解法二：反复修正直到稳定（对照用）
// 所有人先发 1 颗，然后不断扫描，发现违反规则的相邻对就给评分高的补到对方+1，直到一轮无修改。
// 时间复杂度 O(n^2)（最多 n 轮，每轮 O(n)），空间复杂度 O(n)。
func candyBrute(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for changed := true; changed; {
		changed = false
		for i := 0; i < n; i++ {
			if i > 0 && ratings[i] > ratings[i-1] && c[i] <= c[i-1] {
				c[i] = c[i-1] + 1
				changed = true
			}
			if i < n-1 && ratings[i] > ratings[i+1] && c[i] <= c[i+1] {
				c[i] = c[i+1] + 1
				changed = true
			}
		}
	}
	total := 0
	for _, v := range c {
		total += v
	}
	return total
}
