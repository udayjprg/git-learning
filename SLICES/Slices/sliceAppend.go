package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main(){
	fmt.Println(strings.Repeat("-",10))

	//abc:=[]int{1,2,3,4,5}
	//c:=[]int{1,2}
	//c =append(c, abc...)

	a:=make([]int, 9)
	fmt.Println(a)
	fmt.Printf("Length of Slice is %d, Capacity of Slice is %d  \n ",len(a),cap(a))
	fmt.Printf("Length of Slice in Bytes is: %d \n", unsafe.Sizeof(a))
	fmt.Println(strings.Repeat("-",10))
	// d:=make([]int, len(abc))
	// n:=copy(d, abc)
	//m:=copy(a,abc)
	//fmt.Println(d,n,abc)
}