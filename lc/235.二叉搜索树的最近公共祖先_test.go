package lc

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (64.45%)
 * Likes:    336
 * Dislikes: 0
 * Total Accepted:    65.9K
 * Total Submissions: 102.2K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * 输出: 6
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 *
 *
 * 示例 2:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
/*
	从根节点遍历BST，若给定的节点都在一颗子树上，则以那颗子树为根节点继续遍历
	直到分属于不同的子树，则当时的根节点即为最深公共祖先(LCA)
	27/27 cases passed (24 ms)
	Your runtime beats 81.6 % of golang submissions
	Your memory usage beats 100 % of golang submissions (6.6 MB)
	时间O(n)，空间O(1)
*/
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	lca := root
// 	preRoot := root
// 	// 公共祖先
// 	for lca != nil {
// 		preRoot = lca
// 		// 若给定节点为祖先和子节点关系，则已经不需要继续了
// 		if p.Val == lca.Val || q.Val == lca.Val {
// 			break
// 		}

// 		if p.Val < lca.Val && q.Val < lca.Val {
// 			lca = lca.Left
// 		} else if p.Val > lca.Val && q.Val > lca.Val {
// 			lca = lca.Right
// 		} else {
// 			break
// 		}
// 	}
// 	return preRoot
// }

/*
	跟上面一样，对于==的情况，不需要单独判断，统一归到else里返回即可
	27/27 cases passed (24 ms)
	Your runtime beats 81.6 % of golang submissions
	Your memory usage beats 33.33 % of golang submissions (6.8 MB)
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 公共祖先
	lca := root
	for lca != nil {
		if p.Val < lca.Val && q.Val < lca.Val {
			lca = lca.Left
		} else if p.Val > lca.Val && q.Val > lca.Val {
			lca = lca.Right
		} else {
			return lca
		}
	}
	return lca
}

/*
	递归
	终止条件是两个节点不在同一颗子树

	27/27 cases passed (20 ms)
	Your runtime beats 96.48 % of golang submissions
	Your memory usage beats 33.33 % of golang submissions (6.8 MB)
	时间O(n)，空间O(n)，递归需要的栈空间依赖于两个节点的层数差
*/
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if p.Val < root.Val && q.Val < root.Val {
// 		// root即可，不需要再定义遍变量接收返回值
// 		return lowestCommonAncestor(root.Left, p, q)
// 	} else if p.Val > root.Val && q.Val > root.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	} else {
// 		return root
// 	}
// }

// @lc code=end
