package lc

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (58.38%)
 * Likes:    1131
 * Dislikes: 0
 * Total Accepted:    383.4K
 * Total Submissions: 656.7K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */

// @lc code=start
/*
提交结果：
	11509/11509 cases passed (24 ms)
	Your runtime beats 41.4 % of golang submissions
	Your memory usage beats 44 % of golang submissions (5.4 MB)
	时间O(logn)，和位数相关，空间O(1)
*/
// func isPalindrome(x int) bool {
// 	s := strconv.Itoa(x)
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

/*
	不转换为字符串
	反转后的数字计算出来，和原数字比较；
	但存在溢出整型范围的问题，可只反转一半位数，当反转的数位>=原数剩下的位数则说明已经一半了

提交(最后return时判断)：
	11509/11509 cases passed (24 ms)
	Your runtime beats 41.4 % of golang submissions
	Your memory usage beats 52 % of golang submissions (5.2 MB)
判断是否return true的逻辑放到for循环中，速度更快(判断放不同地方，特殊值和临届判断不一样)：
(好吧，貌似对比调整前后跑了几次时间没咋变的，不应该有缓存吧。。)
	11509/11509 cases passed (16 ms)
	Your runtime beats 78.85 % of golang submissions
	Your memory usage beats 88 % of golang submissions (5.2 MB)
时间O(logn)，空间O(1)
*/
func isPalindrome(x int) bool {
	// 负数不为回文
	// 个位数除为回文
	// 个位为0的数，除0外都不为回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	} else if x < 10 {
		return true
	}

	half := 0
	for half < x {
		half = half*10 + x%10
		x = x / 10
		// 若原数位数为奇数，则用最后反转的数字/10判断
		// 判断放到此处，比处理完后放到循环外要快，少了奇数位时的一次反转
		if half == x || half == x/10 {
			return true
		}
	}

	return false
}

// @lc code=end
