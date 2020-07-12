package lc

/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 *
 * https://leetcode-cn.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (58.14%)
 * Likes:    395
 * Dislikes: 0
 * Total Accepted:    100.5K
 * Total Submissions: 172.7K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 *
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * 示例 1:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 *
 * ⁠       [1,2,3],   [1,2,3]
 *
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:      1          1
 * ⁠         /           \
 * ⁠        2             2
 *
 * ⁠       [1,2],     [1,null,2]
 *
 * 输出: false
 *
 *
 * 示例 3:
 *
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 *
 * ⁠       [1,2,1],   [1,1,2]
 *
 * 输出: false
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
	依次比较本节点、左右子树，出现一个不同则不一样，直到都遍历完
	57/57 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 60 % of golang submissions (2.1 MB)
	时间O(n)，
	空间并不是O(1)，需要递归，最优递归判断次数最大log(n)+1，最坏n
	(完全平衡二叉树层数<=log2(n)+1，不过若树退化为链表，则层数为n)(高度为层数-1)
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// 不需要做这么多判断，上面已经过滤了都为nil的情况，所以只要判断其中一个是否为nil，
	// 还有一种都不为nil而值不同的情况，nil的情况上面已经过滤退出了
	// if (p != nil && q == nil) ||
	// 	(p == nil && q != nil) ||
	// 	(p != nil && q != nil && p.Val != q.Val) {
	// 	return false
	// }
	if p == nil ||
		q == nil ||
		p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// @lc code=end
