package lc

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (55.73%)
 * Likes:    495
 * Dislikes: 0
 * Total Accepted:    109K
 * Total Submissions: 195.5K
 * Testcase Example:  '1'
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
 *
 * 注意：整数序列中的每一项将表示为一个字符串。
 *
 * 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 第一项是数字 1
 *
 * 描述前一项，这个数是 1 即 “一个 1 ”，记作 11
 *
 * 描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
 *
 * 描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
 *
 * 描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 * 解释：这是一个基本样例。
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 * 解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 =
 * 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
 *
 */

// @lc code=start
/*

 */

/*
	递归求解前一个
	时间：每个字符串都要来一遍，每遍和长度成线性，O(mn)，m为个数，n为平均长度
	空间O(n)
	空间：递归n次，O(n)

使用fmt.Sprintf构造res字符串：
	18/18 cases passed (4 ms)
	Your runtime beats 48.74 % of golang submissions
	Your memory usage beats 14.29 % of golang submissions (6.4 MB)
使用string.Builder：
	fmt.Sprintf每次会新建string，使用strings.Builder来构建字符串，时间和空间都好很多
	18/18 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 57.14 % of golang submissions (2.2 MB)
*/
// func countAndSay(n int) string {
// 	if n == 1 {
// 		return "1"
// 	}
// 	pre := countAndSay(n - 1)
// 	num := 1
// 	// fmt.Sprintf每次会新建string，使用strings.Builder来更新
// 	var res strings.Builder
// 	for i := 1; i < len(pre); i++ {
// 		if pre[i] == pre[i-1] {
// 			num++
// 		} else {
// 			res.WriteString(strconv.Itoa(num))
// 			res.WriteByte(pre[i-1])
// 			num = 1
// 		}
// 	}
// 	// 最后一次没放进去
// 	res.WriteString(strconv.Itoa(num))
// 	res.WriteByte(pre[len(pre)-1])

// 	return res.String()
// }

/*
	由于使用栈有栈溢出的风险(虽然此处n<=30不存在)，不加另外的栈深度保护的话，调整成循环来处理
	时间复杂度O(mn)(m为个数，n为平均长度)，空间复杂度O(1)
	18/18 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 85.71 % of golang submissions (2.1 MB)
*/
func countAndSay(n int) string {
	var res strings.Builder

	for t := 0; t < n; t++ {
		if t == 0 {
			res.WriteString("1")
			continue
		}
		pre := res.String()
		res.Reset()

		num := 1
		for i := 1; i < len(pre); i++ {
			if pre[i] == pre[i-1] {
				num++
			} else {
				res.WriteString(strconv.Itoa(num))
				res.WriteByte(pre[i-1])
				num = 1
			}
		}
		// 最后一次没放进去
		res.WriteString(strconv.Itoa(num))
		res.WriteByte(pre[len(pre)-1])
	}

	return res.String()
}

// @lc code=end
