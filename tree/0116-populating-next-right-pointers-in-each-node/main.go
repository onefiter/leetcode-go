package populatenext

// 116. 填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
// 给一棵「完美二叉树」（所有叶子在同一层，每个父节点都有两个子节点），
// 让每个节点的 Next 指向同一层右边的节点，最右边的指向 nil。
//
// 核心思路：上一层的 Next 已经连好后，它就是一条「横向链表」，
// 沿着它就能遍历上一层所有节点，顺手把下一层连起来：
//   - 同一父节点：Left.Next = Right；
//   - 相邻父节点：Right.Next = 父.Next.Left。
// 这样不需要队列，只用 O(1) 额外空间。

// Node 带 Next 指针的二叉树节点。
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 解法一：利用已建立的 Next 指针逐层连接 ⭐ 最优方案
// leftmost 指向每层最左节点，沿 Next 遍历该层并连接下一层。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func connect(root *Node) *Node {
	for leftmost := root; leftmost != nil && leftmost.Left != nil; leftmost = leftmost.Left {
		for cur := leftmost; cur != nil; cur = cur.Next {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
		}
	}
	return root
}

// connectRecursive 解法二：递归，连接同父与跨父的两类相邻节点
// 时间复杂度 O(n)，空间复杂度 O(log n)（完美二叉树高度）。
func connectRecursive(root *Node) *Node {
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil || n.Left == nil {
			return
		}
		n.Left.Next = n.Right
		if n.Next != nil {
			n.Right.Next = n.Next.Left
		}
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return root
}

// connectBFS 解法三：BFS 层序遍历（对照用）
// 按层处理，同一层相邻节点依次相连。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func connectBFS(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		next := []*Node{}
		for i, n := range queue {
			if i+1 < len(queue) {
				n.Next = queue[i+1]
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		queue = next
	}
	return root
}
