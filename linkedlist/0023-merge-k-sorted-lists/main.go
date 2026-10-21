package mergeksortedlists

import (
	"container/heap"

	"leetcode-go/template"
)

// 23. 合并 K 个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/
// 给 k 个各自升序的链表，把它们合并成一个升序链表返回。
// 例如 [[1,4,5],[1,3,4],[2,6]] -> [1,1,2,3,4,4,5,6]。
//
// 核心思路：
//   - 最小堆：每条链表的「当前头」放进堆，每次弹出最小的接到结果后面，再把它的 next 入堆；
//   - 分治：像归并排序一样两两合并，log k 轮，每轮总共处理 N 个节点。
// 两者都是 O(N log k)，N 为节点总数。

type ListNode = template.ListNode

// mergeKLists 解法一：分治两两合并 ⭐ 最优方案
// 把 lists 对半拆分，递归合并左右两半，再用「合并两个有序链表」合起来。
// 递归深度 log k，每层合并的节点总数为 N。
// 时间复杂度 O(N log k)，空间复杂度 O(log k)（递归栈）。
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeTwo(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

// mergeTwo 合并两个升序链表（复用原节点）。
func mergeTwo(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			cur.Next, a = a, a.Next
		} else {
			cur.Next, b = b, b.Next
		}
		cur = cur.Next
	}
	if a != nil {
		cur.Next = a
	} else {
		cur.Next = b
	}
	return dummy.Next
}

// nodeHeap 按 Val 排序的最小堆，实现 heap.Interface。
type nodeHeap []*ListNode

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// mergeKListsHeap 解法二：最小堆 ⭐ 最优方案（与分治同复杂度）
// 堆里始终最多 k 个节点（每条链表的当前头），弹出最小值接到结果尾部，
// 若它还有 next 就把 next 入堆。
// 时间复杂度 O(N log k)，空间复杂度 O(k)。
func mergeKListsHeap(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			*h = append(*h, l)
		}
	}
	heap.Init(h)
	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		cur.Next = node
		cur = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	cur.Next = nil
	return dummy.Next
}

// mergeKListsEnumeration 解法三：逐一合并（对照用）
// 用结果链表依次和每条链表合并。第 i 次合并时结果长度约为 i*n，
// 总代价约 n*(1+2+...+k) = O(k²n) = O(kN)。
// 时间复杂度 O(kN)，空间复杂度 O(1)。
func mergeKListsEnumeration(lists []*ListNode) *ListNode {
	var res *ListNode
	for _, l := range lists {
		res = mergeTwo(res, l)
	}
	return res
}
