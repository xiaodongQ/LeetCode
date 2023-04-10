package backtrack

import (
	"log"
	"math"
	"testing"
)

var maxW = math.MinInt32 //存储背包中物品总重量的最大值

func TestBag(t *testing.T) {
	// a所有物品之和超出包的限制重量才有意义，要不全部都能装进去，不存在选择问题
	// 不用管a内的顺序如何，因为下面比的是已存入的总数量，每个数组元素索引从小到大都会回溯到(剪枝之外)，且不会重复
	// a := []int{8, 20, 34, 45, 51, 62, 73, 81, 90, 92}
	// a := []int{10, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	a := []int{55, 66, 10, 11, 22, 33, 44, 77, 88, 99}
	// a := []int{1, 2, 3, 4, 5}
	// 10个物品，选择装入限重100的包里，找出能放如最大的总数量
	f(0, 0, a, 10, 100)
	log.Printf("maxW:%d", maxW)
}

// cw表示当前已经装进去的物品的重量和；i表示考察到哪个物品了；
// w背包重量；items表示每个物品的重量；n表示物品个数
// 假设背包可承受重量100，物品个数10，物品重量存储在数组a中，那可以这样调用函数：
// f(0, 0, a, 10, 100)
func f(i int, cw int, items []int, n, w int) {
	if cw == w || i == n { // cw==w表示装满了;i==n表示已经考察完所有的物品
		if cw > maxW {
			maxW = cw
		}
		return
	}
	f(i+1, cw, items, n, w)
	if cw+items[i] <= w { // 已经超过可以背包承受的重量的时候，就不要再装了
		f(i+1, cw+items[i], items, n, w)
	}
}
