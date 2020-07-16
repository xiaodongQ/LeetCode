package lc

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 *
 * https://leetcode-cn.com/problems/happy-number/description/
 *
 * algorithms
 * Easy (60.27%)
 * Likes:    391
 * Dislikes: 0
 * Total Accepted:    87.3K
 * Total Submissions: 144.8K
 * Testcase Example:  '19'
 *
 * 编写一个算法来判断一个数 n 是不是快乐数。
 *
 * 「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到
 * 1。如果 可以变为  1，那么这个数就是快乐数。
 *
 * 如果 n 是快乐数就返回 True ；不是，则返回 False 。
 *
 *
 *
 * 示例：
 *
 * 输入：19
 * 输出：true
 * 解释：
 * 1^2 + 9^2 = 82
 * 8^2 + 2^2 = 68
 * 6^2 + 8^2 = 100
 * 1^2 + 0^2 + 0^2 = 1
 *
 *
 */

// @lc code=start
/*
	哈希表记录出现过的数，若出现循环则说明不满足条件
	401/401 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 50 % of golang submissions (2.2 MB)
	时间O(logn)，空间O(logn)，证明查看链接供参考

基于这个结论：不会出现值会越来越大，最后接近无穷大这种情况(分析可查看链接)：
	对于3位数的数字，它的各位平方和不可能大于243(即999的平方和)。
		这意味着它要么被困在 243 以下的循环内，要么跌到 1。
		4位或4位以上的数字在每一步都会丢失一位，直到降到3位为止
	最坏的情况下，算法可能会在 243 以下的所有数字上循环，然后回到它已经到过的一个循环或者回到 1。
		但它不会无限期地进行下去
*/
// func isHappy(n int) bool {
// 	// 访问不存在的key返回零值，此处即false
// 	m := make(map[int]bool)
// 	// 非1且未出现过的数则继续
// 	for n != 1 && !m[n] {
// 		v := squareSum(n)
// 		if v == 1 {
// 			return true
// 		}
// 		// 插入哈希表
// 		m[n] = true
// 		n = v
// 	}
// 	return n == 1
// }

/*
	快慢指针，快值等于慢值则说明有循环
	401/401 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
	时间O(logn)，空间O(logn)，证明比较复杂
*/
func isHappy(n int) bool {
	slow := n
	// 第一次就满足的情况也覆盖了
	fast := squareSum(n)
	for fast != 1 && slow != fast {
		slow = squareSum(slow)
		// 出现1则继续计算还是1，所以两次并没有问题
		fast = squareSum(squareSum(fast))
	}
	return fast == 1
}

/*
	数学方法，分析有结论之后很简单
	所有周期只有一个循环：
	4->16->37->58->89->145->42->20->44->16->37->58->89->145->42->20->4。
	所有其他数字都在进入这个循环的链上，或者在进入 1 的链上
	因此，可以硬编码一个包含这些数字的散列集，如果达到其中一个数字，那么就知道在循环中

	401/401 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
*/
// func isHappy(n int) bool {
// 	cycle := map[int]bool{4: true, 6: true, 37: true, 58: true, 89: true, 145: true, 42: true, 20: true}
// 	for n != 1 && !cycle[n] {
// 		n = squareSum(n)
// 	}
// 	return n == 1
// }

// 计算给定值的各位平方和
func squareSum(n int) int {
	res := 0
	bnum := 0
	for n > 0 {
		bnum = n % 10
		res += bnum * bnum
		n /= 10
	}
	return res
}

// @lc code=end
