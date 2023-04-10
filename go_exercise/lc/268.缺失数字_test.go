package lc

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 *
 * https://leetcode-cn.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (56.25%)
 * Likes:    286
 * Dislikes: 0
 * Total Accepted:    74.6K
 * Total Submissions: 132.7K
 * Testcase Example:  '[3,0,1]'
 *
 * 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,0,1]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [9,6,4,2,3,5,7,0,1]
 * 输出: 8
 *
 *
 *
 *
 * 说明:
 * 你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
 *
 */

// @lc code=start
/*
	122/122 cases passed (24 ms)
	Your runtime beats 37.78 % of golang submissions
	Your memory usage beats 50 % of golang submissions (6 MB)
	时间O(n)，空间O(n)
*/
// func missingNumber(nums []int) int {
// 	m := make(map[int]bool, len(nums))
// 	// n个数，传入2个数则0,1,2范围
// 	max := len(nums)
// 	for _, v := range nums {
// 		m[v] = true
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	for i := 0; i <= max; i++ {
// 		if _, ok := m[i]; !ok {
// 			return i
// 		}
// 	}
// 	return 0
// }

/*
	没必要用另外空间，i从0~n相加，减去数组各元素，剩下的就是
	122/122 cases passed (20 ms)
	Your runtime beats 81.95 % of golang submissions
	Your memory usage beats 50 % of golang submissions (6 MB)
	时间O(n)，空间O(1)
*/
func missingNumber(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += i + 1
		sum -= nums[i]
	}
	return sum
}

/*
	利用异或操作，0~n依次和数组元素异或，最后剩下的就是缺少的
	122/122 cases passed (20 ms)
	Your runtime beats 81.95 % of golang submissions
	Your memory usage beats 50 % of golang submissions (6 MB)
	时间O(n)，空间O(1)
*/
// func missingNumber(nums []int) int {
// 	res := 0
// 	for i := 0; i < len(nums); i++ {
// 		res = res ^ (i + 1)
// 		res = res ^ nums[i]
// 	}
// 	return res
// }

// @lc code=end
