package lc

import (
	"log"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (69.37%)
 * Likes:    774
 * Dislikes: 0
 * Total Accepted:    114.5K
 * Total Submissions: 165K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1：
 *
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2：
 *
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 *
 *
 */

func TestCombinationSum(t *testing.T) {
	log.Printf("res:%v", combinationSum([]int{2, 3, 6, 7}, 7))
}

// @lc code=start
/*
	回溯
	候选数组先排序，依次从开始去找能组成和为 target-candidates[pos] 的数，
	直到剩余的和==0则说明符合条件，放到结果里(DFS深度遍历)
	剩余的和<0则说明不符合条件，终止本次递归

	168/168 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 77.5 % of golang submissions (2.7 MB)
	时间O(2^n)，空间O(2^n)
		可以理解为每个数据平均判断两次，是否加到path(DFS)，所以时间为O(2^n)
		空间也是O(2^n)，没大理解；由于可以拿重复元素，每轮回溯不一定只有n轮调用栈，所以不是O(n)

		let s = target / min(nums[i]) T = C(s,1) + C(s, 2) + ... + C(s, s) = 2^s
		分析参考：[Combination - 1 Method for 6 Problems (39, 40, 77, 78, 90, 216)](https://leetcode.com/problems/combination-sum/discuss/389405/Combination-1-Method-for-6-Problems-(39-40-77-78-90-216))
*/
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)

	path := []int{}
	backtrack(candidates, target, 0, &res, path)

	return res
}

// pos用于记录从哪个位置开始找，避免重复结果
// path用于记录本次遍历过的数据
func backtrack(candidates []int, target int, pos int, res *[][]int, path []int) {
	if target == 0 {
		// 复制一个新的slice，避免复用底层数组时后面修改会影响本次结果
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	if target < 0 {
		return
	}
	for i := pos; i < len(candidates); i++ {
		// 候选有序数组中第一个数就比和要大，则肯定没有符合记录的数
		if target < candidates[i] {
			return
		}
		// 添加到遍历集中进行下一次回溯
		path = append(path, candidates[i])
		// 回溯，pos位置从本记录(i)开始
		backtrack(candidates, target-candidates[i], i, res, path)
		// 本次位置回溯完后，从path里删除，用于下一次回溯
		path = path[:len(path)-1]
	}
}

// @lc code=end
