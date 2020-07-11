package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (33.42%)
 * Likes:    213
 * Dislikes: 0
 * Total Accepted:    103.9K
 * Total Submissions: 311K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
 *
 *
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func TestLengthOfLastWord(t *testing.T) {
	log.Printf("res:%d", lengthOfLastWord(" a b df "))
}

// @lc code=start
/*
	使用strings包的api，会增加时间和空间复杂度
	59/59 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 25 % of golang submissions (2.2 MB)
*/
// func lengthOfLastWord(s string) int {
// 	if s == "" {
// 		return 0
// 	}
// 	// " a "会把最后和最开始的字符也当作一个slice的元素，要先trim
// 	s1 := strings.TrimLeft(s, " ")
// 	s1 = strings.TrimRight(s1, " ")
// 	sub := strings.Split(s1, " ")
// 	return len(sub[len(sub)-1])
// }

/*
	不使用api，时间O(n)，空间O(1)，比strings包要省空间
	59/59 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 25 % of golang submissions (2.1 MB)
*/
func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	// 从最后开始遍历，碰到' '则结束，需要排除后面都是' '的情况
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {

	}
	res := 0
	for ; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		res++
	}

	return res
}

// @lc code=end
