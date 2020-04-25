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

func (s *ArrStack) Init(input []interface{}, stacksize int) bool {
	if stacksize == 0 {
		stacksize = len(input)
	}

	if len(input) == 0 || len(input) > stacksize {
		log.Fatalf("error! len:%d, in size:%d", len(input), stacksize)
		return false
	}
	s.cap = stacksize
	s.data = make([]interface{}, 0, stacksize)
	for _, v := range input {
		s.data = append(s.data, v)
		s.top = s.count
		s.count++
	}

	return true
}

func (s *ArrStack) Print(mark string) {
	log.Printf("%s, stack: %+v\n", mark, s.data)
}

func (s *ArrStack) Push(element interface{}) bool {
	// 栈满
	if s.count >= s.cap {
		log.Fatalf("stack full[cap:%d], push error!", s.cap)
		return false
	}
	s.data = append(s.data, element)
	// 栈顶更新
	s.top = s.count
	s.count++
	return true
}

func (s *ArrStack) Pop() interface{} {
	// 栈为空
	if s.count == 0 {
		log.Fatalf("stack empty, pop error!")
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
	s.Print("init")

	if s.Push(888) {
		s.Print("after push")
	}

	log.Printf("pop:%v\n", s.Pop())
	s.Print("after pop")

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
	log.Printf("req:%s, res:%s\n", "(()())(())", removeOuterParentheses("(()())(())"))
}

/*
* 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
* 内存消耗 :6.2 MB, 在所有 Go 提交中击败了50.00%的用户
* 时间O(n)，空间O(1)
 */
func removeOuterParentheses(S string) string {
	var res string
	// 截取位置
	start, end := 0, 0
	// 模拟栈中"("的个数
	num := 0
	// 是否开始下一段
	flag := false

	for k, v := range S {
		if v == '(' {
			num++
			if !flag {
				start = k
				flag = true
			}
		} else if v == ')' {
			num--
			if num == 0 {
				end = k
				res += S[start+1 : end]
				flag = false
			}
		}
	}
	return res
}
