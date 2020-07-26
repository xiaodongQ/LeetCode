package lc

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 *
 * https://leetcode-cn.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (49.03%)
 * Likes:    124
 * Dislikes: 0
 * Total Accepted:    28.2K
 * Total Submissions: 57.5K
 * Testcase Example:  '16'
 *
 * 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 16
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */

// @lc code=start
/*
	和3的幂一样，其他方法 转换进制、最大%n
		另外多一种利用2进制，仅最高位为1，且是2的倍数位(第0位则num为1，满足4^0)
	1060/1060 cases passed (4 ms)
	Your runtime beats 47.92 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.1 MB)
	时间O(logn)，空间O(1)
*/
// func isPowerOfFour(num int) bool {
// 	if num < 1 {
// 		return false
// 	}
// 	for num%4 == 0 {
// 		num = num / 4
// 	}
// 	return num == 1
// }

/*
	只要满足最高为1，且1在偶数位(0位开始算)
	1在0101 0101形式的位上，共32位，4字节
		num&(0x55555555) != 0
		也可以 num&(0xaaaaaaaa) == 0

	1060/1060 cases passed (4 ms)
	Your runtime beats 47.92 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.1 MB)
	时间O(1)，空间O(1)
*/
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&(0x55555555) != 0
}

/*
	另外一个数学解法：满足2的指数且num%3 == 1，即为4的指数
	num%3 == 1的证明：
		4^n % 3 可以写成 (3+1)^n % 3
		(3+1) * (3+1)^(n-1) = 3 * ((3+1)^(n-1)) + (3+1)^(n-1)，前面部分%3肯定为0
		则依此类推，nums%3最后直到 (3+1)^0 % 3，得到1
		所以若为4的指数，则num mod 3 = 1
	时间空间和上面一样
*/
// func isPowerOfFour(num int) bool {
// 	return num > 0 && num&(num-1) == 0 && num%3 == 1
// }

// @lc code=end
