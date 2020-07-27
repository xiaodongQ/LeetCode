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
	n数之和的思路
	递归，从数组里找满足相加等于target的数
*/
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	n := len(candidates)
	// 从最右开始，至少有一个该位置值
	right := n - 1
	for right >= 0 && candidates[right] > target {
		right--
	}
	log.Printf("right:%d", right)
	if right < 0 {
		return nil
	}

	for right >= 0 {
		var r1 []int
		combHelper(candidates, right, target, r1)
		if len(r1) != 0 {
			res = append(res, r1)
		}

		right--
	}

	return res
}

// 题目中已知成员>=1
func combHelper(candidates []int, right, target int, res []int) {
	if target <= 0 || right < 0 {
		return
	}
	if candidates[right] == target {
		res = append(res, target)
		return
	}
	remain := target - candidates[right]
	combHelper(candidates, right, remain, res)

	return
}

// @lc code=end
