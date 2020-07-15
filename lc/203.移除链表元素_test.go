package lc

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 *
 * https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (45.83%)
 * Likes:    408
 * Dislikes: 0
 * Total Accepted:    88.7K
 * Total Submissions: 193.4K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * 删除链表中等于给定值 val 的所有节点。
 *
 * 示例:
 *
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
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
	65/65 cases passed (8 ms)
	Your runtime beats 90.8 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.7 MB)
	时间O(n)，空间O(1)
*/
func removeElements(head *ListNode, val int) *ListNode {
	// 定义一个哨兵
	sentinel := &ListNode{Next: head}
	cur := sentinel.Next
	pre := sentinel
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return sentinel.Next
}

// @lc code=end
