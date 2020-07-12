package lc

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (73.70%)
 * Likes:    599
 * Dislikes: 0
 * Total Accepted:    205.1K
 * Total Submissions: 278.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 *
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 返回它的最大深度 3 。
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
	最大高度=max(左子树高度, 右子树高度) + 1
	39/39 cases passed (4 ms)
	Your runtime beats 92.43 % of golang submissions
	Your memory usage beats 50 % of golang submissions (4.4 MB)
	时间O(n)，空间最好O(logn)，最坏O(n)
*/
// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	lmax := maxDepth(root.Left)
// 	rmax := maxDepth(root.Right)
// 	if lmax > rmax {
// 		return lmax + 1
// 	}
// 	return rmax + 1
// }

/*
	使用迭代，bfs
	39/39 cases passed (4 ms)
	Your runtime beats 92.43 % of golang submissions
	Your memory usage beats 66.67 % of golang submissions (4.4 MB)
	时间O(n)，空间O(n)
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := []*TreeNode{root}
	// 记录每层的个数
	leveln := 1
	// 获取节点后当前位置
	i := 0
	depth := 0
	for len(s) != 0 {
		// 取出一个节点
		n := s[i]
		i++
		if n.Left != nil {
			s = append(s, n.Left)
		}
		if n.Right != nil {
			s = append(s, n.Right)
		}
		if i == leveln {
			i = 0
			s = s[leveln:]
			leveln = len(s)
			depth++
		}
	}
	return depth
}

// @lc code=end
