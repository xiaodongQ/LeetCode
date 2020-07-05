package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (34.95%)
 * Likes:    3918
 * Dislikes: 0
 * Total Accepted:    551K
 * Total Submissions: 1.6M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */
func TestLengthOfLongestSubstring(t *testing.T) {
	log.Printf("res:%d", lengthOfLongestSubstring("dvdf"))
}

// @lc code=start
/*
滑动窗口，
	987/987 cases passed (16 ms)
	Your runtime beats 31.81 % of golang submissions
	Your memory usage beats 67.74 % of golang submissions (2.8 MB)
	时间O(n)，每个字符遍历一遍
	空间O(m)，辅助map，m表示不同的字符个数(字符集大小，ASCII为0-128，可认为m最大128)
*/
// func lengthOfLongestSubstring(s string) int {
// 	// 记录每个字符出现的次数
// 	m := map[byte]int{}
// 	res := 0
// 	right := 0
// 	n := len(s)
// 	for i := 0; i < n; i++ {
// 		if i != 0 {
// 			delete(m, s[i-1])
// 		}
// 		for right < n && m[s[right]] == 0 {
// 			m[s[right]]++
// 			right++
// 		}
// 		if right-i > res {
// 			res = right - i
// 		}
// 	}

// 	return res
// }

/*
 国际站上投票最高的答案
 map记录每个字符的位置，遍历字符串，
 当map中有该字符时，取出其位置作为起始位置(模拟指针)，算出本次位置到起始的长度，和既有最大长度进行比较；
 本次操作后更新该字符的下标为新下标。
 注意，位置指针(保存的索引)只能往后移动

 提交：
	987/987 cases passed (8 ms)
	Your runtime beats 67.12 % of golang submissions
	Your memory usage beats 19.35 % of golang submissions (5.7 MB)
提前定义128大小的map，速度会快一些，空间换时间
不存在上面的map删除操作，效率也会高些
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 记录每个字符的下标
	m := make(map[byte]int, 128)
	// m := map[byte]int{}
	res := 0
	begin := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if pos, ok := m[s[i]]; ok {
			// 每次移动只能向后
			begin = maxInt(begin, pos+1)
		}
		m[s[i]] = i
		res = maxInt(res, i-begin+1)
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
