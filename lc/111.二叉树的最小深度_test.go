package lc

import "math"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (42.85%)
 * Likes:    291
 * Dislikes: 0
 * Total Accepted:    90.2K
 * Total Submissions: 210.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最小深度  2.
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
	根节点的深度为1
	最小深度为 min(左右子树深度)+1，其中nil深度为0
	41/41 cases passed (4 ms)
	Your runtime beats 97.76 % of golang submissions
	Your memory usage beats 66.67 % of golang submissions (5.3 MB)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	mindepth := math.MaxInt32
	if root.Left != nil {
		mindepth = int(math.Min(float64(mindepth), float64(minDepth(root.Left))))
	}
	if root.Right != nil {
		mindepth = int(math.Min(float64(mindepth), float64(minDepth(root.Right))))
	}

	return mindepth + 1
}

// @lc code=end
