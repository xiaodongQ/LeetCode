package lc

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode-cn.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (75.85%)
 * Likes:    502
 * Dislikes: 0
 * Total Accepted:    96.8K
 * Total Submissions: 127.6K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 翻转一棵二叉树。
 *
 * 示例：
 *
 * 输入：
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 *
 * 输出：
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 *
 * 备注:
 * 这个问题是受到 Max Howell 的 原问题 启发的 ：
 *
 * 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
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
方法1:递归交换左右子树
	假设左右子树已经交换好了，直接操作交换好的子树即可
	注意先保存要改变值的节点，或直接利用go的交换
	68/68 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 88.89 % of golang submissions (2.1 MB)
	时间O(n)，空间最好O(logn)，最坏O(n)
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// tmp := invertTree(root.Left)
	// root.Left = invertTree(root.Right)
	// root.Right = tmp
	// 上面的交换左右子树，用下面一句即可
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

// 递归处理也可以这样理解，交换本节点，再交换左右
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	root.Left, root.Right = root.Right, root.Left
// 	invertTree(root.Left)
// 	invertTree(root.Right)

// 	return root
// }

/*
方法2:
	广优先BFS，迭代方式，每层节点依次交换子树
		(不用整层都保存下来首尾交换，首尾交换还存在节点为nil时的判断处理)
		不用记录每层有多少节点，迭代执行知道空就行
	68/68 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 11.11 % of golang submissions (2.1 MB)
	时间O(n)，空间O(n)
*/
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	arrNode := []*TreeNode{root}
// 	for len(arrNode) != 0 {
// 		node := arrNode[0]
// 		arrNode = arrNode[1:]
// 		node.Left, node.Right = node.Right, node.Left
// 		if node.Left != nil {
// 			arrNode = append(arrNode, node.Left)
// 		}
// 		if node.Right != nil {
// 			arrNode = append(arrNode, node.Right)
// 		}
// 	}

// 	return root
// }

// @lc code=end
