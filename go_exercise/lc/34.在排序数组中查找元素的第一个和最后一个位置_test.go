package lc

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (39.99%)
 * Likes:    522
 * Dislikes: 0
 * Total Accepted:    113.5K
 * Total Submissions: 283.9K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

// @lc code=start
/*
	二分法找到目标值后，依次查找其左右
	88/88 cases passed (8 ms)
	Your runtime beats 92.3 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.1 MB)
*/
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			res := []int{mid, mid}
			for res[0] > 0 && nums[res[0]-1] == target {
				res[0]--
			}
			for res[1] < len(nums)-1 && nums[res[1]+1] == target {
				res[1]++
			}
			return res
		}
	}
	return []int{-1, -1}
}

// @lc code=end
