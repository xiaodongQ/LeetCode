/*
 * @Description:
 * @Author: xd
 * @Date: 2020-06-07 22:26:06
 * @LastEditTime: 2020-06-07 22:44:45
 * @LastEditors: xd
 * @Note:
 */
package search

import (
	"log"
	"testing"
)

func TestBsearch(t *testing.T) {
	data := []int{1, 2, 3, 6, 7}
	log.Printf("data:%v", data)
	value := 6
	log.Printf("loop find %d, index:%d\n", value, bsearch(data, value))
	log.Printf("recursive find %d, index:%d\n", value, bsearch2(data, value))
}

// 循环实现
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

// 递归实现
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
