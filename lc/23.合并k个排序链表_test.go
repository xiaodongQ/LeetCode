package lc

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (52.38%)
 * Likes:    801
 * Dislikes: 0
 * Total Accepted:    145.5K
 * Total Submissions: 277.7K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 *
 * 示例:
 *
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
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
	根据合并两个链表
方法1：
	顺序合并：合并后成为一个链表，再继续后面的链表
方法2：
	分治合并：两两配对进行合并，k个链表变为k/2，继续重复直到变成1个链表
方法3：
	使用优先队列合并(go里面并没有实现优先级队列)
*/
/*
	顺序合并
	131/131 cases passed (128 ms)
	Your runtime beats 24.31 % of golang submissions
	Your memory usage beats 77.78 % of golang submissions (5.3 MB)
	假设每个链表最长长度为n，第一次合并后n，第二次2n，...，第k次kn
		代价 n(1+k)k/2，所以时间复杂度为O((k^2) * n)
	空间O(1)
*/
// func mergeKLists(lists []*ListNode) *ListNode {
// 	var res *ListNode
// 	for i := 0; i < len(lists); i++ {
// 		res = mergeTwoList(res, lists[i])
// 	}
// 	return res
// }

/*
	分治合并
	每次选两个组成一对
	131/131 cases passed (8 ms)
	Your runtime beats 95.11 % of golang submissions
	Your memory usage beats 55.56 % of golang submissions (5.4 MB)
	时间O(kn * logk)，空间O(logk)
		第一轮合并，k/2对，每对2n，k/4对，每对4n，...，第i轮：k/(2^i) * (2^i)n
		求和为：kn×logk，所以时间复杂度O(kn * logk)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	return mergeHelper(lists, 0, len(lists)-1)
}

// 选择并合并
func mergeHelper(lists []*ListNode, l, h int) *ListNode {
	if l > h {
		return nil
	}
	if l == h {
		// 同一个链表，随便返回一个
		return lists[l]
	}
	mid := l + (h-l)/2
	// 合并左右两部分
	return mergeTwoList(mergeHelper(lists, l, mid), mergeHelper(lists, mid+1, h))
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	sentinel := &ListNode{}
	cur := sentinel
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return sentinel.Next
}

// @lc code=end
