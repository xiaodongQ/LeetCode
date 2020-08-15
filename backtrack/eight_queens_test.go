package backtrack

import (
	"fmt"
	"testing"
)

var result = make([]int, 8)
var num = 0

func TestCal8queens(t *testing.T) {
	cal8queens(0)
	// 打印出来共有92种放置方法
	fmt.Printf("total num:%d\n", num)
}

/*
	8皇后问题：有一个 8x8 的棋盘，往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子
 	递归尝试每行放置的位置是否符合条件
*/
func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		num++
		return
	}
	// 该行的每列都进行尝试，每次尝试都是一种放法(剩下的行递归下去)
	for column := 0; column < 8; column++ { // 每一行都有8中放法
		if isOk(row, column) { // 有些放法不满足要求
			result[row] = column // 第row行的棋子放到了column列
			cal8queens(row + 1)  // 考察下一行
		}
	}
}

// 判断row行column列放置是否合适，下标从0开始
// 每次只要判断上面行的数据，因为是按顺序放置下来的
func isOk(row, column int) bool {
	// 前一列，后续和行一起--，保持按对角线移动
	leftup := column - 1
	// 后一列
	rightup := column + 1
	// 逐行往上考察每一行
	for i := row - 1; i >= 0; i-- {
		// 查看同一列
		if result[i] == column {
			return false // 第i行的column列有棋子吗？
		}
		// 查看左对角线，过滤已经走过对角线左上顶点的后续情况(虽然leftup可能还在--)
		if leftup >= 0 { // 考察左上对角线：第i行leftup列有棋子吗？
			if result[i] == leftup {
				return false
			}
		}
		// 查看右对角线，过滤已经走过对角线右上最后一个的情况
		if rightup < 8 { // 考察右上对角线：第i行rightup列有棋子吗？
			if result[i] == rightup {
				return false
			}
		}
		// 左右对角线分别移动，有可能某个已经超出了对角线边缘，不过上面会有边界判断
		leftup--
		rightup++
	}
	return true
}

func printQueens(result []int) { // 打印出一个二维矩阵
	// 8行
	for row := 0; row < 8; row++ {
		// 8列
		for column := 0; column < 8; column++ {
			// 该行放的是这一列才打印皇后，否则为空位置，用"*"表示
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")

			}
		}
		fmt.Println()
	}
	fmt.Println()
}
