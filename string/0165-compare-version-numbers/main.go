package compareversion

import (
	"strconv"
	"strings"
)

// 165. 比较版本号
// https://leetcode.cn/problems/compare-version-numbers/
// 给你两个版本号 version1 和 version2，请比较它们。
// 版本号由一个或多个修订号组成，修订号之间以 '.' 连接，每个修订号是整数。
// 比较时忽略修订号的前导零（如 "1.0" 等于 "1.00"），
// 若某个版本号缺少对应修订号，则视为 0。
// version1 > version2 返回 1，version1 < version2 返回 -1，否则返回 0。
// 例如 "1.01" 与 "1.001" -> 0；"1.0" 与 "1.0.0" -> 0；"0.1" 与 "1.1" -> -1。
//
// 一句话理解题意：按 '.' 分段，把每段当整数逐段比较，缺失段补 0。
//
// 核心思路：按点分割成修订号列表，逐段把字符串转成整数比较；
// 长度不同的用 0 补齐，从前往后第一个不相等的段决定大小。

// compareVersion 【解法1】分割 + 逐段整数比较 ⭐ 最优实用方案
// 思路：用 strings.Split 按 '.' 切分，遍历较长的段数，
// 越界的一方取 0，Atoi 转整数后比较（Atoi 天然忽略前导零）。
// 时间复杂度 O(n+m)：n、m 为两版本号长度。
// 空间复杂度 O(n+m)：存放分割后的切片。
func compareVersion(version1, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i := 0; i < len(v1) || i < len(v2); i++ {
		x, y := 0, 0
		if i < len(v1) {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

// compareVersionScan 【解法2】双指针不分割（对照用，省去切片分配）
// 思路：两个指针各自在原字符串上逐段解析出当前修订号的整数值，
// 再比较，直到两边都扫描结束。避免 Split 产生中间切片。
// 时间复杂度 O(n+m)，空间复杂度 O(1)。
func compareVersionScan(version1, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		x := 0
		for i < len(version1) && version1[i] != '.' {
			x = x*10 + int(version1[i]-'0')
			i++
		}
		i++ // 跳过 '.'
		y := 0
		for j < len(version2) && version2[j] != '.' {
			y = y*10 + int(version2[j]-'0')
			j++
		}
		j++ // 跳过 '.'
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
