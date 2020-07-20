package lc

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 *
 * https://leetcode-cn.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (67.13%)
 * Likes:    260
 * Dislikes: 0
 * Total Accepted:    43.6K
 * Total Submissions: 65K
 * Testcase Example:  '38'
 *
 * 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。
 *
 * 示例:
 *
 * 输入: 38
 * 输出: 2
 * 解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
 *
 *
 * 进阶:
 * 你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？
 *
 */

// @lc code=start
/*
	循环
	1101/1101 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.2 MB)
	时间O(nlogn)，空间O(1)
*/
func addDigits(num int) int {
	for num > 9 {
		// 函数调用更清晰一点，不过多了函数入栈出栈操作
		// num = sumBit(num)
		bitSum := 0
		tmp := num
		for tmp > 0 {
			bitSum += tmp % 10
			tmp = tmp / 10
		}
		num = bitSum
	}
	return num
}

func sumBit(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	return res
}

/*
	数学分析：[解法二 数学上](https://leetcode-cn.com/problems/add-digits/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-5-7/)
	数根(又称位数根或数字根Digital root)
	在数学中，数根(又称位数根或数字根Digital root)是自然数的一种性质，换句话说，每个自然数都有一个数根。
		数根是将一正整数的各个位数相加（即横向相加），
		若加完后的值大于10的话，则继续将各位数进行横向相加直到其值小于十为止，
		或是，将一数字重复做数字和，直到其值小于十为止，则所得的值为该数的数根
	用途：
		数根可以计算模运算的同余，对于非常大的数字的情况下可以节省很多时间
		数字根可作为一种检验计算正确性的方法。例如，两数字的和的数根等于两数字分别的数根的和
		另外，数根也可以用来判断数字的整除性，如果数根能被3或9整除，则原来的数也能被3或9整除
	对于给定的 n 有三种情况:
		n 是 0 ，数根就是 0
		n 不是 9 的倍数，数根就是 n 对 9 取余，即 n mod 9
		n 是 9 的倍数，数根就是 9
	原数是 n，树根就可以表示成 (n-1) mod 9 + 1

	1101/1101 cases passed (4 ms)
	Your runtime beats 42.33 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.2 MB)
	时间O(1)，空间O(1)
	运行的时间比上面还久跟用例有关？
*/
// func addDigits(num int) int {
// 	return (num-1)%9 + 1
// }

// @lc code=end
