package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 *
 * https://leetcode-cn.com/problems/roman-to-integer/description/
 *
 * algorithms
 * Easy (61.78%)
 * Likes:    940
 * Dislikes: 0
 * Total Accepted:    217.1K
 * Total Submissions: 351.4K
 * Testcase Example:  '"III"'
 *
 * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V +
 * II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
 *
 * 示例 1:
 *
 * 输入: "III"
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: "IV"
 * 输出: 4
 *
 * 示例 3:
 *
 * 输入: "IX"
 * 输出: 9
 *
 * 示例 4:
 *
 * 输入: "LVIII"
 * 输出: 58
 * 解释: L = 50, V= 5, III = 3.
 *
 *
 * 示例 5:
 *
 * 输入: "MCMXCIV"
 * 输出: 1994
 * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
 *
 */
func TestRomanToInt(t *testing.T) {
	log.Printf("res:%v", romanToInt("III"))
}

// @lc code=start
// var roman = map[byte]int{
// 	'I': 1,
// 	'V': 5,
// 	'X': 10,
// 	'L': 50,
// 	'C': 100,
// 	'D': 500,
// 	'M': 1000,
// }
// var special = map[string]int{
// 	"IV": 4,
// 	"IX": 9,
// 	"XL": 40,
// 	"XC": 90,
// 	"CD": 400,
// 	"CM": 900,
// }

// /*
// 3999/3999 cases passed (8 ms)
// Your runtime beats 75.38 % of golang submissions
// Your memory usage beats 40 % of golang submissions (3.1 MB)
// 时间O(N)，每次要取下一个字符判断是否满足组合，不如下面比较上个值的方法效率高
// */
// func romanToInt(s string) int {
// 	res := 0
// 	i := 0
// 	for ; i < len(s)-1; i++ {
// 		// 除了几个特殊组合，其他只要叠加即可
// 		if s[i] == 'I' || s[i] == 'X' || s[i] == 'C' {
// 			if v, ok := special[s[i:i+2]]; ok {
// 				res += v
// 				i++
// 			} else {
// 				res += roman[s[i]]
// 			}
// 		} else {
// 			res += roman[s[i]]
// 		}
// 	}
// 	if i != len(s) {
// 		res += roman[s[len(s)-1]]
// 	}

// 	return res
// }

/**********************************************/

// type band struct {
// 	i int    // integer
// 	n string // numeral
// }

// var bands = []band{
// 	{1000, "M"},
// 	{900, "CM"},
// 	{500, "D"},
// 	{400, "CD"},
// 	{100, "C"},
// 	{90, "XC"},
// 	{50, "L"},
// 	{40, "XL"},
// 	{10, "X"},
// 	{9, "IX"},
// 	{5, "V"},
// 	{4, "IV"},
// 	{1, "I"},
// }

// /*
// 官方题解：
// 3999/3999 cases passed (4 ms)
// Your runtime beats 94.89 % of golang submissions
// Your memory usage beats 40 % of golang submissions (3.1 MB)
// 按组合大小排列，从大到小的开始找，时间复杂度基本是O(N)，比每个值都遍历一遍效率高，空间O(1)，组合个数是固定的
// 不想定义新类型也可以改成2个数组成员的slice
// */
// func romanToInt(s string) int {
// 	var num int
// 	for _, b := range bands {
// 		for strings.HasPrefix(s, b.n) {
// 			num += b.i
// 			s = s[len(b.n):]
// 		}
// 	}
// 	return num
// }

/**********************************************/

/*
遍历每位，并预先判断下一位，如果下次比本次大则说明本位要-，最后一位直接+即可，时间O(n)，空间O(1)
3999/3999 cases passed (4 ms)
Your runtime beats 94.89 % of golang submissions
Your memory usage beats 100 % of golang submissions (3.1 MB)
*/
var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var num int
	for i := 0; i < len(s)-1; i++ {
		if roman[s[i+1]] > roman[s[i]] {
			num -= roman[s[i]]
		} else {
			num += roman[s[i]]
		}
	}
	num += roman[s[len(s)-1]]

	return num
}

// @lc code=end
