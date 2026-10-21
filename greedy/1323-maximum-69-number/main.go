package maximum69

import (
	"strconv"
	"strings"
)

// 1323. 6 和 9 组成的最大数字
// https://leetcode.cn/problems/maximum-69-number/
// 给一个仅由数字 6 和 9 组成的正整数 num，最多翻转一位数字（6 变 9 或 9 变 6），返回能得到的最大数字。
//
// 核心思路：高位的权重大于所有低位之和，所以只需把「最高位的 6」翻成 9；
// 9 变 6 只会让数字变小，不翻也不会更差。若没有 6，原数即为最大。

// maximum69Number 解法一：数学法找最高位的 6 ⭐ 最优方案
// 时间复杂度 O(log num)，空间复杂度 O(1)。
func maximum69Number(num int) int {
	highest := -1 // 最高位 6 对应的位权（10 的幂）
	for x, base := num, 1; x > 0; x, base = x/10, base*10 {
		if x%10 == 6 {
			highest = base // 越往后越高位，最后一次更新即最高位
		}
	}
	if highest == -1 {
		return num
	}
	return num + 3*highest // 6 -> 9 增加 3 * 位权
}

// maximum69NumberBrute 解法二：枚举每一位翻转取最大（对照用）
// 转成字符串，逐位尝试翻转（6 和 9 互换），解析后取最大值，同时包含「不翻转」的情况。
// 时间复杂度 O(d^2)，d 为位数，空间复杂度 O(d)。
func maximum69NumberBrute(num int) int {
	s := strconv.Itoa(num)
	best := num
	for i := 0; i < len(s); i++ {
		flip := byte('6')
		if s[i] == '6' {
			flip = '9'
		}
		var b strings.Builder
		b.WriteString(s[:i])
		b.WriteByte(flip)
		b.WriteString(s[i+1:])
		v, _ := strconv.Atoi(b.String())
		best = max(best, v)
	}
	return best
}
