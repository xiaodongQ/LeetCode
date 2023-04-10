package lc

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
 *
 * https://leetcode-cn.com/problems/path-sum/description/
 *
 * algorithms
 * Easy (50.84%)
 * Likes:    378
 * Dislikes: 0
 * Total Accepted:    114.5K
 * Total Submissions: 225.2K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,null,1]\n22'
 *
 * 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 *
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \      \
 * ⁠       7    2      1
 *
 *
 * 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
	子节点的路径和 = 父节点路径和-父节点值，转换为左右节点是否满足，直到叶子节点和剩余的数相等
深度优先：
	114/114 cases passed (4 ms)
	Your runtime beats 98.34 % of golang submissions
	Your memory usage beats 66.67 % of golang submissions (4.8 MB)
	时间O(n)，空间最好O(logn)，最坏O(n)
广度优先：
	从根节点开始层序遍历节点，借助两个队列，一个存节点，一个存节点到根节点的路径和(注意不是到子节点)
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点，此处细节：先判断是否满足节点值是否和剩余的sum相等，短路特性减少判断
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	// 此处细节：先减去父节点，避免下面递归时计算两次
	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// @lc code=end
