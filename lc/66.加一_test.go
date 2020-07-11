package lc

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (44.64%)
 * Likes:    496
 * Dislikes: 0
 * Total Accepted:    173.4K
 * Total Submissions: 388.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */

// @lc code=start
/*
	109/109 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 25 % of golang submissions (2.1 MB)
	时间O(n)，空间(n)
*/
// func plusOne(digits []int) []int {
// 	if len(digits) == 0 {
// 		return nil
// 	}

// 	// 多分配一位
// 	res := make([]int, len(digits)+1)
// 	pos := len(res) - 1
// 	plus := 0
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		if i == len(digits)-1 {
// 			res[pos] = (digits[i] + 1) % 10
// 			plus = (digits[i] + 1) / 10
// 		} else {
// 			res[pos] = (digits[i] + plus) % 10
// 			plus = (digits[i] + plus) / 10
// 		}
// 		pos--
// 	}
// 	if plus == 1 {
// 		res[0] = 1
// 	} else {
// 		res = res[1:]
// 	}
// 	return res
// }

/*
	预先分配 和 每个数都计算除法和余数，多了很多不必要的操作
	碰到第一次<9的位，加1，然后返回即可；
	=9的位则赋0后进一位，继续判断下一位，若下一位<9则也是+1后返回
	最后遍历完还没有返回，则说明最高位进位新增了一位

	109/109 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
	时间O(n)(但是比上面执行操作少)，空间O(1) (除了都是9的情况创建一个slice)
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		// 说明此位原来是9，+1变成0后继续下一位判断
		digits[i] = 0
	}
	// 最后还没return说明所有位都是9，并新增一位
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}

// @lc code=end
