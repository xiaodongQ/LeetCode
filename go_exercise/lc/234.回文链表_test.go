package lc

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (42.57%)
 * Likes:    564
 * Dislikes: 0
 * Total Accepted:    105.9K
 * Total Submissions: 248.7K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
不满足要求：
	快慢指针，主要要找到中间节点
	通过慢指针缓存前一半节点
	26/26 cases passed (12 ms)
	Your runtime beats 97.37 % of golang submissions
	Your memory usage beats 50 % of golang submissions (7.1 MB)
	时间O(n)，空间O(n) n/2，不满足题意，需要O(1)
*/
// func isPalindrome(head *ListNode) bool {
// 	if head == nil {
// 		return true
// 	}
// 	slow, fast := head, head
// 	// 要不要先把当前节点缓存，可以举两个简单的奇偶的例子
// 	// 若总数为奇数，停在中间，缓存后最后遍历时也从中间节点比较
// 	// 若总数为偶数，停在中间两个的右边
// 	cache := []*ListNode{}
// 	for fast != nil && fast.Next != nil {
// 		cache = append(cache, slow)
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}
// 	// 需要分奇偶处理，根据快指针走的最后位置来判断奇偶
// 	i := len(cache) - 1
// 	if fast == nil {
// 		// 每次走满了两步，说明为偶数
// 		// 此时slow到达了中间，走一步从cache最后取一个数比较
// 		for slow != nil {
// 			if slow.Val != cache[i].Val {
// 				return false
// 			}
// 			slow = slow.Next
// 			i--
// 		}
// 	} else {
// 		slow = slow.Next
// 		for slow != nil {
// 			if slow.Val != cache[i].Val {
// 				return false
// 			}
// 			slow = slow.Next
// 			i--
// 		}
// 	}

// 	return true
// }

/*
不满足要求：
	如果允许O(n)，则不用像上面那么复杂，直接拷贝链表值到数组，再比较数组是否为回文
	26/26 cases passed (24 ms)
	Your runtime beats 9.59 % of golang submissions
	Your memory usage beats 50 % of golang submissions (7.1 MB)
	时间O(n)，空间O(n)，但是题意要求空间O(1)
*/
// func isPalindrome(head *ListNode) bool {
// 	s := make([]int, 0)
// 	for head != nil {
// 		s = append(s, head.Val)
// 		head = head.Next
// 	}
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

/*
	获取中间节点，反转中间节点后面的一部分链表
	前后部分链表依次比较，最后还原中间节点后面的一部分链表(再次反转)
	操作比较独立，且有重复操作，写成单独的函数

	26/26 cases passed (16 ms)
	Your runtime beats 77.87 % of golang submissions
	Your memory usage beats 100 % of golang submissions (6 MB)
	时间O(n) n+n/2+n+n/2，空间O(1)
*/
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	res := true
	mid := getMidNode(head)
	// 后半部分反转后的链表
	part := reversePartList(mid.Next)
	for part != nil && head != nil {
		if head.Val != part.Val {
			res = false
			break
		}
		head = head.Next
		part = part.Next
	}
	// 反转恢复
	reversePartList(mid.Next)
	return res
}

func getMidNode(head *ListNode) *ListNode {
	slow, fast := head, head
	// 若为偶数则返回前一个中间节点
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reversePartList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur

		cur = tmp
	}
	return pre
}

// @lc code=end
