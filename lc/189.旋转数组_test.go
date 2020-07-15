package lc

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (42.13%)
 * Likes:    622
 * Dislikes: 0
 * Total Accepted:    141.2K
 * Total Submissions: 334.9K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */

// @lc code=start
/*
	利用内建函数拷贝两段slice进行拼接，(总len - k%总len)作为开头
	35/35 cases passed (4 ms)
	Your runtime beats 94.9 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.4 MB)
*/
// func rotate(nums []int, k int) {
// 	if len(nums) != 0 {
// 		copy(nums, append(nums[len(nums)-k%len(nums):], nums[:len(nums)-k%len(nums)]...))
// 	}
// }

/*

 */
func rotate(nums []int, k int) {

}

// @lc code=end
