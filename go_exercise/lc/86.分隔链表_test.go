package lc

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (58.71%)
 * Likes:    226
 * Dislikes: 0
 * Total Accepted:    44.4K
 * Total Submissions: 75.6K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 *
 * 你应当保留两个分区中每个节点的初始相对位置。
 *
 * 示例:
 *
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
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
	利用两个哨兵分别指向大小两部分链表，最后拼接在一起，注意最后的元素的Next需要断开原指向
	166/166 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.4 MB)
	时间O(n)，空间O(1)
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	pre1, pre2 := dummy1, dummy2
	for head != nil {
		if head.Val < x {
			pre1.Next = head
			pre1 = head
		} else {
			pre2.Next = head
			pre2 = head
		}
		head = head.Next
	}
	// 拼接两个链表，修改最后节点指向(否则会保留原指向，导致有环)
	// 考虑仅有最后一个节点
	pre2.Next = nil
	pre1.Next = dummy2.Next

	return dummy1.Next
}

// @lc code=end
