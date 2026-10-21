package queueusingstacks

// 232. 用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks/
// 只用栈的标准操作（栈顶 push/peek/pop、size、empty）实现一个先进先出的队列，
// 支持 push、pop、peek、empty。
//
// 核心思路：把元素从一个栈倒进另一个栈，顺序就反过来了。
// 用 in 栈负责接收新元素，out 栈负责出队；out 为空时才把 in 整个倒进 out，
// 此时 out 的栈顶就是最早进来的元素。每个元素最多被倒一次，所以均摊 O(1)。

// MyQueue 解法一：双栈（惰性倒栈）⭐ 最优方案
// Push/Empty 时间复杂度 O(1)，Pop/Peek 均摊 O(1)；空间复杂度 O(n)。
type MyQueue struct {
	in, out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	x := this.Peek()
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

// MyQueueEnumeration 解法二：每次 push 都整体倒栈（对照用）
// 始终让 s 的栈顶是队头：push 时先把 s 全倒到 tmp，压入新元素，再倒回来。
// Push 时间复杂度 O(n)，Pop/Peek/Empty 时间复杂度 O(1)；空间复杂度 O(n)。
type MyQueueEnumeration struct {
	s, tmp []int
}

func ConstructorEnumeration() MyQueueEnumeration {
	return MyQueueEnumeration{}
}

func (this *MyQueueEnumeration) Push(x int) {
	for len(this.s) > 0 {
		this.tmp = append(this.tmp, this.s[len(this.s)-1])
		this.s = this.s[:len(this.s)-1]
	}
	this.s = append(this.s, x)
	for len(this.tmp) > 0 {
		this.s = append(this.s, this.tmp[len(this.tmp)-1])
		this.tmp = this.tmp[:len(this.tmp)-1]
	}
}

func (this *MyQueueEnumeration) Pop() int {
	x := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return x
}

func (this *MyQueueEnumeration) Peek() int {
	return this.s[len(this.s)-1]
}

func (this *MyQueueEnumeration) Empty() bool {
	return len(this.s) == 0
}
