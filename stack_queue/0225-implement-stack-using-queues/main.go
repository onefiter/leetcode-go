package stackusingqueues

// 225. 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues/
// 只用队列的标准操作（队尾 push、队头 peek/pop、size、empty）实现一个后进先出的栈，
// 支持 push、pop、top、empty。
//
// 核心思路：队列是先进先出，栈是后进先出。每次 push 新元素后，
// 把它前面的所有元素依次出队再入队，新元素就被「转」到了队头，
// 于是队头永远是最后入栈的元素，pop/top 直接操作队头即可。
// 这里用切片模拟队列，只使用 append（队尾入）和 q[0]/q[1:]（队头出）。

// MyStack 解法一：单队列，push 时旋转 ⭐ 最优方案
// Push 时间复杂度 O(n)，Pop/Top/Empty 时间复杂度 O(1)；空间复杂度 O(n)。
type MyStack struct {
	q []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	// 把新元素前面的 n 个元素依次移到队尾。
	for i := 0; i < n; i++ {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}
}

func (this *MyStack) Pop() int {
	x := this.q[0]
	this.q = this.q[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

// MyStackTwoQueues 解法二：双队列
// 新元素先入辅助队列 q2，再把主队列 q1 的元素全部搬到 q2 后面，最后交换 q1、q2。
// 这样 q1 的队头始终是栈顶。
// Push 时间复杂度 O(n)，Pop/Top/Empty 时间复杂度 O(1)；空间复杂度 O(n)。
type MyStackTwoQueues struct {
	q1, q2 []int
}

func ConstructorTwoQueues() MyStackTwoQueues {
	return MyStackTwoQueues{}
}

func (this *MyStackTwoQueues) Push(x int) {
	this.q2 = append(this.q2, x)
	for len(this.q1) > 0 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	this.q1, this.q2 = this.q2, nil
}

func (this *MyStackTwoQueues) Pop() int {
	x := this.q1[0]
	this.q1 = this.q1[1:]
	return x
}

func (this *MyStackTwoQueues) Top() int {
	return this.q1[0]
}

func (this *MyStackTwoQueues) Empty() bool {
	return len(this.q1) == 0
}
