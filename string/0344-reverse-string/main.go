package reversestring

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string/
// 编写一个函数，将输入的字符串反转。输入字符串以字符数组 s 的形式给出。
// 必须原地修改输入数组、使用 O(1) 额外空间解决。
// 例如 s=['h','e','l','l','o'] -> ['o','l','l','e','h']。
//
// 一句话理解题意：把字节数组头尾对调，原地翻转。
//
// 核心思路：双指针分别指向首尾，交换后向中间靠拢，直到相遇。

// reverseString 【解法1】双指针原地交换 ⭐ 最优方案
// 思路：左指针 i 从头、右指针 j 从尾，两两交换后 i++、j--，直到 i>=j。
// 时间复杂度 O(n)：交换 n/2 次。
// 空间复杂度 O(1)：原地操作，只用常数变量。
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
