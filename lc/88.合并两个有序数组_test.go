package lc

import (
	"log"
	"testing"
)

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (48.06%)
 * Likes:    558
 * Dislikes: 0
 * Total Accepted:    173.1K
 * Total Submissions: 360.2K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 *
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */

func TestMerge(t *testing.T) {
	n1 := make([]int, 8)
	n1[0] = 1
	n1[1] = 2
	n1[2] = 3
	log.Printf("len:%d", len(n1))
	merge(n1, 3, []int{2, 5, 6}, 3)
	log.Printf("%v", n1)
}

// @lc code=start
/*
借助一个临时空间，空间O(m)：
	59/59 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 16.67 % of golang submissions (2.3 MB)

若不借助额外空间，插入数组涉及到数组搬移，在nums1中后面空出nums2长度的空间
	归并排序 借助额外的空间做合并，空间O(n)
	快速排序 根据分区点判断来交换位置，空间O(1)

*/
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	// 注意此处 tempn1:=nums1[:m] 这样定义的话，共用底层结构，nums1修改会影响tempn1的值
// 	var tempn1 []int
// 	tempn1 = append(tempn1, nums1[:m]...)
// 	pos := 0
// 	for i, j := 0, 0; i < m || j < n; {
// 		if i < m && j < n {
// 			if tempn1[i] <= nums2[j] {
// 				nums1[pos] = tempn1[i]
// 				i++
// 			} else {
// 				nums1[pos] = nums2[j]
// 				j++
// 			}
// 		} else if i < m {
// 			nums1[pos] = tempn1[i]
// 			i++
// 		} else if j < n {
// 			nums1[pos] = nums2[j]
// 			j++
// 		}
// 		pos++
// 	}
// }

/*
	从后到前合并就不需要另外的空间，时间O(m+n)，空间O(1)
	59/59 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.3 MB)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
		59/59 cases passed (0 ms)
		Your runtime beats 100 % of golang submissions
		Your memory usage beats 25 % of golang submissions (2.3 MB)
		直接用go的sort包最简便
	*/
	// nums1 = append(nums1[:m], nums2[:n]...)
	// sort.Ints(nums1)

	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1
	for p1 >= 0 || p2 >= 0 {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] >= nums2[p2] {
				nums1[p3] = nums1[p1]
				p1--
			} else {
				nums1[p3] = nums2[p2]
				p2--
			}
		} else if p1 >= 0 {
			nums1[p3] = nums1[p1]
			p1--
		} else if p2 >= 0 {
			nums1[p3] = nums2[p2]
			p2--
		}
		p3--
	}
}

// @lc code=end
