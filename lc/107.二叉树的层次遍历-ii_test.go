package lc

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (66.07%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    67.9K
 * Total Submissions: 102.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其自底向上的层次遍历为：
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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
	34/34 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.8 MB)
	时间O(n)，空间O(n)
*/
// func levelOrderBottom(root *TreeNode) [][]int {
// 	if root == nil {
// 		return [][]int{}
// 	}
// 	// 先获取自顶向下的每层节点
// 	leveln := 1
// 	s := []*TreeNode{root}
// 	i := 0
// 	res := [][]int{}
// 	lres := []int{}

// 	for len(s) != 0 {
// 		node := s[i]
// 		i++
// 		lres = append(lres, node.Val)
// 		if node.Left != nil {
// 			s = append(s, node.Left)
// 		}
// 		if node.Right != nil {
// 			s = append(s, node.Right)
// 		}

// 		if i == leveln {
// 			tmp := []int{}
// 			tmp = append(tmp, lres...)
// 			res = append(res, tmp)
// 			lres = lres[0:0]

// 			s = s[leveln:]
// 			leveln = len(s)
// 			i = 0
// 		}
// 	}
// 	// 调换res的位置
// 	n := len(res)
// 	for i := 0; i < n/2; i++ {
// 		res[i], res[n-1-i] = res[n-1-i], res[i]
// 	}

// 	return res
// }

// 上面的层序遍历有点繁琐
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 先获取自顶向下的每层节点
	leveln := 1
	s := []*TreeNode{root}
	res := [][]int{}

	for len(s) != 0 {
		lres := []int{}
		for i := 0; i < leveln; i++ {
			node := s[i]
			lres = append(lres, node.Val)
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
		}
		res = append(res, lres)
		s = s[leveln:]
		leveln = len(s)
	}
	// 调换res的位置
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}

// @lc code=end
