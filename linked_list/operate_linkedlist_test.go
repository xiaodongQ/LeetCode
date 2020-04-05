/*
 * @Description:
 * @Author: xd
 * @Date: 2020-04-05 23:26:07
 * @LastEditTime: 2020-04-06 00:15:03
 * @LastEditors: xd
 * @Note:
 */
package main

import (
	"log"
	"testing"
)

// TestLinkedList 防止和main冲突，用单元测试来执行
func TestLinkedList(t *testing.T) {
	var intArr = []int{1, 0, 1}
	linkinfo := InitLinkInfo(intArr)
	log.Printf("%+v, sum:%d", intArr, getDecimalValue(linkinfo))
}

/*
* 给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。
* 已知此链表是一个整数数字的二进制表示形式。请你返回该链表所表示数字的 十进制值
* 链表非空，结点总数不超过30，结点取值只有0和1
* [1290. 二进制链表转整数](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/)
* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :2 MB, 在所有 Go 提交中击败了58.33%的用户
* 空间O(1)，时间O(n)
* 并不需要先存数组，直接每次把上一次的值左移(*2)即可
 */
func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}
