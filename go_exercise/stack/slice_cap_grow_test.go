package stack

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMain(t *testing.T) {
	fmt.Println("vim-go")
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))

	f := []int{1, 2, 3}
	fmt.Println("cap of f before:", cap(f))
	f = append(f, 4)
	fmt.Println("cap of f after:", cap(f))

	var a int
	var b int32
	fmt.Printf("int len:%d, int32 len:%d\n", unsafe.Sizeof(a), unsafe.Sizeof(b))

	s1 := []int{1, 3, 5, 7}
	s1 = append(s1[0:2], s1[3:]...)
	fmt.Printf("s1:%v\n", s1)
}
