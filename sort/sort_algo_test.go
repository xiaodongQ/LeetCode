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

	// 每次冒泡把最大的元素存到最后
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

// 插入排序
func TestInsertSort(t *testing.T) {
	data := [...]int{3, 5, 4, 1, 2, 6}
	log.Printf("init data:%v", data)

	if len(data) <= 1 {
		return
	}

	for i := 1; i < len(data); i++ {
		// 从未排序区取的数据
		value := data[i]
		j := i - 1

		// 从已排序序区的最后开始比较
		for ; j >= 0; j-- {
			if data[j] > value {
				// 已排序区数据比要插入的数据大则进行后移
				data[j+1] = data[j]
			} else {
				// 小于要插入的数据，则退出循环，并将数据插入到这个数据后面
				break
			}
		}
		data[j+1] = value
	}
	log.Printf("sort data:%v", data)
}

// 选择排序
func TestSelectionSort(t *testing.T) {
	data := [...]int{3, 5, 4, 1, 2, 6}
	log.Printf("init data:%v", data)

	if len(data) <= 1 {
		return
	}

	for i := 0; i < len(data); i++ {
		// 从未排序区找最小的元素
		minIndex := i
		for j := i; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			// 交换位置
			data[i], data[minIndex] = data[minIndex], data[i]
		}
	}
	log.Printf("sort data:%v", data)
}
