package lc

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (66.21%)
 * Likes:    557
 * Dislikes: 0
 * Total Accepted:    127.1K
 * Total Submissions: 192K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
	递归
	从链表头开始，每次交换一对，本节点和其后一节点
	55/55 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 20 % of golang submissions (2.1 MB)
	时间O(n)，空间O(n)
*/
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	first := head
// 	second := head.Next

// 	first.Next = swapPairs(second.Next)
// 	second.Next = first

// 	return second
// }

/*
失败：
	循环
	没考虑到后续迭代完之后，前面一对的最后节点指向应该改变
	新建一个preNode进行操作，借助哨兵(或叫哑节点)更简便
*/
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	res := head.Next
// 	for head != nil && head.Next != nil {
// 		first := head
// 		second := head.Next

// 		first.Next = second.Next
// 		second.Next = first

// 		// 只考虑了下一次跳2个节点迭代，
// 		// 但没考虑到之前最后节点的指向也应该改变
// 		head = first.Next
// 	}

// 	return res
// }

/*
	迭代
	55/55 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 20 % of golang submissions (2.1 MB)
	时间O(n)，空间O(1)
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	preNode := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next
		// 前一节点指向改变
		preNode.Next = second
		// 交换
		first.Next = second.Next
		second.Next = first

		preNode = first
		head = first.Next
	}

	return dummy.Next
}

// @lc code=end
