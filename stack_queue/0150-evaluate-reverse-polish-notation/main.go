package evalrpn

import "strconv"

// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 给定逆波兰表示法（后缀表达式）的 token 数组，求表达式的值。
// 运算符只有 + - * /，整数除法向零截断。例如 ["2","1","+","3","*"] -> 9。
//
// 核心思路：后缀表达式天然适合用栈求值——遇到数字就入栈；
// 遇到运算符就弹出两个数（先弹出的是右操作数），计算后把结果压回栈。
// 最后栈里剩下的唯一元素就是答案。

// evalRPN 解法一：栈 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(n)。
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, tk := range tokens {
		switch tk {
		case "+", "-", "*", "/":
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, apply(a, b, tk))
		default:
			v, _ := strconv.Atoi(tk)
			stack = append(stack, v)
		}
	}
	return stack[0]
}

// evalRPNArray 解法二：数组模拟栈
// 后缀表达式里栈的最大深度不超过 (n+1)/2，预分配定长数组并用下标 top 维护栈顶，
// 省去 append 的切片操作，思路与解法一相同。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func evalRPNArray(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	top := -1
	for _, tk := range tokens {
		if tk == "+" || tk == "-" || tk == "*" || tk == "/" {
			top--
			stack[top] = apply(stack[top], stack[top+1], tk)
		} else {
			top++
			stack[top], _ = strconv.Atoi(tk)
		}
	}
	return stack[0]
}

// apply 计算 a op b，Go 的整数除法本身就是向零截断。
func apply(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
