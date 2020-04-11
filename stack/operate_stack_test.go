package main

import (
	"log"
	"testing"
)

type ArrStack struct {
	// 容量
	cap int
	// 大小
	count int
	// 栈顶(索引)
	top int
	// 栈数据
	data []interface{}
}

type ListGeneralNode struct {
	Val  interface{}
	Next *ListGeneralNode
}

type ListStack struct {
	// 容量
	size int
	// 栈顶
	top *ListGeneralNode
	// 栈数据
	data *ListGeneralNode
}

func (s *ArrStack) Init(elements []interface{}, arrcap int) bool {
	if arrcap == 0 {
		arrcap = 50
	}
	s.data = make([]interface{}, 0, arrcap)

	s1 := make([]int, 5)
	log.Printf("s1 len:%d, cap:%d\n", len(s1), cap(s1))
	s2 := make([]int, 5, 10)
	log.Printf("s2 len:%d, cap:%d\n", len(s2), cap(s2))
	var s3 []int
	log.Printf("s3 len:%d, cap:%d\n", len(s3), cap(s3))

	if len(elements) > arrcap {
		log.Printf("init stack error[input len:%d, cap:%d]!\n", len(elements), arrcap)
		return false
	}
	s.cap = arrcap
	for _, v := range elements {
		s.data = append(s.data, v)
		s.count++
	}
	return true
}

func (s *ArrStack) Print() {
	log.Printf("\nstack: %+v\n", s.data)
}

func (s *ArrStack) Push(element interface{}) bool {
	// 栈满
	if s.count == s.cap {
		return false
	}
	s.data[s.count] = element
	// 栈顶更新
	s.top = s.count
	s.count++
	return true
}

func (s *ArrStack) Pop() interface{} {
	// 栈为空
	if s.count == 0 {
		return nil
	}
	s.count--
	s.top = s.count - 1
	return s.data[s.count]
}

func TestArrStackOperate(t *testing.T) {
	s := &ArrStack{}
	s.Init([]interface{}{1, 2, 3}, 0)
	s.Print()
}
