package lc

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (73.42%)
 * Likes:    513
 * Dislikes: 0
 * Total Accepted:    97.1K
 * Total Submissions: 132.3K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定有序数组: [-10,-3,0,5,9],
 *
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
	每次获取中间节点作为根节点构造子树，自顶向下，前序遍历
	32/32 cases passed (4 ms)
	Your runtime beats 85.91 % of golang submissions
	Your memory usage beats 28.57 % of golang submissions (4.4 MB)
	时间O(n)，空间O(logn)，此处为高度平衡的二叉树，不存在退化问题
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return sortHelper(nums, 0, len(nums)-1)
}

func sortHelper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	root := &TreeNode{}
	root.Val = nums[m]

	root.Left = sortHelper(nums, l, m-1)
	root.Right = sortHelper(nums, m+1, r)
	return root
}

// @lc code=end
