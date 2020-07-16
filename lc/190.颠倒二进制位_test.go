package lc

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 *
 * https://leetcode-cn.com/problems/reverse-bits/description/
 *
 * algorithms
 * Easy (59.82%)
 * Likes:    184
 * Dislikes: 0
 * Total Accepted:    46.7K
 * Total Submissions: 78K
 * Testcase Example:  '00000010100101000001111010011100'
 *
 * 颠倒给定的 32 位无符号整数的二进制位。
 *
 *
 *
 * 示例 1：
 *
 * 输入: 00000010100101000001111010011100
 * 输出: 00111001011110000010100101000000
 * 解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
 * ⁠    因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
 *
 * 示例 2：
 *
 * 输入：11111111111111111111111111111101
 * 输出：10111111111111111111111111111111
 * 解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
 * 因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。
 *
 *
 *
 * 提示：
 *
 *
 * 请注意，在某些语言（如
 * Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
 * 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数
 * -1073741825。
 *
 *
 *
 *
 * 进阶:
 * 如果多次调用这个函数，你将如何优化你的算法？
 *
 */

// @lc code=start
/*
	用strconv库转换为字符串再反转，有点蹩脚。。。
	600/600 cases passed (4 ms)
	Your runtime beats 60.23 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.6 MB)
*/
// func reverseBits(num uint32) uint32 {
// 	s := make([]byte, 32)
// 	s1 := []byte(strconv.FormatUint(uint64(num), 2))
// 	// 高位0不显示，凑成32位
// 	zeros := make([]byte, 32-len(s1))
// 	for i := 0; i < len(zeros); i++ {
// 		zeros[i] = '0'
// 	}
// 	copy(s, append(zeros, s1...))
// 	for i := 0; i < len(s)/2; i++ {
// 		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
// 	}
// 	res, _ := strconv.ParseUint(string(s), 2, 32)
// 	return uint32(res)
// }

/*
	直接每位操作
	600/600 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.5 MB)
	时间O(1) 固定操作32位，可以看做log(N) N为位数，空间O(1)
*/
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		// 也可以低位直接放到最终位置，res += ((num&0x01)<<(31-i))，这种方式i从31开始比较简洁一点
		res = res<<1 + num&0x01
		num >>= 1
	}
	return res
}

/*
	若要求不用循环，可以用 189.旋转数组 中的方式，交换两部分数据，再拆分数据直到间隔为1
	两个16位交换 -> 两部分分别操作：两个8位交换

	600/600 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.5 MB)
	时间O(1)，空间O(1)
*/
// func reverseBits(num uint32) uint32 {
// 	num = (num >> 16) | (num << 16)
// 	// 每8位交换
// 	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
// 	// 每4位
// 	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
// 	// 每2位 1100:c; 0011:3
// 	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
// 	// 每1位 1010:a; 0101:5
// 	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)

// 	return num
// }

// @lc code=end
