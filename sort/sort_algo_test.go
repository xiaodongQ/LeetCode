package sort

import (
	"log"
	"testing"
	"time"
)

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	data := [...]int{3, 5, 4, 1, 2, 6}
	log.Printf("init data:%v", data)
	num := 0

	for i := 0; i < len(data); i++ {
		// 注意此处的终止条件，并不需要len(data)-1，而是再-i，每次冒泡后，后面的冒泡次数的数不需再判断
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
			num++
		}
	}
	log.Printf("sort data:%v, num:%d", data, num)
}

// 冒泡排序优化，某次没有数据交换则说明已经完全有序
func TestBubbleSortOptimize(t *testing.T) {
	data := [...]int{3, 5, 4, 1, 2, 6}
	log.Printf("init data:%v", data)
	num := 0
	bHasSwap := false

	// 定义一个函数(其返回值也是一个func())并defer调用
	// 避免了defer传参时送time.Now()，但是可读性差
	defer func() func() {
		initTime := time.Now()
		// 没有等待的话打印出来就是0s
		time.Sleep(time.Nanosecond)
		return func() {
			log.Printf("cost:%v", time.Since(initTime))
		}
	}()()

	for i := 0; i < len(data); i++ {
		bHasSwap = false
		// len(data)-1-i 避免已经冒泡的几个数再做不必要的判断
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				bHasSwap = true
			}
			num++
		}
		if !bHasSwap {
			break
		}
	}
	log.Printf("sort data:%v, num:%d", data, num)
}
