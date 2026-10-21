package serializebt

import (
	"strconv"
	"strings"

	"leetcode-go/template"
)

// 297. 二叉树的序列化与反序列化
// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
// 设计一个算法，把任意二叉树转换成字符串（序列化），并能从字符串还原出原树（反序列化）。
//
// 核心思路：只记录节点值不够，必须把「空孩子」也记下来，结构才唯一。
//   - 前序 DFS：根,左,右 依次输出，空节点写 "#"；反序列化时按同样顺序递归消费 token。
//   - 层序 BFS：按层输出，空节点写 "#"；反序列化时用队列依次给每个节点挂左右孩子。
// token 之间用 "," 分隔，负数、多位数都能正确切分。

type TreeNode = template.TreeNode

// Codec 解法一：前序 DFS ⭐ 最优方案
// serialize：前序遍历，空节点输出 "#"，用逗号连接。
// deserialize：按逗号切分成 token，递归地「取一个 token 建根，再建左子树、右子树」。
// 时间复杂度 O(n)，空间复杂度 O(n)（字符串 + 递归栈）。
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if sb.Len() > 0 {
			sb.WriteByte(',')
		}
		if node == nil {
			sb.WriteByte('#')
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		tok := tokens[i]
		i++
		if tok == "#" {
			return nil
		}
		v, _ := strconv.Atoi(tok)
		node := &TreeNode{Val: v}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

// CodecBFS 解法二：层序 BFS
// serialize：队列层序遍历，空节点也入结果（写 "#"），但不再展开其孩子。
// deserialize：第一个 token 是根；之后每出队一个节点，就依次读两个 token 作为它的左右孩子。
// 时间复杂度 O(n)，空间复杂度 O(n)。
type CodecBFS struct{}

func ConstructorBFS() CodecBFS {
	return CodecBFS{}
}

func (this *CodecBFS) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	var parts []string
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			parts = append(parts, "#")
			continue
		}
		parts = append(parts, strconv.Itoa(node.Val))
		q = append(q, node.Left, node.Right)
	}
	return strings.Join(parts, ",")
}

func (this *CodecBFS) deserialize(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	if tokens[0] == "#" {
		return nil
	}
	newNode := func(tok string) *TreeNode {
		if tok == "#" {
			return nil
		}
		v, _ := strconv.Atoi(tok)
		return &TreeNode{Val: v}
	}
	root := newNode(tokens[0])
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 && i < len(tokens) {
		node := q[0]
		q = q[1:]
		node.Left = newNode(tokens[i])
		i++
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if i < len(tokens) {
			node.Right = newNode(tokens[i])
			i++
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
