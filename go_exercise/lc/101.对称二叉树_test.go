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
	时间O(n)，
	空间复杂度和递归次数有关(深度优先DFS)：
		完全平衡二叉树时最优，层数为O(logn)；空间复杂度O(logn)
		最坏退化为链表，O(n)
		(因为函数压栈退出跟层数有关)
*/
// func isSymmetric(root *TreeNode) bool {
// 	return check(root, root)
// }
// func check(t1 *TreeNode, t2 *TreeNode) bool {
// 	if t1 == nil && t2 == nil {
// 		return true
// 	}
// 	if t1 == nil || t2 == nil {
// 		return false
// 	}
// 	// 到此处只剩t1和t2都不为nil。
// 	// 取树1的左子树 和 树2的右子树进行递归
// 	// 并且还要取树1的右子树 和 树2的左子树进行递归
// 	return t1.Val == t2.Val && check(t1.Left, t2.Right) && check(t1.Right, t2.Left)
// }

/*
	使用迭代，和层序遍历的思想一样(广度优先BFS)，借助队列，
	此处第一次入队时入两次root，后面入队时两个树节点的左、右子树分别入队
		树1的左节点和树2的右节点，然后树1的右节点和树2的左节点
		出队时每次出两个
	195/195 cases passed (4 ms)
	Your runtime beats 75.31 % of golang submissions
	Your memory usage beats 28.57 % of golang submissions (3.1 MB)
	时间O(n)，空间O(n)
*/
func isSymmetric(root *TreeNode) bool {
	l := make([]*TreeNode, 0)
	l = append(l, root)
	l = append(l, root)
	for len(l) != 0 {
		// 出队，FIFO，出两个节点
		n1 := l[0]
		n2 := l[1]
		l = l[2:]

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		l = append(l, n1.Left)
		l = append(l, n2.Right)
		l = append(l, n1.Right)
		l = append(l, n2.Left)
	}
	return true
}

// @lc code=end
