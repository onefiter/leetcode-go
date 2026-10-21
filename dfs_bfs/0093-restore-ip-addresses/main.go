package restoreip

import "strings"

// 93. 复原 IP 地址
// https://leetcode.cn/problems/restore-ip-addresses/
// 给一个只含数字的字符串 s，在其中插入三个 '.' 形成所有可能的有效 IP 地址。
// 有效 IP 由 4 段 0~255 的整数组成，且除 "0" 本身外不能有前导零，顺序任意。
//
// 核心思路：回溯切分，每层截取 1~3 位作为一段，校验合法后进入下一段。
// 剪枝：剩余字符数必须落在 [剩余段数, 3*剩余段数] 区间内，否则直接返回。

// valid 判断 seg 是否为合法的 IP 段（0~255，无前导零）。
func valid(seg string) bool {
	if len(seg) == 0 || len(seg) > 3 || (len(seg) > 1 && seg[0] == '0') {
		return false
	}
	v := 0
	for i := 0; i < len(seg); i++ {
		v = v*10 + int(seg[i]-'0')
	}
	return v <= 255
}

// restoreIpAddresses 解法一：回溯 + 长度剪枝 ⭐ 最优方案
// 时间复杂度 O(3^4 * |s|)，空间复杂度 O(1)（段数固定为 4，不计结果）。
func restoreIpAddresses(s string) []string {
	res := []string{}
	segs := make([]string, 0, 4)
	var dfs func(start int)
	dfs = func(start int) {
		left, need := len(s)-start, 4-len(segs)
		if left < need || left > need*3 {
			return // 剩余字符不够或太多
		}
		if need == 0 {
			res = append(res, strings.Join(segs, "."))
			return
		}
		for l := 1; l <= 3 && start+l <= len(s); l++ {
			seg := s[start : start+l]
			if !valid(seg) {
				break // 前导零或超过 255，更长只会更不合法
			}
			segs = append(segs, seg)
			dfs(start + l)
			segs = segs[:len(segs)-1]
		}
	}
	dfs(0)
	return res
}

// restoreIpAddressesEnumeration 解法二：三重循环枚举三个点的位置（对照用）
// 每段长度 1~3，直接枚举前三段长度 a、b、c，第四段长度由剩余决定。
// 时间复杂度 O(3^3 * |s|)，空间复杂度 O(1)（不计结果）。
func restoreIpAddressesEnumeration(s string) []string {
	res := []string{}
	n := len(s)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				d := n - a - b - c
				if d < 1 || d > 3 {
					continue
				}
				s1, s2 := s[:a], s[a:a+b]
				s3, s4 := s[a+b:a+b+c], s[a+b+c:]
				if valid(s1) && valid(s2) && valid(s3) && valid(s4) {
					res = append(res, s1+"."+s2+"."+s3+"."+s4)
				}
			}
		}
	}
	return res
}
