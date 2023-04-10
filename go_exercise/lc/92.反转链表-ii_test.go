package lc

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (50.77%)
 * Likes:    429
 * Dislikes: 0
 * Total Accepted:    62.1K
 * Total Submissions: 122.2K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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
	先找到m位置(题目中起始节点算做第一个，即m=1)，要操作n-m+1个节点
	44/44 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.1 MB)
	时间O(n)，空间O(1)
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 操作的是第一个则前一个位置不对
	dummy := &ListNode{Next: head}

	opNode := head
	// 记录m节点前一个位置
	opPreNode := dummy
	for i := 0; i < m-1; i++ {
		opPreNode = opNode
		opNode = opNode.Next
	}

	// 用于反转
	curNode := opNode
	preNode := opPreNode
	for i := 0; i < n-m+1; i++ {
		tmp := curNode.Next
		curNode.Next = preNode

		preNode = curNode
		curNode = tmp
	}
	// m节点及其前一节点的指向需要改变
	opNode.Next = curNode
	opPreNode.Next = preNode

	return dummy.Next
}

// @lc code=end
