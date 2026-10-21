package findthewinner

// 1823. 找出游戏的获胜者
// https://leetcode.cn/problems/find-the-winner-of-the-circular-game/
// n 个朋友编号 1..n 围成一圈，从 1 号开始顺时针数 k 个人（含自己），数到的人出局，
// 再从出局者的下一位继续数，直到只剩一人，返回他的编号。
// 例如 n=5, k=2 -> 3。
//
// 核心思路（约瑟夫环）：设 f(i) 为 i 个人（编号 0..i-1）时胜者的编号。
// 第一轮淘汰编号 (k-1)%i 的人后，从下一个人重新编号，问题变成 i-1 个人；
// 新编号 x 对应原编号 (x+k)%i，所以 f(i) = (f(i-1)+k) % i，f(1)=0。答案是 f(n)+1。

// findTheWinner 解法一：数学递推 ⭐ 最优方案
// 从 1 个人推到 n 个人，只用一个变量。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findTheWinner(n int, k int) int {
	winner := 0 // f(1) = 0
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1
}

// findTheWinnerEnumeration 解法二：队列模拟（对照用）
// 把 1..n 放进队列，每轮把队头的 k-1 个人依次挪到队尾，再弹出队头（出局）。
// 时间复杂度 O(nk)，空间复杂度 O(n)。
func findTheWinnerEnumeration(n int, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	for len(q) > 1 {
		for j := 0; j < k-1; j++ {
			q = append(q[1:], q[0])
		}
		q = q[1:] // 第 k 个人出局
	}
	return q[0]
}
