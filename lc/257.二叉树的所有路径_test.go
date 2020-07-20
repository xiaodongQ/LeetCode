package lc

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (64.41%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    45K
 * Total Submissions: 69.8K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
	BFS，用单独的一个slice来保存要拼接的信息，和节点slice是一致的
	209/209 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.4 MB)
	时间O(n)，空间O(n)
*/
// func binaryTreePaths(root *TreeNode) []string {
// 	if root == nil {
// 		return []string{}
// 	}
// 	res := make([]string, 0)
// 	nodeSlice := make([]*TreeNode, 0)
// 	strSlice := make([]string, 0)
// 	nodeSlice = append(nodeSlice, root)
// 	strSlice = append(strSlice, strconv.Itoa(root.Val))
// 	for len(nodeSlice) != 0 && len(strSlice) != 0 {
// 		node := nodeSlice[0]
// 		nodeSlice = nodeSlice[1:]

// 		tmpstr := strSlice[0]
// 		strSlice = strSlice[1:]
// 		if node.Left == nil && node.Right == nil {
// 			// 叶子节点，上一步已经->加进去了
// 			res = append(res, tmpstr)
// 		}
// 		if node.Left != nil {
// 			str := tmpstr + "->" + strconv.Itoa(node.Left.Val)
// 			nodeSlice = append(nodeSlice, node.Left)
// 			strSlice = append(strSlice, str)
// 		}

// 		if node.Right != nil {
// 			str := tmpstr + "->" + strconv.Itoa(node.Right.Val)
// 			nodeSlice = append(nodeSlice, node.Right)
// 			strSlice = append(strSlice, str)
// 		}
// 	}

// 	return res
// }

/*
	递归，DFS
	"->" + 当前节点，分别递归左右子树
	209/209 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.3 MB)
	时间O(n)，空间和层数有关，最坏O(n)
*/
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	paths(root, "", &res)

	return res
}

// prefix为该节点之前的遍历路径str， res存结果
func paths(root *TreeNode, prefix string, res *[]string) {
	if root == nil {
		return
	}
	if prefix == "" {
		prefix += strconv.Itoa(root.Val)
	} else {
		prefix += "->" + strconv.Itoa(root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, prefix)
		return
	}
	// 可以不判断nil，入口时有处理
	paths(root.Left, prefix, res)
	paths(root.Right, prefix, res)
}

// @lc code=end
