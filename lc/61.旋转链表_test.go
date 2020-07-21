package lc

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (40.46%)
 * Likes:    297
 * Dislikes: 0
 * Total Accepted:    73.1K
 * Total Submissions: 180.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 *
 *
 * 示例 2:
 *
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
 *
 */
func TestRotateRight(t *testing.T) {
	list := InitLinkInfo([]int{1, 2, 3, 4, 5})
	PrintLinkInfo(list)

	PrintLinkInfo(rotateRight(list, 2))
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
错误：
	借助临时空间(不能有环)，重建链表
	有环就死循环了
*/
// func rotateRight(head *ListNode, k int) *ListNode {
// 	s := make([]*ListNode, 0)
// 	for head != nil {
// 		s = append(s, head)
// 		head = head.Next
// 		log.Printf("111len:%d", len(s))
// 	}
// 	start := (len(s) - k%len(s)) % len(s)
// 	res := s[start]
// 	log.Printf("len:%d, start:%d", len(s), start)
// 	for i := 0; i < len(s)-1; i++ {
// 		s[(start+i)%len(s)].Next = s[(start+i+1)%len(s)]
// 		log.Printf("i:%d, (start+i)modlen(s):%d, (start+i+1)modlen(s):%d",
// 			i, (start+i)%len(s), (start+i+1)%len(s))
// 	}
// 	// 原来的尾节点需要断开，否则新的链表有环
// 	s[len(s)-k%len(s)].Next = nil
// 	log.Printf("end")

// 	return res
// }

/*
	新的链表头位置： n-k%n，（k>0，若=0，则要再%n）
		新的链表尾：n-k%n-1 (k>=0都不影响)
	把原链表连接成环，在新尾部断开
	231/231 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.5 MB)
	时间O(n)，空间O(1)
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head
	n := 1
	// 注意找老的链表尾给tail赋值不能一直到nil
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	// 连接成环
	tail.Next = head

	// 找新的链表尾
	newtail := head
	for i := 0; i < n-k%n-1; i++ {
		newtail = newtail.Next
	}
	// 尾部的后一个即新链表头
	newHead := newtail.Next
	newtail.Next = nil

	return newHead
}

// @lc code=end
