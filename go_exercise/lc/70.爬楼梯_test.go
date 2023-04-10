package lc

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (49.98%)
 * Likes:    1127
 * Dislikes: 0
 * Total Accepted:    238.8K
 * Total Submissions: 477.8K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
/*
	经典的递归，f(n) = f(n-1) + f(n-2)，
	爬n阶的方法等于 本次爬1步后剩余步数的方法 加上 本次爬2步后剩余步数的方法
	第0到第0也当作1种，递归终止条件变成
		f(0)=1，f(1)=1
	这样会有很多重复的值，可以用一个哈希表保存算过的值

	递归求解提交时，会超时(n=44)，Time Limit Exceeded，所以不满足要求
*/
// var m map[int]int

// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	a := 0
// 	b := 0
// 	if v, ok := m[n-1]; ok {
// 		a = v
// 	} else {
// 		a = climbStairs(n - 1)
// 	}

// 	if v, ok := m[n-2]; ok {
// 		b = v
// 	} else {
// 		b = climbStairs(n - 2)
// 	}
// 	return a + b
// }

/*
	使用map超时了，涉及哈希表的插入和扩容(扩容时间复杂度要O(n))，使用先创建好的数组，速度满足要求
	这种用法叫 "备忘录"？
	此处递归是自顶向下(第n步找n-1,n-2步)，改成自底向上(从第1步向上)就是动态规划，这个之后的示例做了展示
	45/45 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 36.36 % of golang submissions (2 MB)
	时间O(n)，空间O(n)
*/
var s []int

func climbStairs(n int) int {
	s = make([]int, n+1)
	return climbHelper(n)
}
func climbHelper(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if s[n] == 0 {
		s[n] = climbHelper(n-1) + climbHelper(n-2)
	}
	return s[n]
}

/*
	动态规划
	这种方式和递归差不多，从低到高依次保存低台阶的步数

	45/45 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 90.91 % of golang submissions (1.9 MB)
	时间O(n)，空间O(1)
*/
// func climbStairs(n int) int {
// 	// 分别保存 i-2, i-1, i阶 需要的步数
// 	p, q, r := 0, 0, 1
// 	for i := 1; i <= n; i++ {
// 		p = q
// 		q = r
// 		r = p + q
// 	}
// 	return r
// }

// 和上面一样，按自己的理解顺一点。时间O(n)，空间O(1)
// 如果是生产环境，可以查一下按通项进行求解，用到指数函数pow，其时间O(logn)
// func climbStairs(n int) int {
// 	if n == 1 || n == 0 {
// 		return 1
// 	}
// 	// i-2步
// 	pre2 := 1
// 	// i-1步
// 	pre1 := 1
// 	// 本次步数
// 	cur := 0
// 	for i := 2; i <= n; i++ {
// 		cur = pre2 + pre1
// 		pre2 = pre1
// 		pre1 = cur
// 	}
// 	return cur
// }

// @lc code=end
