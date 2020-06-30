package array

/*
* 给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
	请返回 nums 的动态和。
* [1480. 一维数组的动态和](https://leetcode-cn.com/problems/running-sum-of-1d-array/)
e.g. 输入：nums = [1,2,3,4]
	输出：[1,3,6,10]
	解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
* 说明：
	1 <= nums.length <= 1000
	-10^6 <= nums[i] <= 10^6
*
*/
func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	sum := 0
	for i, v := range nums {
		sum += v
		res[i] = sum
	}
	return res
}

/*
* 给你两个整数，n 和 start 。
	数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。
	请返回 nums 中所有元素按位异或（XOR）后得到的结果
* [1486. 数组异或操作](https://leetcode-cn.com/problems/xor-operation-in-an-array/)
e.g.
	输入：n = 5, start = 0
	输出：8
	解释：数组 nums 为 [0, 2, 4, 6, 8]，其中 (0 ^ 2 ^ 4 ^ 6 ^ 8) = 8 。
		"^" 为按位异或 XOR 运算符
* 说明：
	1 <= n <= 1000
	0 <= start <= 1000
	n == nums.length
*/
// 暴力求解(当前数据量不大) 时间O(n)，空间O(1)
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}

// 利用异或特性(数据量大时适于利用该方法)
// 还没大理解，参考：[O(1) 位运算](https://leetcode-cn.com/problems/xor-operation-in-an-array/solution/o1-wei-yun-suan-by-bruceyuj/)
func xorOperation2(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}

/*
* [1431. 拥有最多糖果的孩子](https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/)
 */
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	res := make([]bool, len(candies))
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	for i, v := range candies {
		if v+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}

/*
* [1470. 重新排列数组](https://leetcode-cn.com/problems/shuffle-the-array/)
 */
// 时间O(n)，空间O(n)
func shuffle(nums []int, n int) []int {
	// 时间换空间
	res := make([]int, len(nums))
	j := 0
	for i := 0; i < n; i++ {
		res[j] = nums[i]
		j++
		res[j] = nums[n+i]
		j++
	}
	return res
}
