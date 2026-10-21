package zigzag

import "strings"

// 6. Z 字形变换
// https://leetcode.cn/problems/zigzag-conversion/
// 将一个给定字符串 s 根据给定的行数 numRows，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING"，行数为 3 时排列如下：
//
//	P   A   H   N
//	A P L S I I G
//	Y   I   R
//
// 之后从左往右逐行读取，产生 "PAHNAPLSIIGYIR"。
// 请你实现这个将字符串进行指定行数变换的函数。
// 注意：当 numRows == 1 时没有折返，直接返回原字符串 s。
//
// 一句话理解题意：按 Z 字（竖直向下 + 斜向上）顺序把字符分配到各行，再逐行拼接。
//
// 核心思路：想象一个在 0 ~ numRows-1 之间来回弹跳的行号指针，
// 竖直向下走时行号 +1，到底后斜向上走时行号 -1，把每个字符追加到对应行。

// convert 【解法1】按行模拟弹跳 ⭐ 最优实用方案
// 思路：为每一行维护一个 Builder，用 row 记录当前字符所在行，
// 用 dir 记录方向（+1 向下 / -1 向上），到达首尾行时翻转方向。
// 时间复杂度 O(n)：每个字符处理一次。
// 空间复杂度 O(n)：存放结果各行字符。
func convert(s string, numRows int) string {
	// numRows==1 时没有折返，直接返回。
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]strings.Builder, numRows)
	row, dir := 0, 1
	for i := 0; i < len(s); i++ {
		rows[row].WriteByte(s[i])
		if row == 0 {
			dir = 1 // 到达顶行，转为向下
		} else if row == numRows-1 {
			dir = -1 // 到达底行，转为向上
		}
		row += dir
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := range rows {
		b.WriteString(rows[i].String())
	}
	return b.String()
}
