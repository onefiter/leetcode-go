package serializebst

import (
	"math"
	"strconv"
	"strings"

	"leetcode-go/template"
)

// 449. 序列化和反序列化二叉搜索树
// https://leetcode.cn/problems/serialize-and-deserialize-bst/
// 设计算法把一棵二叉搜索树编码成尽量紧凑的字符串，并能还原出原树。
//
// 核心思路：BST 的结构完全由「前序序列」决定，不需要像普通二叉树那样记录空节点。
// 反序列化时按前序依次取值，并给每个位置维护一个合法区间 (lower, upper)：
// 当前值落在区间内，就是这个位置的节点；否则说明该位置为空，交给上层的其他位置。

type TreeNode = template.TreeNode

// Codec 解法一：前序值 + 上下界还原 ⭐ 最优方案
// serialize：前序遍历只输出节点值，用逗号连接；空树为 ""。
// deserialize：build(lower, upper) 若下一个值在区间内则建节点，
// 左子树区间 (lower, val)，右子树区间 (val, upper)；每个值只被消费一次。
// 时间复杂度 O(n)，空间复杂度 O(n)（字符串 + 递归栈 O(h)）。
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var parts []string
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		parts = append(parts, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return strings.Join(parts, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, ",")
	vals := make([]int, len(tokens))
	for i, tok := range tokens {
		vals[i], _ = strconv.Atoi(tok)
	}
	i := 0
	var build func(lower, upper int) *TreeNode
	build = func(lower, upper int) *TreeNode {
		if i == len(vals) || vals[i] < lower || vals[i] > upper {
			return nil
		}
		v := vals[i]
		i++
		node := &TreeNode{Val: v}
		node.Left = build(lower, v)
		node.Right = build(v, upper)
		return node
	}
	return build(math.MinInt, math.MaxInt)
}

// CodecInsert 解法二：前序值 + 逐个插入 BST（对照用）
// serialize 与解法一相同；deserialize 把前序值依次插入空 BST，
// 前序中根先于子孙出现，插入后结构与原树一致。
// 时间复杂度 O(n·h)，最坏 O(n²)（链状树），空间复杂度 O(n)。
type CodecInsert struct{}

func ConstructorInsert() CodecInsert {
	return CodecInsert{}
}

func (this *CodecInsert) serialize(root *TreeNode) string {
	c := Constructor()
	return c.serialize(root)
}

func (this *CodecInsert) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var root *TreeNode
	for _, tok := range strings.Split(data, ",") {
		v, _ := strconv.Atoi(tok)
		node := &TreeNode{Val: v}
		if root == nil {
			root = node
			continue
		}
		cur := root
		for {
			if v < cur.Val {
				if cur.Left == nil {
					cur.Left = node
					break
				}
				cur = cur.Left
			} else {
				if cur.Right == nil {
					cur.Right = node
					break
				}
				cur = cur.Right
			}
		}
	}
	return root
}
