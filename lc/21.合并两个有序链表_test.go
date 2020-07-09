package lc

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (63.18%)
 * Likes:    1142
 * Dislikes: 0
 * Total Accepted:    308.1K
 * Total Submissions: 487.4K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
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
	208/208 cases passed (4 ms)
	Your runtime beats 61.26 % of golang submissions
	Your memory usage beats 18.18 % of golang submissions (2.6 MB)
	时间O(m+n)，空间O(m+n)
*/
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// 返回一个链表总是可以借助一个哨兵来简化头节点的操作
// 	sentinel := &ListNode{Val: 0}
// 	preNode := sentinel
// 	for l1 != nil || l2 != nil {
// 		v := 0
// 		if l1 != nil && l2 != nil {
// 			if l1.Val <= l2.Val {
// 				v = l1.Val
// 				l1 = l1.Next
// 			} else {
// 				v = l2.Val
// 				l2 = l2.Next
// 			}
// 		} else if l1 != nil {
// 			v = l1.Val
// 			l1 = l1.Next
// 		} else {
// 			v = l2.Val
// 			l2 = l2.Next
// 		}
// 		// 新定义了一个节点，增加了空间复杂度，题意是拼接已有节点
// 		newNode := &ListNode{Val: v}
// 		preNode.Next = newNode
// 		preNode = newNode
// 	}

// 	return sentinel.Next
// }

/*
	题意是拼接已有节点，所以可以调整对某个链表为空时的处理
	208/208 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 72.73 % of golang submissions (2.5 MB)
	时间O(min(m,n)) 空间O(1)
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 返回一个链表总是可以借助一个哨兵来简化头节点的操作
	sentinel := &ListNode{Val: 0}
	preNode := sentinel
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			preNode.Next = l1
			l1 = l1.Next
		} else {
			preNode.Next = l2
			l2 = l2.Next
		}

		preNode = preNode.Next
	}
	// 执行完后两个链表至少有一个为nil
	if l1 != nil {
		preNode.Next = l1
	} else if l2 != nil {
		preNode.Next = l2
	}

	return sentinel.Next
}

// @lc code=end
