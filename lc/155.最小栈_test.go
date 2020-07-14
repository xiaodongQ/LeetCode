package lc

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (54.80%)
 * Likes:    600
 * Dislikes: 0
 * Total Accepted:    138.6K
 * Total Submissions: 252.9K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 *
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 *
 *
 * 提示：
 *
 *
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 *
 *
*/

// @lc code=start
/*
   18/18 cases passed (16 ms)
   Your runtime beats 95.68 % of golang submissions
   Your memory usage beats 75 % of golang submissions (7.5 MB)
   几个要注意的点：
    slice的len直接从底层结构获取，不用计算
    pop时，可能把最小值给删除了(栈顶元素和最小值一样才需要删除最小值数据栈顶)，这时第二小的数就变成最小了，
        若再pop时也是最小值则原来第三小的数变成最小
        所以另外维护一份处理最小值的数据(也可以当做是栈)
*/
type MinStack struct {
	data []int
	// 可以不用，len(slice)也不耗性能，slice底层为 数据指针+切片大小+切片容量
	// num int

	// 单独维护一份数据存最小值。只用两个int保存最小和第二小的数并不够
	mindata []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.mindata) == 0 || x <= this.mindata[len(this.mindata)-1] {
		this.mindata = append(this.mindata, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	if this.data[len(this.data)-1] == this.mindata[len(this.mindata)-1] {
		this.mindata = this.mindata[:len(this.mindata)-1]
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mindata[len(this.mindata)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
