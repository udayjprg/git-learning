package main

import "fmt"

func main(){

	//Create a Slice From an Array

	arr1 :=[8]int{1,2,3,4,5,6}

	slice1 :=arr1[1:2]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}