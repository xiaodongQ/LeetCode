package backtrace

import (
	"fmt"
	"testing"
)

var result = make([]int, 8)
var num = 0

/*
	8皇后问题：有一个 8x8 的棋盘，往里放 8 个棋子（皇后），每个棋子所在的行、列、对角线都不能有另一个棋子
 	递归尝试每行放置的位置是否符合条件
*/
func TestCal8queens(t *testing.T) {
	cal8queens(0)
	fmt.Printf("total num:%d\n", num)
}
func cal8queens(row int) {
	if row == 8 {
		printQueens(result)
		num++
		return
	}
	for column := 0; column < 8; column++ { // 每一行都有8中放法
		if isOk(row, column) { // 有些放法不满足要求
			result[row] = column // 第row行的棋子放到了column列
			cal8queens(row + 1)  // 考察下一行
		}
	}
}

func isOk(row, column int) bool { //判断row行column列放置是否合适
	leftup := column - 1
	rightup := column + 1
	for i := row - 1; i >= 0; i-- { // 逐行往上考察每一行
		if result[i] == column {
			return false // 第i行的column列有棋子吗？
		}
		if leftup >= 0 { // 考察左上对角线：第i行leftup列有棋子吗？
			if result[i] == leftup {
				return false
			}
		}
		if rightup < 8 { // 考察右上对角线：第i行rightup列有棋子吗？
			if result[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func printQueens(result []int) { // 打印出一个二维矩阵
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
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
