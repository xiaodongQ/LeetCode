/*
 * @Description:
 * @Author: xd
 * @Date: 2020-04-25 23:28:10
 * @LastEditTime: 2020-05-17 00:15:19
 * @LastEditors: xd
 * @Note:
 */
package main

/*
 * 实现一个MyQueue类，该类用两个栈来实现一个队列
 * [面试题 03.04. 化栈为队](https://leetcode-cn.com/problems/implement-queue-using-stacks-lcci/)
 *
 // 题目里是要定义两个栈，把一个栈当队列的入栈，一个当出栈，出列时如果出栈中有数据，则pop一个数据，如果没有则入栈数据依次逆序存到出栈
 // 对于出栈，栈顶数据是原始最先入列的数据，pop时对于原始数据满足先进先出
 // 对于入栈，栈顶是最后进入的，出列时需要最后出，且不能直接pop出最先入栈的数据
 // 只用一个slice，直接存取能通过该题，均击败100%，但和题意不符，评论里的例子也有不少只用一个数组，并没有模拟栈的操作
 * 使用两个栈：
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
*/
type MyQueue struct {
	stackin  []int
	stackout []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stackin = append(this.stackin, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 出栈中有记录则从出栈最后一个位置
	if len(this.stackout) != 0 {
		res := this.stackout[len(this.stackout)-1]
		this.stackout = this.stackout[:len(this.stackout)-1]
		return res
	}
	// 出栈中无记录，则需要从入栈逆序获取一个数据，逆序时将数据放到出栈
	for {
		if len(this.stackin) != 0 {
			if len(this.stackin) == 1 {
				// 如果只剩下一个成员，就不push到出栈里了，直接出列就是该成员
				res := this.stackin[len(this.stackin)-1]
				this.stackin = this.stackin[:len(this.stackin)-1]
				return res
			}
			// 一个一个取出来push到出栈中，并删除入栈的该记录
			this.stackout = append(this.stackout, this.stackin[len(this.stackin)-1])
			this.stackin = this.stackin[:len(this.stackin)-1]
		} else {
			// 注意退出条件
			break
		}
	}

	return -999
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stackout) != 0 {
		return this.stackout[len(this.stackout)-1]
	} else if len(this.stackin) != 0 {
		return this.stackin[0]
	}
	return -999
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stackin) == 0 && len(this.stackout) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
