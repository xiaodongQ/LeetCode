/*
 * @Description:
 * @Author: xd
 * @Date: 2020-06-07 22:26:06
 * @LastEditTime: 2020-06-08 09:01:06
 * @LastEditors: xd
 * @Note:
 */
package search

import (
	"log"
	"testing"
)

func TestBsearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 4, 4, 5, 6, 6, 7}
	log.Printf("data:%v", data)
	value := 4
	log.Printf("loop find %d, index:%d\n", value, bsearch(data, value))
	log.Printf("recursive find %d, index:%d\n", value, bsearch2(data, value))

	log.Printf("search first member, find %d, index:%d\n", value, bsearchFirst(data, value))
}

// 二分查找 循环实现
func bsearch(data []int, value int) int {
	if len(data) == 0 {
		return -1
	}
	low := 0
	high := len(data) - 1
	for low <= high {
		// 可能溢出
		// mid := (low + high) / 2
		// mid := (low + high) >> 1
		mid := low + ((high - low) >> 1)
		if value == data[mid] {
			return mid
		} else if value < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找 递归实现
func bsearch2(data []int, value int) int {
	return bsearchRecusive(data, 0, len(data)-1, value)
}
func bsearchRecusive(data []int, low, high, value int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if value == data[mid] {
		return mid
	} else if value < data[mid] {
		return bsearchRecusive(data, low, mid-1, value)
	} else {
		return bsearchRecusive(data, mid+1, high, value)
	}
}

// 二分查找 查找第一个等于指定值的元素
func bsearchFirst(data []int, value int) int {
	if len(data) == 0 {
		return -1
	}
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value == data[mid] {
			if mid == 0 || data[mid-1] != value {
				return mid
			}
			high = mid - 1
		} else if value < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找 查找第一个等于指定值的元素 实现2
func bsearchFirst2(data []int, value int) int {
	if len(data) == 0 {
		return -1
	}
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value <= data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

		// 低位找到相等元素则返回
		if low < len(data) && data[low] == value {
			return low
		}
	}
	return -1
}
