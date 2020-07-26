package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 *
 * https://leetcode-cn.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (46.93%)
 * Likes:    119
 * Dislikes: 0
 * Total Accepted:    51.8K
 * Total Submissions: 110.5K
 * Testcase Example:  '27'
 *
 * 给定一个整数，写一个函数来判断它是否是 3 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 27
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: true
 *
 * 示例 4:
 *
 * 输入: 45
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */

// @lc code=start
func TestIsPowerOfThree(t *testing.T) {
	log.Printf("res:%t", isPowerOfThree(81))
}

/*
	21038/21038 cases passed (44 ms)
	Your runtime beats 20.64 % of golang submissions
	Your memory usage beats 20 % of golang submissions (6.1 MB)
	时间O(logn)，空间O(1)
*/
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}

	return n == 1
}

/*
失败了，21027/21038 cases passed (N/A)
	部分值还是不满足精度，81，math.Abs(diff) <= 10*epsilon也不行

	不用循环
	利用库函数，判断log3(n)是否为整数，由于没有提供3为底的log，利用换底公式转换
*/
// func isPowerOfThree(n int) bool {
// 	// strconv.FormatInt(n, 3)
// 	l := math.Log10(float64(n)) / math.Log10(3)
// 	log.Printf("l:%f", l)
// 	// 去掉小数位然后比较
// 	m := int(l)
// 	// l:3.0000000000000004, float64(m):3，==为false
// 	log.Printf("l:%v, m:%v", l, float64(m))
// 	// 不行，float精度问题
// 	// return l == float64(m)
// 	diff := l - float64(m)
// 	// 1.0~2.0之间下一个float64，-1.0则为float64最小精度
// 	epsilon := math.Nextafter(1.0, 2.0) - 1
// 	log.Printf("diff:%f, e:%f, %v, res:%t", diff, epsilon, math.Abs(diff), math.Abs(diff) <= 10*epsilon)

// 	// return l == math.Floor(l) //也不行
// 	return math.Abs(diff) <= 2*epsilon
// }

/*
	进制转换，对转换后的字符串进行正则匹配
	3进制，则 8可表示为 22(即2 * 3^1 + 2 * 3^0 = 8)，如果是3的指数，则形式为1000形式(e.g. 9对应3进制为: 100)
	利用正则表达式匹配转换进制后的字符串
	时间复杂度主要是库函数的实现复杂度(通常是用除法，类似上面的O(logn))，空间O(1)
*/
// func isPowerOfThree(n int) bool {
// 	str := strconv.FormatInt(int64(n), 3)
// 	ok, err := regexp.MatchString("^10*$", str)
// 	if err != nil || !ok {
// 		return false
// 	}
// 	return true
// }

/*
	直接用最大的3的指数，对当前值进行除法，能整除则为true
		3为质数，3^n的因数除了1之外(1也满足true)，也只会是3的指数(3^0，3^1，3^2...)
	n为int，64位系统上为64位，2^63部分表示正数
	不过输入整数此处是当作32位了，2^31内最大为3^19 = 1162261467
*/
// func isPowerOfThree(n int) bool {
// 	return n > 0 && 1162261467%n == 0
// }

// @lc code=end
