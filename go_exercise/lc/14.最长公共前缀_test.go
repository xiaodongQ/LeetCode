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
		"aa", "a",
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
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	// 找最短的字符串，题意是只要判断前缀，而不需要判断中间包含相同串的情况
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

// 	// 从最长串开始，每个字符串进行匹配
// 	for l := minlen; l > 0; l-- {
// 		num := 0
// 		for j := 0; j < len(strs); j++ {
// 			if j == minIndex {
// 				continue
// 			}
// 			// 每个字符串的前几位比较
// 			if strs[j][:l] == minstr[:l] {
// 				num++
// 			}
// 		}
// 		if num == len(strs)-1 {
// 			return minstr[:l]
// 		}
// 	}
// 	return res
// }

/*
	从头开始同时依次比较字符，直到有字符串结束或者不同
	118/118 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 23.53 % of golang submissions (2.4 MB)
*/
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	// 其实只要最短长度即可，不过初值不能赋0
// 	maxLen := 0
// 	for _, s := range strs {
// 		if len(s) > maxLen {
// 			maxLen = len(s)
// 		}
// 	}

// 	c := []byte{}
// 	for i := 0; i < maxLen; i++ {
// 		bExist := false
// 		num := 0
// 		var curChar byte
// 		for _, s := range strs {
// 			if !bExist && i < len(s) {
// 				curChar = s[i]
// 				bExist = true
// 			}
// 			if i > len(s)-1 || s[i] != curChar {
// 				return string(c)
// 			}
// 			num++
// 		}
// 		if num == len(strs) {
// 			c = append(c, curChar)
// 		}
// 	}

// 	return string(c)
// }

/*
	参考题解。和上面的思路一样，不过上面逻辑比较晦涩，
	出现长度不足的字符串说明要结束了，不需要上面一些小操作
	时间O(m*n)，m为最小串长度，n为字符串个数
	118/118 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 82.35 % of golang submissions (2.4 MB)
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 随便选一个字符串作为基准，依次比较每位
	for i := 0; i < len(strs[0]); i++ {
		// for _, s := range strs {
		// 比较时除去基准字符串
		for j := 1; j < len(strs); j++ {
			// 注意下标的终止条件
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	// 最后遍历完了还是一样，则说明最长前缀和基准字符串一样
	return strs[0]
}

// @lc code=end
