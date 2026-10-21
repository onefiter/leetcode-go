package longestmountain

// 845. 数组中的最长山脉
// https://leetcode.cn/problems/longest-mountain-in-array/
// 「山脉」是长度 >= 3 的连续子数组，先严格递增到某个山顶，再严格递减。
// 返回最长山脉的长度，不存在返回 0。例如 [2,1,4,7,3,2,5] -> 5（[1,4,7,3,2]）。
//
// 核心思路：用指针 left 定位山脚，从它出发先走完严格上坡，再走完严格下坡，
// 上坡和下坡都非空才算一座山；然后从山的末尾继续找下一座（相邻两座山共享山脚）。

// longestMountain 解法一：双指针枚举山脚 ⭐ 最优方案
// 每个元素只被 right 扫过常数次。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func longestMountain(arr []int) int {
	n := len(arr)
	best := 0
	left := 0
	for left+2 < n {
		right := left
		// 上坡
		for right+1 < n && arr[right] < arr[right+1] {
			right++
		}
		if right == left {
			left++ // 没有上坡，换下一个起点
			continue
		}
		peak := right
		// 下坡
		for right+1 < n && arr[right] > arr[right+1] {
			right++
		}
		if right > peak {
			best = max(best, right-left+1)
			left = right // 下坡终点可作为下一座山的山脚
		} else {
			left = peak // 只有上坡没有下坡，从山顶(平台)继续
		}
	}
	return best
}

// longestMountainDP 解法二：前后缀递推
// up[i]：以 i 结尾的严格上坡长度（不含 i 本身的步数）；down[i]：从 i 开始的严格下坡步数。
// 当 up[i]>0 且 down[i]>0 时，i 可作山顶，长度为 up[i]+down[i]+1。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func longestMountainDP(arr []int) int {
	n := len(arr)
	up := make([]int, n)
	down := make([]int, n)
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			up[i] = up[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			down[i] = down[i+1] + 1
		}
	}
	best := 0
	for i := range n {
		if up[i] > 0 && down[i] > 0 {
			best = max(best, up[i]+down[i]+1)
		}
	}
	return best
}

// longestMountainEnumeration 解法三：枚举山顶向两边扩展（对照用）
// 对每个位置当作山顶，分别向左、向右走严格下降的步数。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func longestMountainEnumeration(arr []int) int {
	best := 0
	for i := 1; i+1 < len(arr); i++ {
		l, r := i, i
		for l > 0 && arr[l-1] < arr[l] {
			l--
		}
		for r+1 < len(arr) && arr[r] > arr[r+1] {
			r++
		}
		if l < i && r > i {
			best = max(best, r-l+1)
		}
	}
	return best
}
