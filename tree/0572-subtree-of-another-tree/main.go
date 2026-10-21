package subtree

import (
	"strconv"
	"strings"

	"leetcode-go/template"
)

// 572. 另一棵树的子树
// https://leetcode.cn/problems/subtree-of-another-tree/
// 判断 subRoot 是否是 root 的某个节点为根的「完整子树」（结构和值都相同，且一直到叶子）。
//
// 核心思路：
//   - 朴素：对 root 的每个节点，都判断以它为根的树是否和 subRoot 相同。
//   - 序列化：把两棵树都按前序序列化（空节点也写出来），subRoot 是子树
//     ⇔ subRoot 的序列化串是 root 序列化串的子串。每个值前后都加分隔符，
//     防止 "12" 误匹配到 "2" 这类情况。

type TreeNode = template.TreeNode

// isSubtree 解法一：前序序列化 + 字符串匹配 ⭐ 最优方案
// 序列化格式：每个节点写成 ",值"，空节点写成 ",#"，开头的逗号保证值的边界。
// strings.Contains 对这种长串匹配在实践中接近线性（也可以换成 KMP 保证 O(m+n)）。
// 时间复杂度 O(m+n)（KMP 保证；Contains 平均情况），空间复杂度 O(m+n)。
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return strings.Contains(serializePre(root), serializePre(subRoot))
}

func serializePre(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString(",#")
			return
		}
		sb.WriteByte(',')
		sb.WriteString(strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// isSubtreeEnumeration 解法二：DFS 枚举每个节点 + 判断相同的树（对照用）
// 对 root 的每个节点 x，调用 sameTree(x, subRoot)；任一为真即可。
// 时间复杂度 O(m*n)，空间复杂度 O(h)。
func isSubtreeEnumeration(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	return sameTree(root, subRoot) ||
		isSubtreeEnumeration(root.Left, subRoot) ||
		isSubtreeEnumeration(root.Right, subRoot)
}

func sameTree(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}
