package lc

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (37.99%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    29.7K
 * Total Submissions: 78.1K
 * Testcase Example:  '1'
 *
 * 给定一个正整数，返回它在 Excel 表中相对应的列名称。
 *
 * 例如，
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "A"
 *
 *
 * 示例 2:
 *
 * 输入: 28
 * 输出: "AB"
 *
 *
 * 示例 3:
 *
 * 输入: 701
 * 输出: "ZY"
 *
 *
 */

// @lc code=start
/*
	进制转换，数字转换为26进制 Z进一位变成 AA (从1开始)
	用一个slice保存，最后翻转
	18/18 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (1.9 MB)
	时间O(n)，空间O(n) (除结果外O(1))
*/
func convertToTitle(n int) string {
	var s []byte
	for n > 0 {
		n--
		c := byte(n%26 + 'A')
		s = append(s, c)
		n /= 26
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	return string(s)
}

// @lc code=end
