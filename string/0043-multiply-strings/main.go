package multiplystrings

// 43. 字符串相乘
// https://leetcode.cn/problems/multiply-strings/
// 给定两个以字符串形式表示的非负整数 num1 和 num2，
// 返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 不能使用任何内置的大整数库，也不能直接将输入转换为整数。
// 例如 num1="2", num2="3" -> "6"；num1="123", num2="456" -> "56088"。
//
// 一句话理解题意：用小学竖式乘法，逐位相乘、错位累加，全程用数组模拟。
//
// 核心思路：长度为 m 和 n 的两数相乘，结果最多 m+n 位。
// num1[i] * num2[j] 的结果会累加到结果数组的第 i+j 和 i+j+1 位上，
// 最后统一处理进位，再去掉前导零转成字符串。

// multiply 【解法1】竖式逐位相乘 + 错位累加 ⭐ 最优方案
// 思路：用长度 m+n 的 int 数组 res 保存每一位（含进位前的累加值）。
// num1[i]*num2[j] 落在 res[i+j]（高位）和 res[i+j+1]（低位）。
// 先把所有位乘积累加进去，再从低位到高位统一处理进位。
// 时间复杂度 O(m*n)：双重循环逐位相乘。
// 空间复杂度 O(m+n)：结果数组。
func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		d1 := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			d2 := int(num2[j] - '0')
			// 乘积落在 res[i+j+1]（低位），进位加到 res[i+j]。
			sum := res[i+j+1] + d1*d2
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	// 跳过前导零（结果非零，所以最多跳过一个）。
	start := 0
	for start < len(res)-1 && res[start] == 0 {
		start++
	}
	buf := make([]byte, 0, len(res)-start)
	for _, d := range res[start:] {
		buf = append(buf, byte(d)+'0')
	}
	return string(buf)
}
