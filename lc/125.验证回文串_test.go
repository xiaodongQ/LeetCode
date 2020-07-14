package lc

import (
	"testing"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (45.80%)
 * Likes:    246
 * Dislikes: 0
 * Total Accepted:    140.4K
 * Total Submissions: 306.6K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */
func TestIsPalindrome(t *testing.T) {
	// log.Printf("res:%t", isPalindrome("A man, a plan, a canal: Panama"))
	// isPalindrome 和 本包中的函数冲突，自测时改下名字
	// log.Printf("res:%t", isPalindrome1("A man, a plan, a canal: Panama"))
}

// @lc code=start
/*
	暴力：用一个哈希表遍历存储字母
	双指针
*/

/*
 ~~~使用哈希超时了~~~ 并不是哈希问题，是自己判断操作下标问题
	481/481 cases passed (8 ms)
	Your runtime beats 20.14 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.7 MB)
	时间O(n)，空间固定，所以为O(1)
*/
// (26字母*2+10)
// var charSet = make(map[byte]byte, 64)

// func isPalindrome(s string) bool {
// 	initChar()
// 	return palinhelper(s)
// }
// func palinhelper(s string) bool {
// 	l, r := 0, len(s)-1
// 	for l <= r {
// 		lc, ok1 := charSet[s[l]]
// 		if !ok1 {
// 			l++
// 			continue
// 		}
// 		rc, ok2 := charSet[s[r]]
// 		if !ok2 {
// 			r--
// 			continue
// 		}
// 		if lc != rc {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
// 	return true
// }
// func initChar() {
// 	var c byte
// 	for c = 'a'; c <= 'z'; c++ {
// 		charSet[c] = c
// 	}
// 	for c = 'A'; c <= 'Z'; c++ {
// 		charSet[c] = c + 32
// 	}
// 	for c = '0'; c <= '9'; c++ {
// 		charSet[c] = c
// 	}
// }

/*
双指针
	481/481 cases passed (4 ms)
	Your runtime beats 77 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.7 MB)
	时间O(n)，空间O(1)
*/
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		lc, ok1 := valid(s[l])
		if !ok1 {
			l++
			continue
		}
		rc, ok2 := valid(s[r])
		if !ok2 {
			r--
			continue
		}
		if lc != rc {
			return false
		}
		l++
		r--
	}
	return true
}

// 按 comma,ok 的习惯
func valid(b byte) (byte, bool) {
	if (b >= 'a' && b <= 'z') ||
		(b >= '0' && b <= '9') {
		return b, true
	}
	if b >= 'A' && b <= 'Z' {
		return b + 32, true
	}

	return b, false
}

/*
	481/481 cases passed (4 ms)
	Your runtime beats 77 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3.4 MB)
	时间O(n)，空间O(n)
*/
// func isPalindrome(s string) bool {
// 	news := make([]byte, len(s))
// 	pos := 0
// 	for i := 0; i < len(s); i++ {
// 		has, b := valid(s[i])
// 		if has {
// 			news[pos] = b
// 			pos++
// 		}
// 	}
// 	// 清理后面的无用字符
// 	news = news[:pos]
// 	log.Printf("news:%s\n", news)
// 	for i := 0; i < len(news)/2; i++ {
// 		if news[i] != news[len(news)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
// func valid(b byte) (bool, byte) {
// 	if (b >= 'a' && b <= 'z') ||
// 		(b >= '0' && b <= '9') {
// 		return true, b
// 	}
// 	if b >= 'A' && b <= 'Z' {
// 		return true, b + 32
// 	}

// 	return false, b
// }

// @lc code=end
