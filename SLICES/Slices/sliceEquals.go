package main

import (
	"fmt"
	"strings"
)

func main(){
	// var n[]int
	// n=[]int{1,2}
	// fmt.Println(strings.Repeat("-", 6))
	// fmt.Println(n==nil)
	// fmt.Println(strings.Repeat("-", 12))
	// m:=[]int{}
	// fmt.Println(m==nil)
	fmt.Println(strings.Repeat("-", 20))

	a,b := []int {1,22,3}, []int{1,2,3}
	//fmt.Println(a==b)
	var eq bool=true
	for i, value:=range a{
		if value !=b[i]{
			eq=false
			break
		}
	}
	if eq {
		fmt.Println(" a & b are equal")
	}else{
		fmt.Println(" a & b are not equal")
	}

}