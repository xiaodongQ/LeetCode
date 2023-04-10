package lc

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (67.51%)
 * Likes:    158
 * Dislikes: 0
 * Total Accepted:    44.2K
 * Total Submissions: 65.4K
 * Testcase Example:  '"A"'
 *
 * 给定一个Excel表格中的列名称，返回其相应的列序号。
 *
 * 例如，
 *
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28
 * ⁠   ...
 *
 *
 * 示例 1:
 *
 * 输入: "A"
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: "AB"
 * 输出: 28
 *
 *
 * 示例 3:
 *
 * 输入: "ZY"
 * 输出: 701
 *
 * 致谢：
 * 特别感谢 @ts 添加此问题并创建所有测试用例。
 *
 */

// @lc code=start
/*
	和 [168] Excel表列名称 相反的转换
	从高位到低位，每次循环先*26
	1000/1000 cases passed (4 ms)
	Your runtime beats 46.12 % of golang submissions
	Your memory usage beats 33.33 % of golang submissions (2.2 MB)
	时间O(n)，空间O(1)
*/
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 26
		res += int(s[i] - 'A' + 1)
	}
	return res
}

// @lc code=end
