package lc

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (52.39%)
 * Likes:    899
 * Dislikes: 0
 * Total Accepted:    173.6K
 * Total Submissions: 331.3K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 *
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
	只操作一个root节点陷入僵局了，两个树同步遍历清晰多了

	195/195 cases passed (4 ms)
	Your runtime beats 75.31 % of golang submissions
	Your memory usage beats 57.14 % of golang submissions (2.9 MB)
	时间O(n)，空间：完全平衡二叉树时最优，为O(logn)；退化为链表时最坏，为O(n)
*/
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}
func check(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	// 到此处只剩t1和t2都不为nil。
	// 取树1的左子树 和 树2的右子树进行递归
	// 并且还要取树1的右子树 和 树2的左子树进行递归
	return t1.Val == t2.Val && check(t1.Left, t2.Right) && check(t1.Right, t2.Left)
}

// @lc code=end
