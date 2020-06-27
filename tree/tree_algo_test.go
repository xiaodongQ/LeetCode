package tree

import (
	"fmt"
	"log"
	"testing"
)

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

var globalTree *BinaryTreeNode

func initTree() {
	// 清空节点
	globalTree = &BinaryTreeNode{}
	// 定义几个节点
	// 第2层(根节点为第一层)
	node1 := BinaryTreeNode{
		value: 6,
	}
	node2 := BinaryTreeNode{
		value: 15,
	}

	// 第3层
	node3 := BinaryTreeNode{
		value: 0,
	}
	node4 := BinaryTreeNode{
		value: 8,
	}
	// 构造一个完全二叉树(叶子节点都在最底下两层，最后一层的叶子节点都靠左，且除最后一层，其他层节点达到最大)
	globalTree.value = 9
	globalTree.left = &node1
	globalTree.right = &node2
	// root的左子树的左右节点
	node1.left = &node3
	node1.right = &node4
	/*
		树结构如(同时也是二叉查找树，任意节点左子树中每个节点的值都小于这个值，右子树中节点值都大于这个值)：
				 9
			   6   15
			  0 8
	*/
}

// 遍历二叉树
func TestTreeTraverse(t *testing.T) {
	initTree()
	preOrder(globalTree) // 9 6 0 8 15
	fmt.Println(" preOrder")

	inOrder(globalTree) // 0 6 8 9 15
	fmt.Println(" inOrder")

	postOrder(globalTree) // 0 8 6 15 9
	fmt.Println(" postOrder")

	levelOrder(globalTree) // 9 6 15 0 8
	fmt.Println(" levelOrder")
}

// 前序遍历
func preOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	// log.Printf打印默认会格式化每行后换行，此处打印到一行里，最后打印一次换行刷新缓存区
	fmt.Printf("%d ", root.value)
	preOrder(root.left)
	preOrder(root.right)
}

// 中序遍历
func inOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("%d ", root.value)
	inOrder(root.right)
}

// 后序遍历
func postOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%d ", root.value)
}

// 层序遍历(借助slice模拟队列)
func levelOrder(root *BinaryTreeNode) {
	var s1 []*BinaryTreeNode

	s1 = append(s1, root)
	for len(s1) != 0 {
		fmt.Printf("%d ", s1[0].value)
		if s1[0].left != nil {
			s1 = append(s1, s1[0].left)
		}
		if s1[0].right != nil {
			s1 = append(s1, s1[0].right)
		}
		// 从队列取出一个节点(模拟队列的FIFO)
		s1 = s1[1:]
	}
}

// 查找(二叉查找树的查找)
func TestFind(t *testing.T) {
	/*
			9
		  6   15
		 0 8
	*/
	initTree()
	v := 8
	node := Find(globalTree, v)
	if node == nil {
		log.Printf("not find")
	} else {
		log.Printf("find node:%p, value:%d", node, node.value)
	}

	node = FindRec(globalTree, v)
	if node == nil {
		log.Printf("not find")
	} else {
		log.Printf("find node:%p, value:%d", node, node.value)
	}
}

// 递归方式
func FindRec(root *BinaryTreeNode, value int) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	if value < root.value {
		return Find(root.left, value)
	} else if value > root.value {
		return Find(root.right, value)
	}
	return root
}

// 循环方式(递推)
func Find(root *BinaryTreeNode, value int) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		if value < node.value {
			node = node.left
		} else if value > node.value {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

func TestInsert(t *testing.T) {
	/*
			9
		  6   15
		 0 8
	*/
	initTree()
	levelOrder(globalTree) // 9 6 15 0 8
	fmt.Println(" levelOrder")

	InsertRec(globalTree, 5)
	levelOrder(globalTree) // 9 6 15 0 8
	fmt.Println(" levelOrder")
}

// 不考虑元素相同
func Insert(root *BinaryTreeNode, v int) {
	node := root
	for node != nil {
		if v < node.value {
			// 若要插入的元素比节点值小，其左子树为空时则直接插入，不为空则进入下一次递归
			if node.left == nil {
				node.left = &BinaryTreeNode{value: v}
				return
			}
			node = node.left
		} else {
			if node.right == nil {
				node.right = &BinaryTreeNode{value: v}
				return
			}
			node = node.right
		}
	}
}
func InsertRec(root *BinaryTreeNode, v int) {
	if root == nil {
		return
	}
	if v < root.value {
		if root.left == nil {
			root.left = &BinaryTreeNode{value: v}
			return
		}
		InsertRec(root.left, v)
		return
	}

	if root.right == nil {
		root.right = &BinaryTreeNode{value: v}
		return
	}
	InsertRec(root.right, v)
}
