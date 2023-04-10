package lc

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (39.03%)
 * Likes:    897
 * Dislikes: 0
 * Total Accepted:    199.7K
 * Total Submissions: 511.7K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
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
	~~快慢指针~~ 应该是双指针，
	快指针先走n步，而后两个指针同步走，走到尾部则慢指针为倒数第n个
	208/208 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 64.29 % of golang submissions (2.2 MB)
*/
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	sentinel := &ListNode{Next: head}
// 	slow, fast := head, head
// 	preNode := sentinel
// 	for i := 0; i < n; i++ {
// 		fast = fast.Next
// 	}
// 	for fast != nil {
// 		preNode = slow
// 		slow = slow.Next
// 		fast = fast.Next
// 	}
// 	// fast到底了说明slow已经倒数第n个
// 	preNode.Next = slow.Next

// 	return sentinel.Next
// }

// 和上面一样，fast多走一步，则最后slow停留在倒数第n个的前一个节点
// 时间O(n)，空间O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
	// 都从哨兵开始走
	slow, fast := sentinel, sentinel
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// fast到底了说明slow已经倒数第n个的前一个
	slow.Next = slow.Next.Next

	return sentinel.Next
}

// @lc code=end
