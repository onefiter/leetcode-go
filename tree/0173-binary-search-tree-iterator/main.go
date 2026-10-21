package bstiterator

import "leetcode-go/template"

// 173. 二叉搜索树迭代器
// https://leetcode.cn/problems/binary-search-tree-iterator/
// 实现一个按中序（从小到大）遍历 BST 的迭代器：Next() 返回下一个最小值，HasNext() 判断是否还有。
// 进阶要求：Next/HasNext 均摊 O(1) 时间，O(h) 内存。
//
// 核心思路：把「迭代中序遍历」拆开执行——栈里始终存着「从当前位置一路向左」的节点。
// 每次 Next 弹出栈顶（就是当前最小值），再把它右子树的左链压栈。
// 每个节点只进出栈一次，所以 n 次 Next 总共 O(n)，均摊 O(1)。

type TreeNode = template.TreeNode

// BSTIterator 解法一：受控递归（显式栈）⭐ 最优方案
// 时间复杂度：Next 均摊 O(1)，HasNext O(1)；空间复杂度 O(h)。
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

// pushLeft 把 node 及其一路向左的节点全部压栈。
func (it *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

func (it *BSTIterator) Next() int {
	node := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	it.pushLeft(node.Right) // 右子树的最左链就是接下来要访问的
	return node.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

// BSTIteratorEnumeration 解法二：预先中序展开成数组（对照用）
// 构造时一次性中序遍历存入数组，之后用下标读取。
// 时间复杂度：构造 O(n)，Next/HasNext O(1)；空间复杂度 O(n)。
type BSTIteratorEnumeration struct {
	vals []int
	idx  int
}

func ConstructorEnumeration(root *TreeNode) BSTIteratorEnumeration {
	it := BSTIteratorEnumeration{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		it.vals = append(it.vals, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return it
}

func (it *BSTIteratorEnumeration) Next() int {
	v := it.vals[it.idx]
	it.idx++
	return v
}

func (it *BSTIteratorEnumeration) HasNext() bool {
	return it.idx < len(it.vals)
}
