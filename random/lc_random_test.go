package random

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
