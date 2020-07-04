package random

import (
	"strings"
)

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

/*
* [LCP 01. 猜数字](https://leetcode-cn.com/problems/guess-numbers/)
 */
func game(guess []int, answer []int) int {
	n := 0
	for k, g := range guess {
		if g == answer[k] {
			n++
		}
	}
	return n
}

/*
* [771. 宝石与石头](https://leetcode-cn.com/problems/jewels-and-stones/)
 */
func numJewelsInStones(J string, S string) int {
	/*
		//时间O(J*S)，空间O(1)
		n := 0
		for _, s := range S {
			for _, j := range J {
				if s == j {
					n++
					break
				}
			}
		}
		return n
	*/
	// 时间O(J+S)，空间O(J)
	n := 0
	jm := make(map[rune]bool)
	for _, j := range J {
		jm[j] = true
	}
	for _, s := range S {
		if _, ok := jm[s]; ok {
			n++
		}
	}
	return n
}

/*
* [1281. 整数的各位积和之差](https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/)
 */
func subtractProductAndSum(n int) int {
	// 时间O(logn)，空间O(1)
	// 对于时间复杂度：其中执行次数为 1+log10(n)，根据换底公式 1+log2(n) / log2(10)，即 Clogn + 1，因此复杂度为O(logn)
	sum := 0
	product := 1
	temp := 0
	for n != 0 {
		temp = n % 10
		sum += temp
		product *= temp
		n = n / 10
	}
	return product - sum
}

/*
* [1108. IP 地址无效化](https://leetcode-cn.com/problems/defanging-an-ip-address/)
 */
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

/*
* 给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1
* [1342. 将数字变成 0 的操作次数](https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/)
 */
func numberOfSteps(num int) int {
	/*
		count := 0
		for num != 0 {
			// 偶数
			if (num & 1) != 1 {
				num = num >> 1
			} else {
				num--
			}
			count++
		}
		return count
	*/

	/*
		其他思路(时间复杂度O(logn)，循环次数为其位数)：
		从右到左遍历二进制位，碰到0需要除2，所以次数+1；碰到1则减1后再除2，所以次数+2；
		最后一位肯定为1，只需要除2即可将其变为0，所以按上面方法算的次数减1即为总次数
		来源：[国际站讨论](https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/discuss/502809/just-count-number-of-0-and-1-in-binary)
		其中还包括C++ O(1)的实现，借助内建函数实现：return num ? __builtin_popcount(num) + 31 - __builtin_clz(num) : 0;
	*/
	if num == 0 {
		return 0
	}
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count = count + 2
		} else {
			count++
		}
		num = num >> 1
	}
	return count - 1
}
