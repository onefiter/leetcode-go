package reversewords

import "strings"

// 151. 反转字符串中的单词
// https://leetcode.cn/problems/reverse-words-in-a-string/
// 给你一个字符串 s，请你反转字符串中「单词」的顺序。
// 单词是由非空格字符组成的字符串，单词间以至少一个空格分隔。
// 返回的字符串中，单词间应只有一个空格，且不含前导或尾随空格。
// 例如 s="the sky is blue" -> "blue is sky the"；
// s="  hello world  " -> "world hello"；s="a good   example" -> "example good a"。
//
// 一句话理解题意：把多余空格清掉，按空格切成单词，再倒序拼回去。
//
// 核心思路：先切分出所有非空单词，反转单词顺序，用单个空格连接。

// reverseWords 【解法1】切分 + 反转 ⭐ 最优实用方案
// 思路：strings.Fields 会按任意数量空白切分并自动忽略多余空格，
// 得到单词切片后原地反转，再用单个空格 Join。
// 时间复杂度 O(n)：切分与反转都是线性。
// 空间复杂度 O(n)：存放单词切片与结果。
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// reverseWordsManual 【解法2】手动扫描（对照用，不依赖 Fields/Join）
// 思路：从右往左扫描字符串，跳过空格，定位每个单词的 [start, end)，
// 依次把单词追加到结果 builder 中，单词之间补一个空格。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reverseWordsManual(s string) string {
	var b strings.Builder
	i := len(s) - 1
	for i >= 0 {
		// 跳过尾部空格
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i < 0 {
			break
		}
		// 定位单词 [j+1, i]
		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(s[j+1 : i+1])
		i = j
	}
	return b.String()
}
