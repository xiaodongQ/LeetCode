package backtrack

import (
	"log"
	"testing"
)

var matched = false
var pattern = []byte("*c*")
var plen = len(pattern)

// 若直接 go test regex_match_test.go，如果所有用例都pass，则下面的打印不会打印出来，通过-v可以打印过程
// 若没有pass，则下面打印会进行打印(没有加-v的情况下也会)
func TestRegMatch(t *testing.T) { // 文本串及长度
	text := []byte("abc123hh")
	tlen := len(text)
	matched = false
	rmatch(0, 0, text, tlen)
	log.Printf("match:%t", matched)
	log.Fatal("xx")
}

func rmatch(ti, pj int, text []byte, tlen int) {
	if matched {
		return // 如果已经匹配了，就不要继续递归了
	}
	if pj == plen { // 正则表达式到结尾了
		if ti == tlen {
			matched = true // 文本串也到结尾了
		}
		return
	}
	if pattern[pj] == '*' { // *匹配任意个字符
		for k := 0; k <= tlen-ti; k++ {
			rmatch(ti+k, pj+1, text, tlen)
		}
	} else if pattern[pj] == '?' { // ?匹配0个或者1个字符
		rmatch(ti, pj+1, text, tlen)
		rmatch(ti+1, pj+1, text, tlen)
	} else if ti < tlen && pattern[pj] == text[ti] { // 纯字符匹配才行
		rmatch(ti+1, pj+1, text, tlen)
	}
}
