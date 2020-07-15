package lc

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (63.82%)
 * Likes:    663
 * Dislikes: 0
 * Total Accepted:    186.5K
 * Total Submissions: 292.2K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */

// @lc code=start
/*
	记录每个数据的次数
	46/46 cases passed (20 ms)
	Your runtime beats 93.49 % of golang submissions
	Your memory usage beats 80 % of golang submissions (5.9 MB)
	时间O(n)，空间O(n)
*/
// func majorityElement(nums []int) int {
// 	limit := len(nums) / 2
// 	m := make(map[int]int, limit)
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]]++
// 		if m[nums[i]] > limit {
// 			return nums[i]
// 		}
// 	}
// 	// 放到上面的循环里判断，已经出现结果则不需要继续
// 	// for k, v := range m {
// 	// 	if v > limit {
// 	// 		return k
// 	// 	}
// 	// }
// 	return -1
// }

/*
	Boyer-Moore 投票算法(摩尔投票算法，多数投票算法)
	如果我们把众数记为 +1，把其他数记为 −1，将它们全部加起来，显然和大于 0，从结果本身我们可以看出众数比其他数多

	维护一个候选众数 candidate 和它出现的次数 count。
		初始时 candidate 可以为任意值，count 为 0；
		遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，
		如果 count 的值为 0，我们先将 x 的值赋予 candidate，随后我们判断 x：相等则计数器+1，否则-1
	在遍历完成后，candidate 即为整个数组的众数

	46/46 cases passed (16 ms)
	Your runtime beats 99.61 % of golang submissions
	Your memory usage beats 50 % of golang submissions (5.9 MB)
	时间O(n)，空间O(1)
*/
func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// @lc code=end
