package lc

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (54.27%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    112.4K
 * Total Submissions: 207K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 *
 * 输入为 非空 字符串且只包含数字 1 和 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 *
 *
 * 提示：
 *
 *
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 *
 *
 */

// @lc code=start
/*
	294/294 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.2 MB)
*/
// func addBinary(a string, b string) string {
// 	maxlen := 0
// 	if len(a) > len(b) {
// 		maxlen = len(a)
// 	} else {
// 		maxlen = len(b)
// 	}
// 	res := make([]byte, maxlen+1)
// 	pos := len(res) - 1

// 	i := len(a) - 1
// 	j := len(b) - 1
// 	inc := 0
// 	ia := 0
// 	ib := 0
// 	for i >= 0 && j >= 0 {
// 		ia = int(a[i] - '0')
// 		ib = int(b[j] - '0')
// 		res[pos] = byte((ia ^ ib ^ inc) + '0')
// 		if ia+ib+inc >= 2 {
// 			inc = 1
// 		} else {
// 			inc = 0
// 		}
// 		pos--
// 		i--
// 		j--
// 	}
// 	for i != -1 {
// 		ia = int(a[i] - '0')
// 		res[pos] = byte((ia ^ inc) + '0')
// 		if ia+inc >= 2 {
// 			inc = 1
// 		} else {
// 			inc = 0
// 		}
// 		pos--
// 		i--
// 	}
// 	for j != -1 {
// 		ib = int(b[j] - '0')
// 		res[pos] = byte((ib ^ inc) + '0')
// 		if ib+inc >= 2 {
// 			inc = 1
// 		} else {
// 			inc = 0
// 		}
// 		pos--
// 		j--
// 	}
// 	if inc == 1 {
// 		res[0] = '1'
// 		return string(res[:])
// 	}
// 	return string(res[1:])
// }

/*
	上面的方式不简洁
	94/294 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.2 MB)
	时间O(max(m,n))，空间O(max(m,n))，若除去结果的空间则空间O(1)
*/
func addBinary(a string, b string) string {
	maxlen := len(a)
	if len(b) > maxlen {
		maxlen = len(b)
	}
	res := make([]byte, maxlen+1)
	i := len(a) - 1
	j := len(b) - 1
	pos := len(res) - 1
	carry := 0
	// 没有的位当作是0也不影响，相当于补0
	// 对应两个数组/链表的处理，用||还是&&并不都一样，
	// 合并链表时&&之后，剩余的做依次指向即可
	// 合并有序数组时，从后往前，||貌似还是更简洁一点
	for i >= 0 || j >= 0 {
		var ia, ib = 0, 0
		if i >= 0 {
			ia = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			ib = int(b[j] - '0')
			j--
		}
		res[pos] = byte(ia ^ ib ^ carry + '0')
		carry = (ia + ib + carry) >> 1
		pos--
	}
	if carry == 1 {
		res[0] = '1'
		return string(res)
	}
	return string(res[1:])
}

// @lc code=end
