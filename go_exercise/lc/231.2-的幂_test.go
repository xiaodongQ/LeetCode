package lc

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 *
 * https://leetcode-cn.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (48.31%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    68K
 * Total Submissions: 140.8K
 * Testcase Example:  '1'
 *
 * 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: true
 * 解释: 2^0 = 1
 *
 * 示例 2:
 *
 * 输入: 16
 * 输出: true
 * 解释: 2^4 = 16
 *
 * 示例 3:
 *
 * 输入: 218
 * 输出: false
 *
 */

// @lc code=start
/*
	1108/1108 cases passed (4 ms)
	Your runtime beats 48.84 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.2 MB)
	时间O(logn)，空间O(1)
*/
// func isPowerOfTwo(n int) bool {
// 	if n < 1 {
// 		return false
// 	}
// 	v := 1
// 	for v < n {
// 		v <<= 1
// 	}
// 	return v == n
// }

/*
	有更简单的解法
	位运算，x&(-x) 会保留最后一位的1，若结果还是等于x，则说明只有一个1，说明为2的幂次方
		-x的补码表示为：所有位取反 + 1
	e.g.
		7:  0000 0111
		-7: 各位取反1111 1000，加1得:
			1111 1001
		按位与，则第一个1保留，其他位都是0，其十进制为1，和7不相等，所以不为2的指数
	e.g.
		4:  0000 0100
		-4: 各位取反1111 1011，加1得:
			1111 1100
		按位与，则第一个1保留，其他位都是0，其十进制为4，和4相等，所以为2的指数

	1108/1108 cases passed (4 ms)
	Your runtime beats 48.84 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.2 MB)
	时间O(1)，空间O(1)
*/
// func isPowerOfTwo(n int) bool {
// 	// 运算符优先级？ go里面按位&、|、^ 比 == 优先级高
// 	// C里面，==优先级比&高，需要加括号先算&
// 	// 最好还是用括号来明确执行顺序
// 	return (n&(-n)) == n && n > 0
// }

/*
	另外一种位运算思路
	x&(x-1)：可以去掉最右边的1，去掉之后为0则说明只有一个1，说明是2的幂次方
	1108/1108 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.2 MB)
*/
func isPowerOfTwo(n int) bool {
	// 要过滤n为0的情况
	return (n&(n-1)) == 0 && n > 0
}

// @lc code=end
