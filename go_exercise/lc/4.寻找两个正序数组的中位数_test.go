package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (38.32%)
 * Likes:    2858
 * Dislikes: 0
 * Total Accepted:    218.8K
 * Total Submissions: 570.9K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
 *
 * 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 *
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 *
 *
 * 示例 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * 则中位数是 2.0
 *
 *
 * 示例 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * 则中位数是 (2 + 3)/2 = 2.5
 *
 *
 */
func TestFindMedianSortedArrays(t *testing.T) {
	log.Printf("res:%f", findMedianSortedArrays([]int{1, 3}, []int{2}))
}

// @lc code=start
/*
	找两个有序数组的中间节点，两个则平均，
	奇数则(m+n)/2的下标；偶数则(m+n)/2 - 1 和 (m+n)/2
	由于要求O(log(m+n))，一般会用到二分思想
	长度m和n，k 和 (m+n)/2 比较，若小于则可过滤一半数据，根据情况判断是两个数组各自哪一半
	2085/2085 cases passed (20 ms)
	Your runtime beats 49.47 % of golang submissions
	Your memory usage beats 90.91 % of golang submissions (5.6 MB)
	时间O(log(m+n))，空间O(1)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	mid := total / 2
	if total%2 == 0 {
		return float64(findKth(nums1, nums2, mid-1)+findKth(nums1, nums2, mid)) / 2.0
	}
	return float64(findKth(nums1, nums2, mid))
}

func findKth(nums1 []int, nums2 []int, k int) int {
	for {
		m1 := len(nums1) / 2
		m2 := len(nums2) / 2

		// 终止条件
		if len(nums1) == 0 {
			// 则剩余要找的数在nums2中
			return nums2[k]
		} else if len(nums2) == 0 {
			return nums1[k]
		} else if k == 0 {
			// 只剩下一个数要找，则取小的那个头结点
			if nums1[0] <= nums2[0] {
				return nums1[0]
			}
			return nums2[0]
		}

		// k在总数的左半部分则过滤大于的数组的右半边
		// 注意此处mid 和 m1+m2还是有差别的
		if k <= m1+m2 {
			if nums1[m1] <= nums2[m2] {
				nums2 = nums2[:m2]
			} else {
				nums1 = nums1[:m1]
			}
		} else {
			// 过滤左边
			if nums1[m1] <= nums2[m2] {
				nums1 = nums1[m1+1:]
				k -= m1 + 1
			} else {
				nums2 = nums2[m2+1:]
				k -= m2 + 1
			}
		}
	}
}

// @lc code=end
