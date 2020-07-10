package lc

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (39.73%)
 * Likes:    495
 * Dislikes: 0
 * Total Accepted:    194.3K
 * Total Submissions: 489.3K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */

// @lc code=start
/*
	74/74 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 60 % of golang submissions (2.3 MB)
	时间O((m-n)*n)，所以大致为O(mn)，空间O(1)，按字符比较，不用语言特性
*/
// func strStr(haystack string, needle string) int {
// 	// return strings.Index(haystack, needle)
// 	if len(needle) == 0 {
// 		return 0
// 	}
// 	if len(needle) > len(haystack) {
// 		return -1
// 	}

// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {
// 		j := 0
// 		for ; j < len(needle); j++ {
// 			if needle[j] != haystack[i+j] {
// 				break
// 			}
// 		}
// 		if j == len(needle) {
// 			return i
// 		}
// 	}
// 	return -1
// }

/*
	利用slice的截取子串
	74/74 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 60 % of golang submissions (2.3 MB)
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	length := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
