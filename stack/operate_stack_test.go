package main

import (
	"log"
	"reflect"
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

func (s *ArrStack) Init(input interface{}, arrcap int) bool {
	log.Printf("type:%v\n", reflect.TypeOf(input).String())
	elements, ok := input.([]byte)
	if !ok {
		log.Printf("init stack type error, input:%+v!\n", input)
		return false
	}
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

/*
* 有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
* 如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
* 给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
* 对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S
* e.g. 输入："(()())(())"，输出："()()()"
	原语化分解得到 "(()())" + "(())"
	删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"
* 说明：S.length <= 10000
	S[i] 为 "(" 或 ")"
	S 是一个有效括号字符串
* [1021. 删除最外层的括号](https://leetcode-cn.com/problems/remove-outermost-parentheses/)
*/
func TestRemoveOuterParentheses(t *testing.T) {
	s := &ArrStack{}
	s.Init([]byte("(()())(())"), 100)
	s.PrintWithMark("init")

	if s.Push(888) {
		s.PrintWithMark("after push")
	}

	log.Printf("pop:%v\n", s.Pop())
	s.PrintWithMark("after pop")

}
func removeOuterParentheses(S string) string {
	return ""
}
