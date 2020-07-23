package lc

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (28.71%)
 * Likes:    2412
 * Dislikes: 0
 * Total Accepted:    283.7K
 * Total Submissions: 987.9K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
/*
失败：
	暴力方法，取一个数a，在剩下的数里找两个满足 0-a，取b在剩下的数找满足：0-a-b
		为了区分某个值已经用了，还要一个n大小的数组做标志
		时间O(n^3)，空间O(n)

	暴力法超时了
	Time Limit Exceeded
	311/313 cases passed (N/A)
*/
// func threeSum(nums []int) [][]int {
// 	res := [][]int{}
// 	for i := 0; i < len(nums); i++ {
// 		a := nums[i]
// 		for j := 0; j != i && j < len(nums); j++ {
// 			b := nums[j]
// 		lp:
// 			for k := 0; k != i && k != j && k < len(nums); k++ {
// 				c := nums[k]
// 				if a+b+c == 0 {
// 					// 去重
// 					pair := []int{a, b, c}
// 					sort.Ints(pair)
// 					for r := 0; r < len(res); r++ {
// 						if pair[0] == res[r][0] &&
// 							pair[1] == res[r][1] &&
// 							pair[2] == res[r][2] {
// 							continue lp
// 						}
// 					}
// 					res = append(res, pair)
// 				}
// 			}
// 		}
// 	}

// 	return res
// }

/*
	上面的暴力法超出时间限制
	先对数据排序，再进行遍历：
		第一层循环(a)从小到大，并排除和上一次遍历值相同的元素
		第二层循环(b)从a位置后面开始，且从最后一个元素开始取c，
			有b<=c
			如果b+c > target-a，则c往前移，直到c移到当前b的位置
				(只是到值和b相同并不对0 0)，还不满足条件则再移动也不会有满足的数据
		~~第三层循环即上面c的移动~~ 此处需要优化

		去掉第三层循环，通过c的取值和移动(内层相当于双指针)

	313/313 cases passed (28 ms)
	Your runtime beats 99.18 % of golang submissions
	Your memory usage beats 66.67 % of golang submissions (6.9 MB)

*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		// 去重，同a值并不需要再遍历
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		// c的位置
		third := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			b := nums[j]

			// 这里是精髓，双指针，移动右边指针
			for b < nums[third] && b+nums[third] > -a {
				third--
			}
			if j == third {
				// c左移到了当前取b的位置(并不是数组值为b)
				break
			}

			if a+b+nums[third] == 0 {
				res = append(res, []int{a, b, nums[third]})
			}
		}
	}

	return res
}

// @lc code=end
