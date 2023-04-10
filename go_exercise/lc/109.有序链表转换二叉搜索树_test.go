package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (73.03%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    34.5K
 * Total Submissions: 47.2K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 *
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 *
 *
 */
func TestSortedListToBST(t *testing.T) {
	l := InitLinkInfo([]int{-10, -3, 0, 5, 9})
	log.Printf("res:%v", sortedListToBST(l))
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
	中间节点作为根节点，分前后两部分进行递归，
	前一部分的链表需要从中间节点的左侧断开(中间节点使用后其之前的节点不能指向它了)

中间取左侧(左侧多了一些nil节点，取右侧更简洁)：
	32/32 cases passed (4 ms)
	Your runtime beats 99.26 % of golang submissions
	Your memory usage beats 100 % of golang submissions (6.2 MB)
中间取右侧：
	32/32 cases passed (4 ms)
	Your runtime beats 99.26 % of golang submissions
	Your memory usage beats 100 % of golang submissions (6.1 MB)
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMiddle(head)
	tree := &TreeNode{Val: mid.Val}
	// 需要考虑只有一个节点的情况
	// 偶数取右边，这里没有问题，
	/*
		若取左边不能直接return，
		如果是两个节点取左边，则该节点还要断开和后续的连接
	*/
	// 取右侧中间
	if head == mid {
		return tree
	}

	// 取左侧中间(left不用继续，虽然tree的left置nil了，但原链表head还是指向下一个节点，并没有断开)
	// if head == mid {
	// 	head = nil
	// }
	tree.Left = sortedListToBST(head)
	tree.Right = sortedListToBST(mid.Next)

	return tree
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	var pre *ListNode
	// 此处偶数取左边
	// for fast.Next != nil && fast.Next.Next != nil {
	// 	pre = slow
	// 	slow = slow.Next
	// 	fast = fast.Next.Next
	// }
	// 此处偶数取右边
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 从中间节点左侧断开
	if pre != nil {
		pre.Next = nil
	}
	return slow
}

// @lc code=end
