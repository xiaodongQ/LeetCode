package lc

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (34.20%)
 * Likes:    580
 * Dislikes: 0
 * Total Accepted:    75.8K
 * Total Submissions: 221.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须原地修改，只允许使用额外常数空间。
 *
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */

// @lc code=start
/*
错误：
	思路：从最后往前比较，碰到后面比前面大，则交换后返回。若最后无交换，说明最大，则倒置

	交换相邻的两位，可能会导致高位上放的并不是刚好比现在大的数，而是选择了更大的高位
		如：1 3 2，交换后得到3 1 2，而下一个更大的排列应该是2 1 3
*/
// func nextPermutation(nums []int) {
// 	i := len(nums) - 1
// 	for ; i > 0; i-- {
// 		if nums[i] > nums[i-1] {
// 			nums[i], nums[i-1] = nums[i-1], nums[i]
// 			return
// 		}
// 	}
// 	if i == 0 {
// 		// 倒排
// 		for j := 0; j < len(nums)/2; j++ {
// 			nums[j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[j]
// 		}
// 	}
// }

/*
	上面演示了错误的处理
	应该从后面开始，碰到后面比前面大，则从最后开始找到最近一个比前面大的数，并交换位置
	1 3 2， 2 3 1也不行，3也要交换位置
	针对不同的用例又踩不同的坑了。。。按用例修改问题并不简单

	找第一个后面(i)比前面(i-1)大的数，互换后，将原来i-1的数放到合适位置，再将i位置后面的数据倒排
		前面不符合用例，主要是没考虑到后面是从大到小排列的
*/
func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {

			// 从最后开始找一个比 nums[i-1]大的数
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] <= nums[i-1] {
					continue
				}
				// 交换nums[i-1]和找到的nums[j]
				nums[j], nums[i-1] = nums[i-1], nums[j]

				// 把交换的数放到从右边起第一个比它小的位置
				k := len(nums) - 1
				for k > j && nums[j] < nums[k] {
					k--
				}
				nums[k], nums[j] = nums[j], nums[k]

				return
			}
		}
	}
	if i == 0 {
		// 倒排
		for j := 0; j < len(nums)/2; j++ {
			nums[j], nums[len(nums)-1-j] = nums[len(nums)-1-j], nums[j]
		}
	}
}

// @lc code=end
