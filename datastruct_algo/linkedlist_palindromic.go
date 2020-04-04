package main

import (
	"log"
)

// [234. 回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/submissions/)

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
* 执行用时 :16 ms, 在所有 Go 提交中击败了75.91%的用户
* 内存消耗 :7.1 MB, 在所有 Go 提交中击败了19.51%的用户
 */
func isPalindromeByArr(head *ListNode) bool {
	if head == nil {
		return true
	}
	valArr := make([]int, 0)
	for head != nil {
		valArr = append(valArr, head.Val)
		head = head.Next
	}

	for i := 0; i < len(valArr); i++ {
		if i >= (len(valArr) / 2) {
			return true
		}
		if valArr[i] != valArr[len(valArr)-1-i] {
			return false
		}
	}

	return false
}

func isPalindrome(head *ListNode) bool {
	// 遍历
	var flag1 = head
	var flag2 = head
	nodeCount := 0
	for flag1 != nil {
		if flag2 == nil || flag2.Next == nil {
			// 说明此时flag1到了中间点,flag2到了最后一个结点
			break
		}
		nodeCount++
		flag1 = flag1.Next
		flag2 = flag2.Next.Next
	}
	// 偶数次，则flag1位置为中间点
	if nodeCount%2 == 0 {

	}

	return false
}

func main() {
	iArr := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	// iArr := []int{1, 3, 4, 5, 4, 3, 2, 1}
	linkinfo := initLinkInfo(iArr)
	// 打印链表信息
	tempptr := linkinfo
	for tempptr != nil {
		log.Printf("node value:%v\n", tempptr.Val)
		tempptr = tempptr.Next
	}

	// isPalindromic := judgePalindromic(linkinfo)
	isPalindromic := isPalindromeByArr(linkinfo)
	if isPalindromic {
		log.Printf("linkinfo:%+v is Palindromic", linkinfo)
	} else {
		log.Printf("linkinfo:%+v is not Palindromic", linkinfo)
	}
}

func initLinkInfo(intArr []int) *ListNode {
	// 反向遍历
	var nodePtr *ListNode // 保存前一个位置的指针
	for i := len(intArr) - 1; i >= 0; i-- {
		note := ListNode{
			Val:  intArr[i],
			Next: nodePtr,
		}
		nodePtr = &note
	}

	return nodePtr
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
