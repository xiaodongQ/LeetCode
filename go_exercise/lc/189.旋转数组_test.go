package lc

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (42.13%)
 * Likes:    622
 * Dislikes: 0
 * Total Accepted:    141.2K
 * Total Submissions: 334.9K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */

// @lc code=start
/*
	利用内建函数拷贝两段slice进行拼接，(总len - k%总len)作为开头
	35/35 cases passed (4 ms)
	Your runtime beats 94.9 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.4 MB)
*/
// func rotate(nums []int, k int) {
// 	if len(nums) != 0 {
// 		copy(nums, append(nums[len(nums)-k%len(nums):], nums[:len(nums)-k%len(nums)]...))
// 	}
// }

/*
	移动k次，每次1步
	时间O(k*n)，空间O(1)
	35/35 cases passed (108 ms)
	Your runtime beats 12.38 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.2 MB)
*/
// func rotate(nums []int, k int) {
// 	for i := 0; i < k; i++ {
// 		// 第一个位置处，前一个值为最后一个元素
// 		pre := nums[len(nums)-1]
// 		for j := 0; j < len(nums); j++ {
// 			// 本次位置赋值为前一个位置，并更新前一个位置为本位置原来的值
// 			nums[j], pre = pre, nums[j]
// 		}
// 	}
// }

/*
	使用环状替换，替换时 每次+k位置 进行替换，起始和结束位置相同则结束一轮
		记录移动的元素个数，个数和数组大小相同则说明都移动完了
	(参考官方题解[旋转数组](https://leetcode-cn.com/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode/))
	35/35 cases passed (4 ms)
	Your runtime beats 94.9 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.2 MB)
	时间O(n) 每个元素只移动一次，空间O(1)

	自己最开始的方式是 i移动到 (i+k)/n位置，移动n次，
	但是出现k个之后的元素其实是原来位置移动过来的，这些位置原有数据丢失，解法错误
*/
// func rotate(nums []int, k int) {
// 	count := 0
// 	for i := 0; count < len(nums); i++ {
// 		start := i
// 		pre := nums[start]
// 		next := (start + k) % len(nums)
// 		for start != next {
// 			pre, nums[next] = nums[next], pre
// 			count++
// 			next = (next + k) % len(nums)
// 		}
// 		// 还差一次第一个
// 		pre, nums[next] = nums[next], pre
// 		count++
// 	}
// }

/*
	使用反转，最终尾部的 k%n 个元素移动到数组头部，其他元素都后移
	所以可以总数组先反转，再反转前k%n个，再反转剩下的元素

	35/35 cases passed (4 ms)
	Your runtime beats 94.9 % of golang submissions
	Your memory usage beats 25 % of golang submissions (3.2 MB)
	时间O(n) 每个元素访问两次，2n，空间O(1)
*/
func rotate(nums []int, k int) {
	roReverse(nums)
	roReverse(nums[:k%len(nums)])
	roReverse(nums[k%len(nums):])
}

func roReverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

// @lc code=end
