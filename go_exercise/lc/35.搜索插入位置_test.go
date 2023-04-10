package lc

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.77%)
 * Likes:    552
 * Dislikes: 0
 * Total Accepted:    180.8K
 * Total Submissions: 395K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 你可以假设数组中无重复元素。
 *
 * 示例 1:
 *
 * 输入: [1,3,5,6], 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [1,3,5,6], 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 * 输入: [1,3,5,6], 7
 * 输出: 4
 *
 *
 * 示例 4:
 *
 * 输入: [1,3,5,6], 0
 * 输出: 0
 *
 *
 */

// @lc code=start
/*
	62/62 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3.1 MB)
	二分查找后进行判断，时间O(logn)，空间主要是二分查找的栈调用，O(logn)，可以改成循环，则为O(1)
*/
// func searchInsert(nums []int, target int) int {
// 	i, last := halfSearch(nums, 0, len(nums)-1, target)
// 	if i != -1 {
// 		return i
// 	}
// 	if last >= len(nums) || target < nums[last] {
// 		return last
// 	}

// 	return last + 1
// }

// // 先二分查找是否有元素，有则返回下标，没有则下标-1，并返回最后一个判断数据的索引
// func halfSearch(nums []int, begin, end, target int) (int, int) {
// 	if begin > end {
// 		return -1, begin
// 	}
// 	mid := begin + (end-begin)/2
// 	if target == nums[mid] {
// 		return mid, 0
// 	} else if target < nums[mid] {
// 		return halfSearch(nums, begin, mid-1, target)
// 	} else {
// 		return halfSearch(nums, mid+1, end, target)
// 	}
// }

/*
	循环的话，没必要定一个递归函数，时间O(logn)，空间O(1)
	62/62 cases passed (4 ms)
	Your runtime beats 90.22 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3.1 MB)
*/
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = l + (r-l)/2
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}

// @lc code=end
