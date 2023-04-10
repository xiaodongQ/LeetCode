package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 *
 * https://leetcode-cn.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (39.92%)
 * Likes:    310
 * Dislikes: 0
 * Total Accepted:    41.1K
 * Total Submissions: 103.1K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，返回 n! 结果尾数中零的数量。
 *
 * 示例 1:
 *
 * 输入: 3
 * 输出: 0
 * 解释: 3! = 6, 尾数中没有零。
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: 1
 * 解释: 5! = 120, 尾数中有 1 个零.
 *
 * 说明: 你算法的时间复杂度应为 O(log n) 。
 *
 */
func TestTrailingZeroes(t *testing.T) {
	log.Printf("res:%d", trailingZeroes(30))
}

// @lc code=start
/*
	因数中包含5的个数就是0的个数，5：1个0，25：2个0，50：2个

方法正确，提交失败：
	超时了，而且实现得不大简洁，时间O(nlogn)，空间O(logn)
	Time Limit Exceeded
	500/502 cases passed (N/A)

	Testcase
		1808548329
*/
// func trailingZeroes(n int) int {
// 	res := 0
// 	// 换底公式算log5(n)
// 	mNum := make(map[int]int, int(math.Log10(float64(n))/math.Log10(5)))
// 	for i := 5; i <= n; i += 5 {
// 		if i%5 != 0 {
// 			continue
// 		}
// 		// 算该值里包含的0
// 		num := 0
// 		for n := i; n > 0 && n%5 == 0; n /= 5 {
// 			if nn, ok := mNum[n]; ok {
// 				// 已经有这个值包含的5个数了，则叠加后不用再继续重复计算
// 				num += nn
// 				break
// 			} else {
// 				num++
// 			}
// 		}
// 		mNum[i] = num

// 		res += num
// 	}
// 	return res
// }

/*
方法错误：
	包含个数，循环每个值，求log5(i)，借助标准库
*/
// func trailingZeroes(n int) int {
// 	res := 0
// 	for i := 1; i <= n; i++ {
// 		if i%5 != 0 {
// 			continue
// 		}
// 		// log5(i)换底
// 		// 不能这样算，30只包含一个5，这样会算出2个
// 		num := int(math.Log2(float64(i)) / math.Log2(5))
// 		if math.Pow(5, float64(num)) > float64(i) {
// 			num--
// 		}
// 		res += num
// 	}

// 	return res
// }

/*
	包含5的个数+1，包含25个数+2，则可以 /5得到的个数+/25得到的个数，依次类推
		(同一个25在/5时+1次，并在/25时再+1次)
	证明可参考：[Trailing Number of Zeros](https://brilliant.org/wiki/trailing-number-of-zeros/)

	502/502 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
	时间O(logn)，空间O(1)
*/
func trailingZeroes(n int) int {
	res := 0
	i := 5
	for i <= n {
		res += n / i
		i *= 5
	}

	return res
}

// @lc code=end
