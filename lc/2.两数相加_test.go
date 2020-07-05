package lc

import "testing"

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (37.70%)
 * Likes:    4541
 * Dislikes: 0
 * Total Accepted:    466.6K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

func TestAddTwoNumbers(t *testing.T) {
	list1 := InitLinkInfo([]int{1})
	list2 := InitLinkInfo([]int{9, 9})
	PrintLinkInfo(addTwoNumbers(list1, list2))
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 提交：
1563/1563 cases passed (8 ms)
Your runtime beats 93.49 % of golang submissions
Your memory usage beats 58.62 % of golang submissions (5 MB)
时间O(max(M,N))，空间O(max(M,N))，新链表最大长度为max(M,N)+1
*/
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// 定义一个哨兵节点，只指向链表，不存数据，
//	// 哨兵技巧：链表初始化无可用节点，或者操作过程中头指针可能变化，用一个哨兵可使问题简单很多
// 	var newListHead = &ListNode{}
// 	var preNode = newListHead
// 	inc := 0
// 	for l1 != nil && l2 != nil {
// 		newNode := &ListNode{}
// 		newNode.Val = (l1.Val + l2.Val + inc) % 10
// 		inc = (l1.Val + l2.Val + inc) / 10
// 		preNode.Next = newNode

// 		l1 = l1.Next
// 		l2 = l2.Next
// 		preNode = newNode
// 	}

// 	var tmp *ListNode
// 	if l1 != nil {
// 		tmp = l1
// 	} else if l2 != nil {
// 		tmp = l2
// 	}
// 	for tmp != nil {
// 		newNode := &ListNode{}
// 		newNode.Val = (tmp.Val + inc) % 10
// 		inc = (tmp.Val + inc) / 10

// 		preNode.Next = newNode

// 		tmp = tmp.Next
// 		preNode = newNode
// 	}

// 	if inc == 1 {
// 		newNode := &ListNode{}
// 		newNode.Val = inc

// 		preNode.Next = newNode
// 	}

// 	return newListHead.Next
// }

/*
	和上面一样，优化重复代码(自己的代码比较丑陋。。。)
1563/1563 cases passed (8 ms)
Your runtime beats 93.49 % of golang submissions
Your memory usage beats 58.62 % of golang submissions (5 MB)
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义一个哨兵节点，只指向链表，不存数据
	var newListHead = &ListNode{}
	var preNode = newListHead
	sum := 0
	for l1 != nil || l2 != nil {
		sum /= 10
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		newNode := &ListNode{Val: sum % 10}
		preNode.Next = newNode
		preNode = newNode
	}

	if sum/10 == 1 {
		newNode := &ListNode{Val: 1}
		preNode.Next = newNode
	}

	return newListHead.Next
}

// @lc code=end
