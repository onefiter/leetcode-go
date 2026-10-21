package basiccalculator

// 224. 基本计算器
// https://leetcode.cn/problems/basic-calculator/
// 给定一个只包含数字、'+'、'-'、'('、')' 和空格的合法表达式字符串 s，计算它的值。
// 注意 '-' 可以作为一元运算符（如 "-(2+3)"），不允许使用 eval 之类的内置函数。
//
// 核心思路：表达式里只有加减，那么每个数字最终都是以「+」或「-」的身份累加到结果里。
// 一个数字的最终符号 = 它前面的符号 × 外层所有括号前的符号。
// 用一个栈记录「每层括号外的符号」，就能一次遍历算出答案。

// calculate 解法一：符号栈 ⭐ 最优方案
// ops 栈顶表示当前括号层的整体符号（初始为 +1）。
//   - 遇到 '+'：sign = 栈顶；遇到 '-'：sign = -栈顶；
//   - 遇到 '('：把当前 sign 压栈，作为括号内的整体符号；遇到 ')'：弹栈；
//   - 遇到数字：读完整个数，res += sign * num。
//
// 时间复杂度 O(n)，空间复杂度 O(n)。
func calculate(s string) int {
	ops := []int{1}
	sign, res := 1, 0
	for i := 0; i < len(s); {
		switch c := s[i]; {
		case c == ' ':
			i++
		case c == '+':
			sign = ops[len(ops)-1]
			i++
		case c == '-':
			sign = -ops[len(ops)-1]
			i++
		case c == '(':
			ops = append(ops, sign)
			i++
		case c == ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += sign * num
		}
	}
	return res
}

// calculateStack 解法二：结果栈（遇括号保存现场）
// 维护当前层的 res 和 sign。遇到 '(' 就把外层的 res、sign 压栈，括号内从 0 重新算；
// 遇到 ')' 时，括号内结果 res 乘以压栈时的 sign，再加回外层的 res。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func calculateStack(s string) int {
	stack := []int{}
	res, sign := 0, 1
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			i-- // for 循环会再 ++
			res += sign * num
		case c == '+':
			sign = 1
		case c == '-':
			sign = -1
		case c == '(':
			stack = append(stack, res, sign)
			res, sign = 0, 1
		case c == ')':
			prevSign, prevRes := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res = prevRes + prevSign*res
		}
	}
	return res
}

// calculateEnumeration 解法三：递归下降（对照用）
// 按文法 expr := term (('+'|'-') term)*，term := 数字 | '(' expr ')' | '-' term 直接递归求值，
// 思路最直观，但递归深度与括号嵌套深度成正比。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func calculateEnumeration(s string) int {
	p := &parser{s: s}
	return p.expr()
}

type parser struct {
	s   string
	pos int
}

func (p *parser) skip() {
	for p.pos < len(p.s) && p.s[p.pos] == ' ' {
		p.pos++
	}
}

func (p *parser) expr() int {
	res := p.term()
	for {
		p.skip()
		if p.pos >= len(p.s) || (p.s[p.pos] != '+' && p.s[p.pos] != '-') {
			return res
		}
		op := p.s[p.pos]
		p.pos++
		if op == '+' {
			res += p.term()
		} else {
			res -= p.term()
		}
	}
}

func (p *parser) term() int {
	p.skip()
	switch p.s[p.pos] {
	case '-':
		p.pos++
		return -p.term()
	case '(':
		p.pos++
		v := p.expr()
		p.skip()
		p.pos++ // 跳过 ')'
		return v
	}
	num := 0
	for ; p.pos < len(p.s) && p.s[p.pos] >= '0' && p.s[p.pos] <= '9'; p.pos++ {
		num = num*10 + int(p.s[p.pos]-'0')
	}
	return num
}
