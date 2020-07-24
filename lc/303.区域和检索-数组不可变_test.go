package lc

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (62.21%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    46.6K
 * Total Submissions: 74.8K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n' +
  '[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
 *
 * 示例：
 *
 * 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 * 说明:
 *
 *
 * 你可以假设数组不可变。
 * 会多次调用 sumRange 方法。
 *
 *
*/

// @lc code=start
/*
   16/16 cases passed (68 ms)
   Your runtime beats 37.37 % of golang submissions
   Your memory usage beats 100 % of golang submissions (8.2 MB)
   时间O(n)，空间O(1)

   题目中的：会多次调用 sumRange 方法
   给定数组后，使用过程中会送不同范围，
       以上面的方法，每次调用都会访问数组然后累加成员，O(n)
   可以优化处理，对结果进行缓存
*/
// type NumArray struct {
// 	data []int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{data: nums}
// }

// func (this *NumArray) SumRange(i int, j int) int {
// 	res := 0
// 	if j > len(this.data)-1 {
// 		return -1
// 	}

// 	for p := i; p <= j; p++ {
// 		res += this.data[p]
// 	}
// 	return res
// }

/*
   16/16 cases passed (32 ms)
   Your runtime beats 98.71 % of golang submissions
   Your memory usage beats 100 % of golang submissions (8.2 MB)
   时间O(1)，空间O(n)
   可以在make数组的时候，len(nums)+1，用一个没有用的辅助0位置，避免每次的i==0判断
*/
type NumArray struct {
	// 定义从0-原数组每个元素的累加和
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
	}
	return NumArray{sum}
}

// i和j范围内，用的是累加到j和i-1的和，再相减 (因为sum[i]包含了i，不应该减去i)
func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
// @lc code=end
