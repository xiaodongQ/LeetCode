package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (38.37%)
 * Likes:    1131
 * Dislikes: 0
 * Total Accepted:    302.9K
 * Total Submissions: 789.6K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */
func TestLongestCommonPrefix(t *testing.T) {
	log.Printf("res:%s", longestCommonPrefix([]string{
		"aca", "cba",
	}))
}

// @lc code=start
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	// 找最短的字符串，从大到小依次找组合，组合存到一个map里
// 	res := ""
// 	minIndex := 0
// 	minlen := len(strs[0])
// 	for i, v := range strs {
// 		if len(v) < minlen {
// 			minIndex = i
// 			minlen = len(v)
// 		}
// 	}
// 	minstr := strs[minIndex]

// 	log.Printf("minstr:%s, minlen:%d", minstr, minlen)
// 	// 值为长度
// 	m := map[string]int{}

// 	for l := minlen; l > 0; l-- {
// 		for i := 0; i < minlen-l+1; i++ {
// 			prefix := minstr[i : i+l]
// 			if _, ok := m[prefix]; ok {
// 				continue
// 			}
// 			m[prefix] = 0

// 			num := 0
// 			for j := 0; j < len(strs); j++ {
// 				if j == minIndex {
// 					continue
// 				}
// 				if !strings.HasPrefix(strs[j], prefix) {
// 					break
// 				}
// 				num++
// 			}

// 			if num == len(strs)-1 {
// 				return prefix
// 			}
// 		}
// 	}
// 	return res
// }

/*
	118/118 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 82.35 % of golang submissions (2.4 MB)
	时间O(n*m)，n为字符串个数，m为最小串长度，空间O(1)
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 找最短的字符串，题意是只要判断前缀，而不需要判断中间包含相同串的情况
	res := ""
	minIndex := 0
	minlen := len(strs[0])
	for i, v := range strs {
		if len(v) < minlen {
			minIndex = i
			minlen = len(v)
		}
	}
	minstr := strs[minIndex]

	for l := minlen; l > 0; l-- {
		num := 0
		for j := 0; j < len(strs); j++ {
			if j == minIndex {
				continue
			}
			if strs[j][:l] == minstr[:l] {
				num++
			}
		}
		if num == len(strs)-1 {
			return minstr[:l]
		}
	}
	return res
}

// @lc code=end
