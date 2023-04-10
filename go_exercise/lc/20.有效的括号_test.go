package lc

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (41.91%)
 * Likes:    1661
 * Dislikes: 0
 * Total Accepted:    317.7K
 * Total Submissions: 758K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	// 栈，依次遍历将左边入栈，LIFO弹出元素看是否与当前匹配
	stk := []rune{}
	var p = map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stk = append(stk, c)
		} else {
			// 从最后一个位置取元素当做出栈，若为空或与不对应右括号则不匹配
			if len(stk) == 0 || stk[len(stk)-1] != p[c] {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}

	// 最后栈里还有元素则说明不匹配
	return len(stk) == 0
}

// @lc code=end
