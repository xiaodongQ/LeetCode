/*
 * @Description:
 * @Author: xd
 * @Date: 2020-04-05 23:26:07
 * @LastEditTime: 2020-04-07 23:25:59
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
* 说明：链表非空，结点总数不超过30，结点取值只有0和1
* [1290. 二进制链表转整数](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/)
*
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

/*
* 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
* 说明：链表至少包含两个节点；所有节点的值都是唯一的；给定的节点为非末尾节点并且一定是链表中的一个有效节点；函数不用返回任何结果
* [237. 删除链表中的节点](https://leetcode-cn.com/problems/delete-node-in-a-linked-list/)
*
* 注意此处只送了要删除的结点，并不是从链表中找指定值的结点进行删除，直接操作相邻结点即可，太简单而让人怀疑。。
* 执行用时 :4 ms, 在所有 Go 提交中击败了80.65%的用户
* 内存消耗 :2.9 MB, 在所有 Go 提交中击败了68.42%的用户
* 时间O(1)，空间O(1)
 */
func deleteNode(node *ListNode) {
	/*
	   之前想的是将下一个结点指针赋值给当前结点，并将结点中的值替换(还想着需要深拷贝。。。没意义)
	   不过结点指针替换后，前一个结点的后继指针还是指向原来的结点指针，整个链表并没有改变
	   node := node.Next
	*/

	// 直接将值和下一个的指针重新指向为下一个指针
	// 下面语句可以一个语句完成：*node = *node.Next，提交后
	// 执行用时 :4 ms, 在所有 Go 提交中击败了85.37%的用户，性能略高
	// 内存消耗 :2.9 MB, 在所有 Go 提交中击败了100.00%的用户
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	/*
	   C++实现时，注意要delete的结点是要删除结点node的Next指针，而不是node本身
	   吐槽一下，用Go来练LeetCode导致写C++的时候老用`.`来解引用指针中的成员。。。

	   此处C++提交的性能比Go差不少：
	   执行用时 :16 ms, 在所有 C++ 提交中击败了24.86%的用户
	   内存消耗 :8 MB, 在所有 C++ 提交中击败了100.00%的用户
	*/
}

/*
* 实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值
* 说明：给定的 k 保证是有效的
* [面试题 02.02. 返回倒数第 k 个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)
*
* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :2.5 MB, 在所有 Go 提交中击败了100.00%的用户
* 时间O(n)，空间O(1)
 */
func TestKthToLast(t *testing.T) {
	var intArr = []int{1, 2, 3, 4, 5}
	k := 2
	linkinfo := InitLinkInfo(intArr)
	log.Printf("%+v, k:%d, return:%d", intArr, k, kthToLast2(linkinfo, k))
}
func kthToLast(head *ListNode, k int) int {
	// 拷贝到数组，取len-k个位置，时间O(n)，空间O(n)
	// 双指针，跟找中间结点类似，只是相对移动的结点为fast，且步长为k
	slow := head
	fast := head
	// head即为尾结点，k为1的话，直接返回slow
	for slow != nil {
		fast = slow
		// 需要k次判断(而不是<k-1)，否则k=2时，只判断一次，循环中的if并不会执行第二次
		for i := 0; i < k; i++ {
			if fast.Next == nil {
				// 说明k步之内会到达尾部
				log.Printf("already i:%d, slow:%d\n", i, slow.Val)
				return slow.Val
			}
			fast = fast.Next
		}
		slow = slow.Next
	}
	return slow.Val
}

/*
* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :2.2 MB, 在所有 Go 提交中击败了100.00%的用户
* 时间O(n)，空间O(1)
 */
func kthToLast2(head *ListNode, k int) int {
	slow := head
	fast := head
	// 已知k是有效的，则可以让fast先移动k-1步，再同步移动两个指针，直到fast为尾结点
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Val
}

/*
* 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
* 说明：0 <= 节点个数 <= 5000
* [面试题24. 反转链表](https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/)
* [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :2.6 MB, 在所有 Go 提交中击败了100.00%的用户
* 时间O(n)，空间O(1)
 */
func TestReverseList(t *testing.T) {
	var intArr = []int{1, 2, 3, 4, 5}
	linkinfo := InitLinkInfo(intArr)
	log.Printf("%+v, return:%+v", intArr, reverseList(linkinfo))
}
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	var next *ListNode
	for head != nil {
		// 先保存下一个结点
		next = head.Next

		head.Next = previous
		previous = head
		head = next
	}
	return previous
}

/*
* 递归方式
* 从第一个结点开始进入递归，把当前结点(head)的后一个结点的Next指向当前结点(head)；
*
* 第一个结点的处理依旧还停留在递归调用的那行语句，后续的递归分析：

* 进入递归后(调用递归的那行语句)会直到递归到最后一个结点，
	最后一个结点(也是最后一层)传入递归调用时，触发终止条件(head.Next==nil)，返回最后结点的指针
	往上每层的递归返回，都是将最后结点的指针层层返回上去
* 每层递归中除了层层返回最后结点的指针外，还会把传入结点的后一个结点指向传入结点
	并将传入结点的Next置nil(从最后一个结点看好理解一点)
	所以里层的递归将传入结点(假设称A)的Next置nil，而外面一层递归(传入的是A的上一个结点B)又将A(即B的下一个结点)的Next指向B
	所以对每个非首结点(还没执行到下一行)和非尾结点(只改了一次Next而没有置nil，而且是传入其上一个结点的那层递归改的)来说，
	其Next都是先置为nil，再置为其上一个结点的指针
* 对于首结点，递归调用行层层返回尾结点的指针，然后继续执行下去(此时是递归最外层，已经没有其他递归了)
	执行将首结点的下一个结点指向首结点，然后首结点的Next置nil，
	然后整个调用就结束了
*/
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListRecursion(head.Next) // 递归调用行，直到最后一个结点触发终止条件
	head.Next.Next = head
	head.Next = nil

	return p
}

/*
* 给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
* 如果有两个中间结点，则返回第二个中间结点。
* 说明：给定链表的结点数介于 1 和 100 之间
* [876. 链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/)
*
* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
* 时间O(n)，空间O(1)
 */
func TestMiddleNode(t *testing.T) {
	var intArr = []int{1, 2, 3, 4, 5}
	linkinfo := InitLinkInfo(intArr)
	log.Printf("%+v, return:%+v", intArr, middleNode(linkinfo))
}
func middleNode(head *ListNode) *ListNode {
	// 双指针法，fast步长为2
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 另外若有两个中间结点，则要返回后面的结点
	if fast.Next != nil && fast.Next.Next == nil {
		return slow.Next
	}

	return slow
}
