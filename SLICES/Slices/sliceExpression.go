package main

import (
	"fmt"
	"strings"
)

func main(){
	// a:=[4]int{1,2,3,4}

	// b:=a[3:4]
	// fmt.Println(b)

	// c:=[]int{1,2,3,4,5}
	// cd:=c[0:3]
	// fmt.Println(cd)

	// arry1:=[4]int{1,2,3,4}
	// sliceA:=[]int{}
	// sliceA=arry1[0:3]
	// fmt.Println(sliceA, arry1)
	// sliceA=append(sliceA, 6)
	 fmt.Println(strings.Repeat("-", 10))
	// fmt.Println(sliceA, arry1)

	// sliceB:=[]int{1,2,3,4,5,6,7}
	// sliceC:=[]int{}
	// sliceC=sliceB[0:4]
	//fmt.Println("SliceC values are:",sliceC)
	// sliceC=append(sliceC, 100,200)
	// fmt.Println("SliceC values are:",sliceC,  "SliceB values are:",sliceB)

	
	// fmt.Println("SliceB values are:",sliceB)
	// sliceB[2]=66
	// fmt.Println(strings.Repeat("-",7))
	// fmt.Println("SliceC values are:",sliceC)
	// fmt.Println("SliceB values are:",sliceB)

	xyz:=[]int{1,2,3}
	fmt.Println("XYZ values are:", xyz)
	xy:=[]int{}
	xy=append(xy, xyz[0:2]...)
	fmt.Println("slice XY Append values are:",xy)
	xyz[2]=66
	fmt.Println("slice XYZ Changing values are:",xyz)
	fmt.Println("slice XY values are:",xy)

	fmt.Println(strings.Repeat("-", 20))

	sliceOne:=[]int{1,2,3}
	fmt.Println("sliceOne values are:",sliceOne)
	sliceTwo:=[]int{}
	sliceTwo=sliceOne[0:2]
	fmt.Println("sliceTwo from Slice Expression values are:",sliceTwo)
	//fmt.Println("sliceOne values are:",sliceOne)
	sliceOne[0]=0
	fmt.Println("sliceOne changingIndex 0 values are:",sliceOne)
	fmt.Println("slicetwo values are:",sliceTwo)

	fmt.Println(strings.Repeat("-", 20))

	slice1:=[]int{4,5,6,88,99}
	fmt.Println("slice1 values are:",slice1)
	slice2:=[]int{}
	slice2=append(slice2, slice1[0:5]...)
	fmt.Println("slice2 append values are:",slice2)
	//slice1=append(slice1,7,8,9)
	slice1[0]=0
	//fmt.Println("slice1 Append values are:",slice1)
	fmt.Println("slice1 values are:",slice1)
	fmt.Println("slice2 append values are:",slice2)
	fmt.Println(strings.Repeat("-", 20))

}