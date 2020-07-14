package lc

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 *
 * https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/
 *
 * algorithms
 * Easy (55.60%)
 * Likes:    722
 * Dislikes: 0
 * Total Accepted:    125.4K
 * Total Submissions: 225.6K
 * Testcase Example:  '8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3'
 *
 * 编写一个程序，找到两个单链表相交的起始节点。
 *
 * 如下面的两个链表：
 *
 *
 *
 * 在节点 c1 开始相交。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2,
 * skipB = 3
 * 输出：Reference of the node with value = 8
 * 输入解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为
 * [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB
 * = 1
 * 输出：Reference of the node with value = 2
 * 输入解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为
 * [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
 * 输出：null
 * 输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为
 * 0，而 skipA 和 skipB 可以是任意值。
 * 解释：这两个链表不相交，因此返回 null。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 如果两个链表没有交点，返回 null.
 * 在返回结果后，两个链表仍须保持原有的结构。
 * 可假定整个链表结构中没有循环。
 * 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
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
	空间O(n)：hash表存节点指针，查找，不满足要求
	空间O(1)：若有交点(不存在环的情况)，则遍历完list1再从list2开始遍历，走相同的距离会相交到交点
		两个起点同时移动，直到有交点或移动完两个链表m+n长度
		46/46 cases passed (44 ms)
		Your runtime beats 89.71 % of golang submissions
		Your memory usage beats 80 % of golang submissions (7.5 MB)
		时间O(m+n)，空间O(1)
*/
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	// 两个起点
// 	n1 := headA
// 	n2 := headB
// 	// 标识自身是否遍历完
// 	end1 := false
// 	end2 := false
// 	for n1 != nil && n2 != nil {
// 		if n1 == n2 {
// 			return n1
// 		}

// 		n1 = n1.Next
// 		if !end1 && n1 == nil {
// 			n1 = headB
// 			end1 = true
// 		}

// 		n2 = n2.Next
// 		if !end2 && n2 == nil {
// 			n2 = headA
// 			end2 = true
// 		}
// 	}

// 	return nil
// }

/*
	调整一下逻辑更简洁
	46/46 cases passed (48 ms)
	Your runtime beats 70.66 % of golang submissions
	Your memory usage beats 80 % of golang submissions (7.9 MB)
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	// 两个起点
	n1 := headA
	n2 := headB

	for n1 != n2 {
		// 同时移动
		n1 = n1.Next
		n2 = n2.Next
		if n1 == n2 {
			return n1
		}
		// 移动到尾部了则从另一个链表继续，
		// 不会出现第二个遍历完重复再赋值遍历的情况，没有交点则最后肯定会同时为nil
		if n1 == nil {
			n1 = headB
		}
		if n2 == nil {
			n2 = headA
		}
	}

	// 若最后面n1==n2，则退出for循环，并不会从上面的return返回，所以此处不能直接返回nil
	return n1
}

// @lc code=end
