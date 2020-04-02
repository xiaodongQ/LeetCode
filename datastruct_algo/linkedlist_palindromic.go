package main

import (
	"log"
)

/*
 * @Description:
 * @Author: xd
 * @Date: 2020-04-02 22:14:23
 * @LastEditTime: 2020-04-02 23:48:23
 * @LastEditors: xd
 * @Note:
 */

type linkedInfo struct {
	value string
	next  *linkedInfo
}

func judgePalindromic(nodePtr *linkedInfo) bool {
	// 遍历
	var flag1 = nodePtr
	var flag2 = nodePtr
	nodeCount := 0
	for flag1 != nil {
		if flag2 == nil || flag2.next == nil {
			// 说明此时flag1到了中间点,flag2到了最后一个结点
			break
		}
		nodeCount++
		flag1 = flag1.next
		flag2 = flag2.next.next
	}
	// 说明有奇数个，则flag1位置处不需要
	if nodeCount%2 == 0 {
	}

	return false
}

func initLinkInfo(str string) *linkedInfo {
	// 反向遍历
	var nodePtr *linkedInfo // 保存前一个位置的指针
	for i := len(str) - 1; i >= 0; i-- {
		note := linkedInfo{
			value: str[i : i+1],
			next:  nodePtr,
		}
		nodePtr = &note
	}

	// 打印链表信息
	for nodePtr != nil {
		log.Printf("node value:%s, ptr:%v\n", nodePtr.value, nodePtr)
		nodePtr = nodePtr.next
	}

	return nodePtr
}

func main() {
	// const str = "123454321"
	const str = "123abc"
	linkinfo := initLinkInfo(str)

	isPalindromic := judgePalindromic(linkinfo)
	if isPalindromic {
		log.Printf("linkinfo:%+v is Palindromic", linkinfo)
	} else {
		log.Printf("linkinfo:%+v is not Palindromic", linkinfo)
	}
}
