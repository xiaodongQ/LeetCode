package lc

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 *
 * https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (54.71%)
 * Likes:    329
 * Dislikes: 0
 * Total Accepted:    107.3K
 * Total Submissions: 196.1K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 *
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *
 * 说明:
 *
 *
 * 返回的下标值（index1 和 index2）不是从零开始的。
 * 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 *
 *
 * 示例:
 *
 * 输入: numbers = [2, 7, 11, 15], target = 9
 * 输出: [1,2]
 * 解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 *
 */

// @lc code=start
/*
	借助map，时间O(n)，空间O(n)
	17/17 cases passed (4 ms)
	Your runtime beats 96.33 % of golang submissions
	Your memory usage beats 33.33 % of golang submissions (3.2 MB)
*/
// func twoSum(numbers []int, target int) []int {
// 	// 存需要配对的值，并记录自己的索引
// 	m := make(map[int]int, len(numbers))
// 	for i, v := range numbers {
// 		if t, ok := m[v]; ok {
// 			return []int{t + 1, i + 1}
// 		}
// 		m[target-v] = i
// 	}
// 	return []int{}
// }

/*
	双指针，由于是排序好的，左右缩小范围，时间O(n)，空间O(1)
	17/17 cases passed (4 ms)
	Your runtime beats 96.33 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3 MB)
*/
// func twoSum(numbers []int, target int) []int {
// 	low := 0
// 	high := len(numbers) - 1
// 	for low < high {
// 		if numbers[low]+numbers[high] > target {
// 			high--
// 		} else if numbers[low]+numbers[high] < target {
// 			low++
// 		} else {
// 			return []int{low + 1, high + 1}
// 		}
// 	}
// 	return []int{}
// }

/*
	上面的双指针法，可以对low和high提前缩小范围，
		二分法找第一个右边值满足：最小(low)+右边值 > target，右边一些值不需要一步步缩小范围
	17/17 cases passed (4 ms)
	Your runtime beats 96.33 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3 MB)
*/
func twoSum(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1

	mid := high / 2
	for mid < high-1 {
		if numbers[low]+numbers[mid] > target {
			high = mid
		} else {
			mid = (mid + high) / 2
		}
	}

	for low < high {
		if numbers[low]+numbers[high] > target {
			high--
		} else if numbers[low]+numbers[high] < target {
			low++
		} else {
			return []int{low + 1, high + 1}
		}
	}
	return []int{}
}

// @lc code=end
