package lc

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (72.22%)
 * Likes:    595
 * Dislikes: 0
 * Total Accepted:    201.5K
 * Total Submissions: 278.8K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [1,3,2]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
	68/68 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 7.69 % of golang submissions (2.1 MB)
	时间O(n)，空间O(logn)
*/
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}
// 	res := []int{}
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)

// 	return res
// }

/*
	循环
	68/68 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	// 需要后进先出，和层序遍历不同
	s := []*TreeNode{}
	cur := root

	for cur != nil || len(s) != 0 {
		for cur != nil {
			// 左边节点都入栈
			s = append(s, cur)
			cur = cur.Left
		}
		// 弹出一个节点
		cur = s[len(s)-1]
		s = s[:len(s)-1]
		// 左子树结果
		res = append(res, cur.Val)

		cur = cur.Right
	}

	return res
}

// @lc code=end
