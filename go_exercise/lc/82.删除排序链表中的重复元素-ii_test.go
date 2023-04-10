package lc

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (48.48%)
 * Likes:    322
 * Dislikes: 0
 * Total Accepted:    55.6K
 * Total Submissions: 114.7K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
	双指针，每次相对当前位置，用另外指针往后查找
	166/166 cases passed (4 ms)
	Your runtime beats 89.15 % of golang submissions
	Your memory usage beats 50 % of golang submissions (3 MB)
	时间O(n)，空间O(1)

	另外一种方法也可以借助辅助空间，记录每个值出现的次数，再遍历原链表，次数>1的跳过
*/
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	preNode := dummy

	for head != nil && head.Next != nil {
		// 记录当前节点
		cur := head
		// 预先判断该节点后面的节点是否相等
		for cur.Next != nil && head.Val == cur.Next.Val {
			cur = cur.Next
		}
		if cur != head {
			preNode.Next = cur.Next
			head = cur.Next
		} else {
			preNode = head
			head = head.Next
		}
	}
	return dummy.Next
}

// @lc code=end
