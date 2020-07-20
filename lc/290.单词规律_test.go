package lc

import (
	"log"
	"strings"
	"testing"
)

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 *
 * https://leetcode-cn.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (43.02%)
 * Likes:    171
 * Dislikes: 0
 * Total Accepted:    27.3K
 * Total Submissions: 63.5K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
 *
 * 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
 *
 * 示例1:
 *
 * 输入: pattern = "abba", str = "dog cat cat dog"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:pattern = "abba", str = "dog cat cat fish"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: pattern = "aaaa", str = "dog cat cat dog"
 * 输出: false
 *
 * 示例 4:
 *
 * 输入: pattern = "abba", str = "dog dog dog dog"
 * 输出: false
 *
 * 说明:
 * 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 *
 */
func TestWordPattern(t *testing.T) {
	log.Printf("res:%v", wordPattern("abba", "dog cat cat dog"))
}

// @lc code=start
/*
	两个串都映射成同一个第三方集合里
	37/37 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
*/
func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	m := make(map[string]byte, len(s))
	var ch byte = 'a'
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		var curCh byte
		if c, ok := m[s[i]]; !ok {
			m[s[i]] = ch
			curCh = ch
			ch++
		} else {
			curCh = c
		}
		sb.WriteByte(curCh)
	}

	var sbpat strings.Builder
	mpat := make(map[byte]byte, 0)
	ch = 'a'
	for i := 0; i < len(pattern); i++ {
		var curCh byte
		if c, ok := mpat[pattern[i]]; !ok {
			mpat[pattern[i]] = ch
			curCh = ch
			ch++
		} else {
			curCh = c
		}
		sbpat.WriteByte(curCh)
	}

	return sb.String() == sbpat.String()
}

// @lc code=end
