package lc

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (63.79%)
 * Likes:    1658
 * Dislikes: 0
 * Total Accepted:    247.7K
 * Total Submissions: 388.2K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例：
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 *
 */

// @lc code=start
/*
	双指针，起始和结尾，算一个面积，然后移动对应值小的那个指针，比较出最大面积
	证明要参考题解。
	理解：左右两个指针定好后，小的指针不动，任意移动大的那个指针(向内)，组成的面积肯定是<=原位置组成的面积
		下一次循环移动对应值小的指针，无论怎么移动另一个也是<=确定好小指针位置后的面积
		于是就变成了每次移动小指针，求面积(宽度*小值高度)后，再比较面积求最大
	50/50 cases passed (16 ms)
	Your runtime beats 79.35 % of golang submissions
	Your memory usage beats 16.67 % of golang submissions (5.8 MB)
	时间O(n)，空间O(1)
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		area := (j - i) * minHeight(height[i], height[j])
		if area > res {
			res = area
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func minHeight(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
