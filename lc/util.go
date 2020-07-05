/*
 * @Description:
 * @Author: xd
 * @Date: 2020-07-05 16:52:29
 * @LastEditTime: 2020-07-05 17:07:17
 * @LastEditors: xd
 * @Note:
 */
package lc

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

// InitLinkInfo 构造一个链表
func InitLinkInfo(intArr []int) *ListNode {
	// 反向遍历
	var preNodePtr *ListNode // 保存前一个位置的指针
	for i := len(intArr) - 1; i >= 0; i-- {
		node := &ListNode{
			Val:  intArr[i],
			Next: preNodePtr,
		}
		preNodePtr = node
	}

	return preNodePtr
}

// PrintLinkInfo 打印链表信息
func PrintLinkInfo(head *ListNode) {
	log.Printf("linkinfo:%v", ConvertLinkInfo(head))
}

// ConvertLinkInfo 用于打印链表
func ConvertLinkInfo(head *ListNode) []int {
	var intArr []int
	cur := head
	for cur != nil {
		intArr = append(intArr, cur.Val)
		cur = cur.Next
	}

	return intArr
}
