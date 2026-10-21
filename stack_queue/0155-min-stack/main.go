package minstack

// 155. 最小栈
// https://leetcode.cn/problems/min-stack/
// 设计一个支持 push、pop、top 操作，并能在常数时间内检索到最小元素的栈。
//
// 核心思路：栈的特点是「只在栈顶变化」，所以每个元素入栈时，
// 「从栈底到它为止的最小值」是固定不变的。给每个元素顺手记下这个最小值，
// 弹出时一起弹出，栈顶记录的最小值就始终是当前整个栈的最小值。

// MinStack 解法一：辅助最小值栈 ⭐ 最优方案
// stack 存元素，mins[i] 存 stack[0..i] 的最小值，两栈同步进出。
// 所有操作时间复杂度 O(1)，空间复杂度 O(n)。
type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.mins) == 0 || val < this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

// MinStackEnumeration 解法二：每次 GetMin 遍历整个栈（对照用）
// Push/Pop/Top 时间复杂度 O(1)，GetMin 时间复杂度 O(n)；空间复杂度 O(n)。
type MinStackEnumeration struct {
	stack []int
}

func ConstructorEnumeration() MinStackEnumeration {
	return MinStackEnumeration{}
}

func (this *MinStackEnumeration) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStackEnumeration) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStackEnumeration) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStackEnumeration) GetMin() int {
	m := this.stack[0]
	for _, v := range this.stack[1:] {
		m = min(m, v)
	}
	return m
}
