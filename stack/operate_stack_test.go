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
		arrcap = len(elements)
	}
	s.data = make([]interface{}, 0, arrcap)
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

func (s *ArrStack) PrintWithMark(mark string) {
	log.Printf("\n%s, stack: %+v\n", mark, s.data)
}
func (s *ArrStack) Print() {
	log.Printf("\nstack: %+v\n", s.data)
}

func (s *ArrStack) Push(element interface{}) bool {
	// 栈满
	if s.count >= s.cap {
		log.Printf("stack full[cap:%d], push error!", s.cap)
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
		log.Printf("stack empty, pop error!")
		return nil
	}
	// 将slice的最后一个元素删除
	res := s.data[s.top]
	s.data = s.data[:s.count-1]
	s.count--
	s.top = s.count - 1
	return res
}

func TestArrStackOperate(t *testing.T) {
	s := &ArrStack{}
	s.Init([]interface{}{1, 2, 3}, 10)
	s.PrintWithMark("init")

	if s.Push(888) {
		s.PrintWithMark("after push")
	}

	log.Printf("pop:%v\n", s.Pop())
	s.PrintWithMark("after pop")

}
