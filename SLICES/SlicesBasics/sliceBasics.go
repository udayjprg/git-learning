package main

import (
	 "fmt"
	 "strings"
)

func main(){
	// var abc []string
	// fmt.Println(abc)

	// //abc[0]="uday"

	// fmt.Println(abc)
	fmt.Println(strings.Repeat("-", 20))
	//fmt.Printf("%#v\n", abc)

	abs := make([]int, 2)
		fmt.Println(len(abs))
		abs[0]=10
		fmt.Println(abs)
		fmt.Println(strings.Repeat("-",20))
		fmt.Printf("abs %#v\n", abs)

		// Another way to create Slice using Make function
		nums :=make([] int,2)
		fmt.Printf("%#v\n",nums)
}