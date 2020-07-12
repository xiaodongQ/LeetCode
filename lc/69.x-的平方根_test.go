package lc

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.77%)
 * Likes:    440
 * Dislikes: 0
 * Total Accepted:    172.8K
 * Total Submissions: 445.7K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
/*
	要求的整数满足最大的 k*k <= x，二分查找依次找满足条件的中间值
	1017/1017 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 14.29 % of golang submissions (2.2 MB)
	时间O(logx)，空间O(1)

	牛顿迭代法，效率更高一些，但数学推导和横轴的交点，需要参考链接理解一下
	要用的话可以先直接记迭代公式：x(i+1) = 1/2(x(i) + C/x(i))，C取值为x，直到x(i+1)和x(i)取值相差10^(-7)就当作结束
*/
func mySqrt(x int) int {
	res := -1
	l, r := 0, x

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}

/*
	牛顿迭代法借助泰勒级数，效率更高一些，但数学推导和横轴的交点，需要参考链接理解一下
	要用的话可以先直接记迭代公式：x(i+1) = 1/2(x(i) + C/x(i))，C取值为x，直到x(i+1)和x(i)取值相差10^(-7)就当作结束
	从大到小收敛

	1017/1017 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 14.29 % of golang submissions (2.2 MB)
	时间O(logn)，空间O(1)

*/
// func mySqrt(x int) int {
// 	if x == 0 {
// 		return 0
// 	}
// 	res := 0
// 	cur := float64(0.0)
// 	C, pre := float64(x), float64(x)
// 	for {
// 		cur = 0.5 * (pre + C/pre)
// 		if math.Abs(pre-cur) < 1e-7 {
// 			res = int(cur)
// 			break
// 		}
// 		pre = cur
// 	}
// 	return res
// }

// @lc code=end
