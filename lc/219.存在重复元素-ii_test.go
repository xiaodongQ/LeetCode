package lc

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (39.65%)
 * Likes:    182
 * Dislikes: 0
 * Total Accepted:    51.9K
 * Total Submissions: 130.9K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的
 * 绝对值 至多为 k。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 *
 */

// @lc code=start
/*
	23/23 cases passed (16 ms)
	Your runtime beats 93.5 % of golang submissions
	Your memory usage beats 20 % of golang submissions (7.6 MB)
	时间O(n)，空间O(n)
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	// 记录索引
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			if i-v <= k {
				return true
			}
			// 更新最新的下标
			m[nums[i]] = i
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

// @lc code=end
