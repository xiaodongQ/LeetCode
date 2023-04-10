package lc

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (69.16%)
 * Likes:    1370
 * Dislikes: 0
 * Total Accepted:    236.2K
 * Total Submissions: 341.6K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,1]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: [4,1,2,1,2]
 * 输出: 4
 *
 */

// @lc code=start
/*
	哈希表保存遍历过的字符数量，>1则删除哈希表元素(删除O(1))，时间O(n)，空间O(n)
	m := make(map[int]int) 删节点
	16/16 cases passed (12 ms)
	Your runtime beats 76.42 % of golang submissions
	Your memory usage beats 16.67 % of golang submissions (5.9 MB)

	m := make(map[int]int, len(nums)) 删节点
	16/16 cases passed (20 ms)
	Your runtime beats 7.37 % of golang submissions
	Your memory usage beats 8.33 % of golang submissions (5.9 MB)

	m := make(map[int]int, len(nums)) 且不删节点
	16/16 cases passed (8 ms)
	Your runtime beats 98.43 % of golang submissions
	Your memory usage beats 8.33 % of golang submissions (5.9 MB)
*/
// func singleNumber(nums []int) int {
// 	m := make(map[int]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]]++
// 		// if m[nums[i]] > 1 {
// 		// 	// 题目中说了其他都为2次，才允许删除，若有更多次则不行
// 		// 	delete(m, nums[i])
// 		// }
// 	}
// 	// 发现还能这么用for...range(不获取value)：for k := range m{}
// 	for k, v := range m {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return -1
// }

/*
	不借助额外空间，空间O(1)
	利用异或操作，异或所有的元素，相同元素异或后都是0，最后剩下的就是落单的数
	16/16 cases passed (8 ms)
	Your runtime beats 98.43 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.7 MB)
	时间O(n)，空间O(1)
*/
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// @lc code=end
