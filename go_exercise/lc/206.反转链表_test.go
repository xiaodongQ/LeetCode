package lc

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (69.82%)
 * Likes:    1094
 * Dislikes: 0
 * Total Accepted:    281.2K
 * Total Submissions: 402.7K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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
	27/27 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 52.17 % of golang submissions (2.6 MB)
	时间O(n)，空间O(1)
*/
// func reverseList(head *ListNode) *ListNode {
// 	var pre, tmp *ListNode
// 	var cur = head
// 	for cur != nil {
// 		tmp = cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = tmp
// 	}
// 	return pre
// }

/*
	递归
	27/27 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 17.39 % of golang submissions (2.9 MB)
	时间O(n)，空间O(n)
*/
func reverseList(head *ListNode) *ListNode {
	// 链表为空，或者到了最后一个节点
	if head == nil || head.Next == nil {
		return head
	}
	// 下一个节点已经反转
	// 原链表：a -> b -> c，当前a，a之后已经变为 a -> b <- c
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return p
}

// @lc code=end
