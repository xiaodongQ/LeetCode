package lc

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (34.33%)
 * Likes:    388
 * Dislikes: 0
 * Total Accepted:    65.6K
 * Total Submissions: 191K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */

// @lc code=start
/*
	埃氏筛，
		保存n大小的数组并全部标记为true(方便起见，2对应索引2，算个数时也从2开始即可)
		从2开始(假设本次为i)，i的倍数的数说明有两个以上因数，即不为质数，将数组对应值标记为false
		最后获取数组为true的元素个数，即<n的质数个数
	也叫埃拉托斯特尼筛法 / 厄拉多塞筛法（sieve of Eratosthenes），简称埃氏筛，是一种简单且年代久远的算法，用来找出一定范围内所有的质数
	20/20 cases passed (8 ms)
	Your runtime beats 95.92 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.9 MB)
	时间复杂度：O(nloglogn)，证明就不看了。。。
*/
func countPrimes(n int) int {
	res := 0
	// 由于默认零值为false，那就以值为false时当作满足质数，一开始所有都当作质数
	// 质数(prime) 和 合数(composite)，定义一个有意义的名字
	isComposite := make([]bool, n)
	// 只需要遍历到sqrt(n)，i比其小时会乘到比其大的数，比其大相当于重复判断了一遍
	for i := 2; i*i < n; i++ {
		if isComposite[i] {
			// i不为质数则不需要继续标记其倍数
			continue
		}
		// j为i依次递增的倍数，假设此时i=2，则2*3, 和后续i=3时，3*2会重复，所以每次直接从i*i开始标记
		for j := i * i; j < n; j += i {
			isComposite[j] = true
		}
	}
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			res++
		}
	}
	return res
}

// @lc code=end
