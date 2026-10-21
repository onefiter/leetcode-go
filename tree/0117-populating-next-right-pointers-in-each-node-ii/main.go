package populatenextii

// 117. 填充每个节点的下一个右侧节点指针 II
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
// 和 116 一样让每个节点的 Next 指向同层右侧节点，但树不再是完美二叉树，
// 子节点可能缺失，因此不能简单用 父.Next.Left 连接。
//
// 核心思路：把每一层看成一条链表。遍历当前层（沿已建立的 Next）时，
// 用一个「哑节点」dummy 作为下一层链表的头，tail 指向下一层链表的尾，
// 遇到子节点就挂到 tail 后面。当前层走完后，dummy.Next 就是下一层的最左节点。

// Node 带 Next 指针的二叉树节点。
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// connect 解法一：哑节点逐层串联 ⭐ 最优方案
// 外层循环按层推进，内层沿 Next 遍历当前层，把子节点依次接到下一层链表尾部。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func connect(root *Node) *Node {
	for head := root; head != nil; {
		dummy := &Node{}
		tail := dummy
		for cur := head; cur != nil; cur = cur.Next {
			if cur.Left != nil {
				tail.Next = cur.Left
				tail = tail.Next
			}
			if cur.Right != nil {
				tail.Next = cur.Right
				tail = tail.Next
			}
		}
		head = dummy.Next // 下一层的最左节点
	}
	return root
}

// connectBFS 解法二：BFS 层序遍历（对照用）
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
