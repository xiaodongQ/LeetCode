package lc

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (38.38%)
 * Likes:    845
 * Dislikes: 0
 * Total Accepted:    149.8K
 * Total Submissions: 390.3K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
/*
	看到logn差不多就是二分思想
	196/196 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 71.43 % of golang submissions (2.6 MB)
	时间O(logn)，2次O(logn)操作，空间O(1)

	对于只一次O(logn)操作的解法，参考题解
		理解起来需要动些脑筋，要是实践项目里用感觉自己容易搞错边界，倾向简洁清晰的逻辑
*/
func search(nums []int, target int) int {
	// 每次最好先把len临时存储，可以防止操作过程里slice大小改变
	n := len(nums)
	// 后续练习命名都调整成左右指针left和right
	low, high := 0, n-1
	// 二分先找起点
	for low < high {
		mid := low + (high-low)/2
		// 此处重点：应该比较中点和右边的值，比high大才说明起点在右边
		// (比low大并不能说明在右边，正常升序就不满足 0 1 2 3 5)
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	// 然后从起点开始重新用二分
	left, right := low, low+n-1
	for left <= right {
		// mid不先做mod，后续的left和right调整依赖于它
		mid := left + (right-left)/2
		midVal := nums[mid%n]
		// 直接判断退出，省略下面两次判断
		if target == midVal {
			return mid % n
		}
		if target < midVal {
			right = mid - 1
		} else if target > midVal {
			left = mid + 1
		}
	}

	return -1
}

// @lc code=end
