package lc

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (61.68%)
 * Likes:    656
 * Dislikes: 0
 * Total Accepted:    176.8K
 * Total Submissions: 286.5K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

// @lc code=start
/*
	双指针，按顺序填值，位置指针指向要填的位置
	21/21 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.8 MB)
	时间O(n)，空间O(1)
*/
// func moveZeroes(nums []int) {
// 	pos := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != 0 {
// 			nums[pos] = nums[i]
// 			if i != pos {
// 				nums[i] = 0
// 			}
// 			// 0则不需要++
// 			pos++
// 		}
// 	}
// 	// 不需要做这个操作了，前面已经置0了
// 	// 非0元素都移到前面了，剩下的位置都置0
// 	// for i := pos; i < len(nums); i++ {
// 	// 	nums[i] = 0
// 	// }
// }

// 和上面一样，理解方式不同，相当于交换操作
func moveZeroes(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 若i和pos一样，则说明原地不同
			nums[pos], nums[i] = nums[i], nums[pos]
			// 0则不需要++
			pos++
		}
	}
}

// @lc code=end
