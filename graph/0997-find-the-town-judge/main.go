package townjudge

// 997. 找到小镇的法官
// https://leetcode.cn/problems/find-the-town-judge/
// 小镇有 n 个人（编号 1..n），trust[i] = (a, b) 表示 a 信任 b。
// 法官不信任任何人，且其他所有人都信任法官，最多只有一个法官。
// 存在法官则返回其编号，否则返回 -1。
//
// 核心思路：把信任关系看成有向图，法官就是「出度 0、入度 n-1」的点。
// 用一个数组记录 入度 - 出度，恰好等于 n-1 的点即为法官。

// findJudge 解法一：度数差计数 ⭐ 最优方案
// 出度 > 0 会让差值小于 n-1，所以只看「入度 - 出度 == n-1」即可同时满足两个条件。
// 时间复杂度 O(n + m)，空间复杂度 O(n)。
func findJudge(n int, trust [][]int) int {
	score := make([]int, n+1)
	for _, t := range trust {
		score[t[0]]--
		score[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if score[i] == n-1 {
			return i
		}
	}
	return -1
}

// findJudgeBruteForce 解法二：邻接矩阵暴力检查（对照用）
// 对每个候选人 i，检查他不信任任何人、且其他人都信任他。
// 时间复杂度 O(n² + m)，空间复杂度 O(n²)。
func findJudgeBruteForce(n int, trust [][]int) int {
	trusts := make([][]bool, n+1)
	for i := range trusts {
		trusts[i] = make([]bool, n+1)
	}
	for _, t := range trust {
		trusts[t[0]][t[1]] = true
	}
	for i := 1; i <= n; i++ {
		ok := true
		for j := 1; j <= n && ok; j++ {
			if j == i {
				continue
			}
			if trusts[i][j] || !trusts[j][i] {
				ok = false
			}
		}
		if ok {
			return i
		}
	}
	return -1
}
