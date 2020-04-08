/*
 * @Description:
 * @Author: xd
 * @Date: 2020-04-05 23:26:07
 * @LastEditTime: 2020-04-08 21:50:47
 * @LastEditors: xd
 * @Note:
 * [234. 回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/submissions/)
 */
package main

import (
	"log"
)

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
* 使用数组先存储链表结点的数据，再依次(len/2次)比较两头的数据是否满足回文
* 执行用时 :16 ms, 在所有 Go 提交中击败了75.91%的用户
* 内存消耗 :7.1 MB, 在所有 Go 提交中击败了19.51%的用户
* 空间复杂度O(n)，时间复杂度 复制O(n)+判断O(n/2)=O(n)
 */
func isPalindromeByArr(head *ListNode) bool {
	if head == nil {
		// 输入为空时当作满足回文
		return true
	}
	valArr := make([]int, 0)
	for head != nil {
		valArr = append(valArr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(valArr)/2; i++ {
		if valArr[i] != valArr[len(valArr)-1-i] {
			return false
		}
	}

	return true
}

/*
* 双指针，不同的步进长度
* 执行用时 :16 ms, 在所有 Go 提交中击败了75.98%的用户
* 内存消耗 :6 MB, 在所有 Go 提交中击败了79.41%的用户
* 空间复杂度O(1), 时间复杂度 中间结点O(n/2)+反转O(n/2)+比较O(n/2)+反转O(n/2) = O(n)
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 遍历
	var slow = head
	var fast = head
	// 这样能保证奇数个数时，slow在中间位置
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 从中间之后的后半部分(中间结点不算在后半部分)反转
	var current = slow.Next
	var previous *ListNode
	for current != nil {
		nextNode := current.Next
		current.Next = previous
		previous = current
		current = nextNode
	}

	// 依次比较前半部分和反转后的后半部分的结点(个数以反转部分为准)
	oriHead := head
	for previous != nil {
		if previous.Val != oriHead.Val {
			return false
		}
		previous = previous.Next
		oriHead = oriHead.Next
	}

	// 再对之前的反转还原
	current = slow.Next
	for current != nil {
		nextNode := current.Next
		current.Next = previous
		previous = current
		current = nextNode
	}

	return true
}

func main() {
	var intArr = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	linkinfo := InitLinkInfo(intArr)
	// 打印链表信息
	tempptr := linkinfo
	for tempptr != nil {
		log.Printf("node value:%v\n", tempptr.Val)
		tempptr = tempptr.Next
	}

	// isPalindromic := isPalindromeByArr(linkinfo)
	isPalindromic := isPalindrome(linkinfo)
	if isPalindromic {
		log.Printf("linkinfo:%+v is Palindromic", linkinfo)
	} else {
		log.Printf("linkinfo:%+v is not Palindromic", linkinfo)
	}
}

// InitLinkInfo 构造一个链表
func InitLinkInfo(intArr []int) *ListNode {
	// 反向遍历
	var preNodePtr *ListNode // 保存前一个位置的指针
	for i := len(intArr) - 1; i >= 0; i-- {
		node := ListNode{
			Val:  intArr[i],
			Next: preNodePtr,
		}
		preNodePtr = &node
	}

	return preNodePtr
}

// ConvertLinkInfo 用于打印链表
func ConvertLinkInfo(head *ListNode) []int {
	var intArr []int
	for head != nil {
		intArr = append(intArr, head.Val)
		head = head.Next
	}

	return intArr
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
