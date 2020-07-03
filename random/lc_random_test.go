package random

/*
 * 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部
 * [剑指 Offer 58 - II. 左旋转字符串](https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/)
 * e.g. 输入: s = "abcdefg", k = 2
	输出: "cdefgab"
 * 限制：1 <= k < s.length <= 10000
*/
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
