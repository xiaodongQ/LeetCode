package lc

import "math"

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (52.42%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    89K
 * Total Submissions: 169.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
 *
 *
 * 示例 1:
 *
 * 给定二叉树 [3,9,20,null,null,15,7]
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回 true 。
 *
 * 示例 2:
 *
 * 给定二叉树 [1,2,2,3,3,null,null,4,4]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   2
 * ⁠   / \
 * ⁠  3   3
 * ⁠ / \
 * ⁠4   4
 *
 *
 * 返回 false 。
 *
 *
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
	依次判断每个节点的左右子树，高度差<=1
	自顶向下 和 自底向上 两种方式

自顶向下，本节点左右子树高度差<=1，然后依次递归左右子树判断是否也满足：
	227/227 cases passed (12 ms)
	Your runtime beats 57.09 % of golang submissions
	Your memory usage beats 100 % of golang submissions (5.7 MB)
	时间：平衡二叉树高度不超过logn，每个节点都要判断，所以时间O(nlogn)
	空间：最坏O(n) 最好O(logn)?
*/
// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 &&
// 		isBalanced(root.Left) &&
// 		isBalanced(root.Right)
// }

// // 时间O(m)，空间O(logm)，m为传入子树的所有子树数量
// func height(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	lh := height(root.Left)
// 	rh := height(root.Right)
// 	if lh > rh {
// 		return lh + 1
// 	}
// 	return rh + 1
// }

/*
	上面自顶向下的方式算节点高度，存在大量冗余，高度更高的节点算高度时会遍历到高度更低的节点
自底向上：
	先判断子树是否高度平衡，再判断父节点
	227/227 cases passed (8 ms)
	Your runtime beats 90.93 % of golang submissions
	Your memory usage beats 60 % of golang submissions (5.7 MB)
	时间O(n)？
		我理解的n个节点，每个节点递归logn次，应该是O(nlogn)
	空间最好O(logn)，最坏O(n)
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	h := 0
	return helper(root, &h)
}

// 返回是否平衡，并返回高度
func helper(root *TreeNode, pheight *int) bool {
	if root == nil {
		// 一颗空的树为平衡，且高度-1
		*pheight = -1
		return true
	}
	// 左右子树的高度
	lh, rh := 0, 0
	// 左子树平衡 且 右子树也平衡，且左右子树的高度差<=1，
	// 则传入节点之下的树是高度平衡的
	if helper(root.Left, &lh) &&
		helper(root.Right, &rh) &&
		math.Abs(float64(rh-lh)) <= 1 {
		// 本传入节点的高度max(左,右子树高度)+1
		if lh > rh {
			*pheight = lh + 1
		} else {
			*pheight = rh + 1
		}
		return true
	}

	return false
}

// @lc code=end
