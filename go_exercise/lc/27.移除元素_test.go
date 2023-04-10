package lc

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 *
 * https://leetcode-cn.com/problems/remove-element/description/
 *
 * algorithms
 * Easy (58.39%)
 * Likes:    589
 * Dislikes: 0
 * Total Accepted:    199.1K
 * Total Submissions: 340.9K
 * Testcase Example:  '[3,2,2,3]\n3'
 *
 * 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
 *
 * 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
 *
 * 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 * 示例 1:
 *
 * 给定 nums = [3,2,2,3], val = 3,
 *
 * 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
 *
 * 你不需要考虑数组中超出新长度后面的元素。
 *
 *
 * 示例 2:
 *
 * 给定 nums = [0,1,2,2,3,0,4,2], val = 2,
 *
 * 函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
 *
 * 注意这五个元素可为任意顺序。
 *
 * 你不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 *
 * 说明:
 *
 * 为什么返回数值是整数，但输出的答案是数组呢?
 *
 * 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
 *
 * 你可以想象内部操作如下:
 *
 * // nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
 * int len = removeElement(nums, val);
 *
 * // 在函数里修改输入数组对于调用者是可见的。
 * // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 *
 *
 */

// @lc code=start
/*
	和删除有序数组重复项类似，双指针思想，从第0个位置开始填数字

	113/113 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 9.09 % of golang submissions (2.1 MB)
	时间O(n)，空间O(1)
	对于时间复杂度，移动的数据做了两次判断
	各个++，+1的调整，出错了好多次。。。这样写临界条件很容易弄错
	作为反例。
*/
// func removeElement(nums []int, val int) int {
// 	if len(nums) == 0 {
// 		// 返回的是数组长度
// 		return 0
// 	}

// 	p := 0
// 	n := len(nums) - 1
// 	for p < n {
// 		if nums[p] != val {
// 			p++
// 		} else {
// 			if nums[n] != val {
// 				// p位置要删除，n位置保留，直接移动即可
// 				nums[p] = nums[n]
// 			}
// 			n--
// 		}
// 	}
// 	if nums[p] != val {
// 		return p + 1
// 	}

// 	return p
// }

/*
	还是遍历数组的方式更直观，
	之前想岔了，想着移动赋值后会覆盖原来位置，其实只有要删除的位置才会被赋值，并不影响

	113/113 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 9.09 % of golang submissions (2.1 MB)
	时间O(n)，空间O(1)
	每个元素只会判断一次，不过不等于val的元素都会被赋值一次
*/
// func removeElement(nums []int, val int) int {
// 	// 记录长度
// 	reslen := len(nums)
// 	cur := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == val {
// 			reslen--
// 			continue
// 		}
// 		nums[cur] = nums[i]
// 		cur++
// 	}
// 	return reslen
// }

/*
	113/113 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 9.09 % of golang submissions (2.1 MB)
	跟第一种方式的思路其实一样，只是自己实现得很绕
	双指针，后指针从最后开始
*/
func removeElement(nums []int, val int) int {
	p := 0
	n := len(nums)
	for p < n {
		if nums[p] != val {
			p++
		} else {
			// 不用管最后一个元素即nums[n-1]是否为val，留给下一次判断，简化逻辑
			// 直接移动即可
			nums[p] = nums[n-1]
			n--
		}
	}
	return n
}

// @lc code=end
