package lc

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (38.03%)
 * Likes:    519
 * Dislikes: 0
 * Total Accepted:    93.3K
 * Total Submissions: 245.2K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

// @lc code=start
/*
	和三数类似，最后两层用双指针
	282/282 cases passed (12 ms)
	Your runtime beats 73.41 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.7 MB)
	时间O(n^3)，空间O(1)
*/
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	// 剩余中取3个数，不用到len(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			need := target - nums[i] - nums[j]
			low, high := j+1, len(nums)-1
			for low < high {
				twoSum := nums[low] + nums[high]
				if twoSum == need {
					res = append(res, []int{nums[i], nums[j], nums[low], nums[high]})
					// 去重
					low++
					high--
					for low < high && nums[low] == nums[low-1] {
						low++
					}
					for low < high && nums[high] == nums[high+1] {
						high--
					}
				} else if twoSum < need {
					low++
				} else {
					high--
				}
			}
		}
	}

	return res
}

// @lc code=end
