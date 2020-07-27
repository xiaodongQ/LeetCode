package lc

import (
	"log"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (42.72%)
 * Likes:    520
 * Dislikes: 0
 * Total Accepted:    121.2K
 * Total Submissions: 283.7K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

func TestMerge1(t *testing.T) {
	log.Printf("res:%v", merge1([][]int{[]int{1, 4}, []int{0, 0}}))
}

// @lc code=start
/*
	按区间起始值排序，
		当前区间结束位置小于下一个开始位置，则不需合并
		当前区间结束位置>=下一个开始位置，则需合并这两个
	169/169 cases passed (12 ms)
	Your runtime beats 80.96 % of golang submissions
	Your memory usage beats 11.11 % of golang submissions (6.1 MB)
	时间O(nlogn)，O(nlogn)是排序需要的时间(快排)，O(n)是合并
	空间O(n) 用于保存结果
*/
func merge1(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Sort(ByStart(intervals))
	for i := 0; i < len(intervals); i++ {
		// 和res里面的区间比较
		// cur, next := intervals[i], intervals[i+1]
		cur := intervals[i]
		// 不用合并，直接添加记录
		if len(res) == 0 || res[len(res)-1][1] < cur[0] {
			res = append(res, cur)
			continue
		}
		// 需要合并，更新结束位置
		if res[len(res)-1][1] <= cur[1] {
			res[len(res)-1][1] = cur[1]
		}
	}

	return res
}

/*
	不用额外空间，直接在传入的slice上操作
	169/169 cases passed (32 ms)
	Your runtime beats 12.55 % of golang submissions
	Your memory usage beats 100 % of golang submissions (4.6 MB)
	时间O(nlogn)，比上面时间久，涉及slice的拼接操作，空间~~O(1)~~，空间O(logn)排序辅助空间
*/
// func merge(intervals [][]int) [][]int {
// 	sort.Sort(ByStart(intervals))
// 	for i := 0; i < len(intervals)-1; {
// 		// 和res里面的区间比较
// 		cur, next := intervals[i], intervals[i+1]
// 		if cur[1] < next[0] {
// 			i++
// 			continue
// 		}

// 		// 比下一个的结束小，则当前结束更新
// 		if cur[1] < next[1] {
// 			cur[1] = next[1]
// 		}
// 		// 删除next
// 		intervals = append(intervals[:i+1], intervals[i+2:]...)
// 	}

// 	return intervals
// }

// 将输入的数据按区间起始位置升序排序，实现sort包的接口，然后就可以使用排序接口了
// [[1,3], [5,8]] 形式的slice
type ByStart [][]int

func (bs ByStart) Len() int {
	return len(bs)
}

func (bs ByStart) Less(i, j int) bool {
	return bs[i][0] < bs[j][0]
}

func (bs ByStart) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

// @lc code=end
