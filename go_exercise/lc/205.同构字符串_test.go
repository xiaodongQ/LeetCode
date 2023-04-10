package lc

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (47.22%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    44.1K
 * Total Submissions: 93.5K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *
 * 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 *
 * 示例 1:
 *
 * 输入: s = "egg", t = "add"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "foo", t = "bar"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: s = "paper", t = "title"
 * 输出: true
 *
 * 说明:
 * 你可以假设 s 和 t 具有相同的长度。
 *
 */

// @lc code=start
/*
反例：提交后各种情形不满足，调整后又出现新的情形不满足，靠巧合编程。。。
	map记录替换的内容
*/
// func isIsomorphic(s string, t string) bool {
// 	m := make(map[byte]byte)
// 	// 记录被谁哈希
// 	reversem := make(map[byte]byte)
// 	var sbyte byte
// 	for i := 0; i < len(s); i++ {
// 		sbyte = s[i]
// 		if r, ok := m[s[i]]; ok {
// 			sbyte = r
// 		}
// 		if sbyte != t[i] {
// 			// 记录这个字符已经替换为本身或新字符了
// 			if replace, ok := m[sbyte]; !ok {
// 				// 记录之前，需要查询之前是否已经有字符占用了要替换成的字符
// 				// 因为要求两个字符不能哈希到同一个字符
// 				if replace2, ok2 := reversem[t[i]]; ok2 {
// 					if replace2 != sbyte {
// 						return false
// 					}
// 				}

// 				// 还要判断两个数不能互为哈希
// 				if r3, ok3 := m[t[i]]; ok3 {
// 					if r3 == sbyte {
// 						return false
// 					}
// 				}

// 				m[sbyte] = t[i]
// 				reversem[t[i]] = sbyte
// 			} else {
// 				if replace != t[i] {
// 					return false
// 				}
// 			}
// 		} else {
// 			// 相同的位也记录map
// 			m[sbyte] = sbyte
// 			reversem[sbyte] = sbyte
// 		}
// 	}

// 	return true
// }

/*
解法一：
	双向都满足替换关系，判断两次(上面把问题混一起比较难处理)
	32/32 cases passed (4 ms)
	Your runtime beats 76.63 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.7 MB)
	时间O(n)，空间O(n)

解法二：
	另外一种解法：每个字符串都映射成其他字符组成的字符串
	"ab" -> "12"
	"cd" -> "12"

	"abbad" -> "12213"
	"cddfa" -> "12234"

	"abbac" -> "12213"
	"bccbd" -> "12213"
*/
func isIsomorphic(s string, t string) bool {
	return omorphicHelper(s, t) && omorphicHelper(t, s)
}

func omorphicHelper(s, t string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		// 判断方式有问题，先判断了原值
		// if s[i] != t[i] {
		if replace, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		} else {
			if replace != t[i] {
				return false
			}
		}
	}

	return true
}

// @lc code=end
