package copyrandomlist

// 138. 随机链表的复制
// https://leetcode.cn/problems/copy-list-with-random-pointer/
// 链表每个节点除了 Next，还有一个可能指向任意节点或 nil 的 Random 指针。
// 要求深拷贝：返回一条全新的链表，新节点的 Next / Random 都只能指向新节点。
//
// 核心思路：难点在于复制 Random 时，目标节点的「副本」可能还没创建。
//   - 哈希表：先建好「原节点 -> 新节点」的映射，再统一连指针；
//   - 节点交织：把每个副本插到原节点后面（A -> A' -> B -> B'），
//     这样 A.Random 的副本就是 A.Random.Next，不需要哈希表，最后再把两条链拆开。

// Node 带随机指针的链表节点。
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 解法一：节点交织 ⭐ 最优方案
// 三步：① 每个节点后插入它的副本；② 副本.Random = 原.Random.Next；
// ③ 拆分，恢复原链表并串起副本链表。
// 时间复杂度 O(n)，空间复杂度 O(1)（不计结果链表）。
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// ① 交织：A -> A' -> B -> B' -> ...
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	// ② 设置副本的 Random。
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	// ③ 拆分两条链表。
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		cp := cur.Next
		cur.Next = cp.Next
		if cp.Next != nil {
			cp.Next = cp.Next.Next
		}
	}
	return newHead
}

// copyRandomListEnumeration 解法二：哈希表映射（对照用）
// 第一遍为每个原节点创建副本并记录映射；第二遍通过映射连好 Next 和 Random。
// map 中查不到的 key（包括 nil）返回零值 nil，正好处理指针为空的情况。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func copyRandomListEnumeration(head *Node) *Node {
	m := map[*Node]*Node{}
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
	}
	return m[head]
}
