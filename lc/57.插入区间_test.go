package lc

import "sort"

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (37.42%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    27.3K
 * Total Submissions: 73K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 *
 *
 * 示例 2：
 *
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 *
 *
 *
 * 注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。
 *
 */

// @lc code=start
/*
	可和[56] 合并区间一样，插入后排序，再合并
		时间O(nlogn)
		这样浪费了原来区间已经有序的条件
	154/154 cases passed (32 ms)
	Your runtime beats 5.39 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.6 MB)
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	// 复用56题中定义的slice结构，其实现了sort需要的Interface
	sort.Sort(ByStart2(intervals))
	for i := 0; i < len(intervals)-1; {
		cur, next := intervals[i], intervals[i+1]
		if cur[1] < next[0] {
			i++
		} else {
			// 需要合并
			if cur[1] < next[1] {
				// 后一个结束位置大则作为结束位置，否则以本区间结束为结束
				cur[1] = next[1]
			}
			// 删除下一个区间
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		}
	}

	return intervals
}

// 为了避免和56题里定义冲突
type ByStart2 [][]int

func (bs ByStart2) Len() int {
	return len(bs)
}

func (bs ByStart2) Less(i, j int) bool {
	return bs[i][0] < bs[j][0]
}

func (bs ByStart2) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

/*
	利用已经有序的前提条件，找到插入区间起始start和结束end 对应于有序区间列表 的位置
		二分法查找，第一个结束<=start的区间，和第一个结束<=end的区间
*/
// func insert(intervals [][]int, newInterval []int) [][]int {
// 	var res [][]int

// 	left := 0
// 	right := len(intervals) - 1
// 	// 区间结束位置小于新插入起始的，都添加到结果里
// 	for left < len(intervals) && newInterval[0] > intervals[left][1] {
// 		res = append(res, intervals[left])
// 		left++
// 	}
// 	for right >= 0 && newInterval[1] < intervals[right][0] {
// 		right--
// 	}

// 	return res
// }

// @lc code=end
