package tree

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
* 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
* [108. 将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/)

* 可以选择中间数字作为二叉搜索树的根节点，这样分给左右子树的数字个数相同或只相差 11，可以使得树保持平衡。
 如果数组长度是奇数，则根节点的选择是唯一的，
 如果数组长度是偶数，则可以选择中间位置左边的数字作为根节点或者选择中间位置右边的数字作为根节点，选择不同的数字作为根节点则创建的平衡二叉搜索树也是不同的。

 确定平衡二叉搜索树的根节点之后，其余的数字分别位于平衡二叉搜索树的左子树和右子树中，左子树和右子树分别也是平衡二叉搜索树，
 因此可以通过递归的方式创建平衡二叉搜索树

 为什么这么建树一定能保证是「平衡」的呢？
 这里可以参考「1382. 将二叉搜索树变平衡」，这两道题的构造方法完全相同，这种方法是正确的，
 1382 题解中给出了这个方法的正确性证明：
 [1382. 将二叉搜索树变平衡 官方题解](https://leetcode-cn.com/problems/balance-a-binary-search-tree/solution/jiang-er-cha-sou-suo-shu-bian-ping-heng-by-leetcod/)

*/
func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	// nums := []int{}
	LevelOrder(sortedArrayToBST(nums))
	// LevelOrder(sortedArrayToBSTLoop(nums))

}

func sortedArrayToBST(nums []int) *TreeNode {
	/* 递归方式
	执行用时：4 ms, 在所有 Go 提交中击败了86.31%的用户
	内存消耗：4.4 MB, 在所有 Go 提交中击败了28.57%的用户
	*/
	return doArraySort(nums, 0, len(nums)-1)
}
func doArraySort(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 不使用 (left+right)/2，有溢出风险
	mid := left + ((right - left) >> 1)
	root := &TreeNode{Val: nums[mid]}
	root.Left = doArraySort(nums, left, mid-1)
	root.Right = doArraySort(nums, mid+1, right)
	return root
}

//执行用时：4 ms, 在所有 Go 提交中击败了86.31%的用户
// 内存消耗：5.8 MB, 在所有 Go 提交中击败了28.57%的用户
func sortedArrayToBSTLoop(nums []int) *TreeNode {
	// 循环方式，利用两个队列，类似二叉树层序遍历的方式(BFS宽度优先搜索)
	if len(nums) == 0 {
		return nil
	}
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	if len(nums) == 1 {
		return root
	}

	// 保存下次要处理的索引范围，分别对应左右子树
	srange := make([][2]int, 0, 64)
	srange = append(srange, [2]int{left, mid - 1})
	srange = append(srange, [2]int{mid + 1, right})
	// 记录父节点，记录两次，分别用于左、右子树的处理
	// 和保存的范围个数是一一对应的
	snode := make([]*TreeNode, 0, 64)
	snode = append(snode, root)
	snode = append(snode, root)
	for len(srange) != 0 && len(snode) != 0 {
		// 取出一次父节点，取出后就去掉一条记录，模拟队列
		curnode := snode[0]
		snode = snode[1:]

		// 取出一个范围
		left = srange[0][0]
		right = srange[0][1]
		srange = srange[1:]

		// 注意不能像下面两种方式直接用 >>1 来替换/2，若left和right相等，则右移一位等于-1，会影响后面的判断并导致越界访问
		// mid = ((right + left) >> 1)
		// mid = left + ((right - left) >> 1)
		mid = left + ((right - left) / 2) // 还是老实用/2吧
		// 创建一个新节点
		newnode := &TreeNode{Val: nums[mid]}
		if newnode.Val < curnode.Val {
			curnode.Left = newnode
		} else {
			curnode.Right = newnode
		}

		// 把新节点和范围加到队列中
		if left < right {
			srange = append(srange, [2]int{left, mid - 1})
			srange = append(srange, [2]int{mid + 1, right})
			snode = append(snode, newnode)
			snode = append(snode, newnode)
		}
	}

	return root
}
