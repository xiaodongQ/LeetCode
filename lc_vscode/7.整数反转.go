package lcvscode

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
/*
提交结果：
	1032/1032 cases passed (4 ms)
	Your runtime beats 46.77 % of golang submissions
	Your memory usage beats 42.86 % of golang submissions (2.2 MB)
时间O(logn)，数字个数大概为log10(n)
空间O(1)
*/
func reverse(x int) int {
	// 会报溢出
	// max := int32(0x7fffffff)
	// min := int32(0x80000000)
	max := 2147483647
	min := -2147483648
	sum := 0
	for x != 0 {
		sum = sum*10 + x%10
		// 没计算完却已经越界则返回0
		if sum > max || sum < min {
			return 0
		}
		x = x / 10
	}

	return sum
}

// @lc code=end
