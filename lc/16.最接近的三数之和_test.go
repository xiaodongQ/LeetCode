package lc

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.75%)
 * Likes:    516
 * Dislikes: 0
 * Total Accepted:    137.2K
 * Total Submissions: 300K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 *
 *
 * 示例：
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 10^3
 * -10^3 <= nums[i] <= 10^3
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
/*
	和三数之和一样的思路，先定义一个接近值
	131/131 cases passed (4 ms)
	Your runtime beats 95.1 % of golang submissions
	Your memory usage beats 25 % of golang submissions (2.7 MB)
	时间O(n^2) (排序用的O(nlogn)数量级被包含了)，空间O(1) (改变了原切片，若要求不能改变，则临时空间O(n))
*/
func threeSumClosest(nums []int, target int) int {
	res := 0
	diff := math.MaxInt64
	// 先排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSum := target - nums[i]
		// 双指针
		low := i + 1
		high := len(nums) - 1
		for low < high {
			// 为了避免下面的重复计算
			tmpDiff := myAbs(nums[i] + nums[low] + nums[high] - target)
			tmpRes := nums[i] + nums[low] + nums[high]

			if nums[low]+nums[high] == twoSum {
				// 如果有三数之和正好为指定值的，则直接返回即可
				return target
			} else if nums[low]+nums[high] > twoSum {
				if tmpDiff < diff {
					res = tmpRes
					diff = tmpDiff
				}
				// 此处优化点，对于相同的元素不需要一步步移动(high--)，
				// 此处high不应该放for里面，有相等元素才进入，没有则high不变，就死循环了
				// 注意h0赋值时还要先-1
				h0 := high - 1
				// 循环判断h0，而不用h0-1，上面已经high-1了
				for h0 > low && nums[h0] == nums[h0+1] {
					h0--
				}
				high = h0
			} else {
				if tmpDiff < diff {
					res = tmpRes
					diff = tmpDiff
				}
				// 优化 l0显示起来像10，还是low0可读性好些
				low0 := low + 1
				for low0 < high && nums[low0] == nums[low0-1] {
					low0++
				}
				low = low0
			}
		}
	}

	return res
}

func myAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
