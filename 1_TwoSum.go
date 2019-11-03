/*
 * @Description:
 * @Author: xd
 * @Date: 2019-10-29 08:42:51
 * @LastEditTime: 2019-11-03 19:41:50
 * @LastEditors: xd
 * @Note:
	source: https://leetcode-cn.com/problems/two-sum/
	Given an array of integers, return indices of the two numbers such that they add up to a specific target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
*/

package main

import (
	"fmt"
	"net"
)

// 结果：执行用时:4 ms, golang中>97.39%; 内存消耗:3.7MB, golang中>46.28%
// 时间：O(n) 空间：O(n)
func twoSum(nums []int, target int) []int {
	mapInt := make(map[int]int)
	for index, v := range nums {
		iNum, ok := mapInt[target-v]
		if ok {
			return []int{iNum, index}
		}
		mapInt[v] = index
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	res := twoSum(nums, target)

	for _, v := range res {
		fmt.Printf("index:%d, value:%d", v, nums[v])
		return
	}
	_, err := net.ResolveTCPAddr("tcp", ":4040")
	fmt.Println()
}
