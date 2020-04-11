package main

import (
	"fmt"
	"unsafe"
)

func main() {
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
}
