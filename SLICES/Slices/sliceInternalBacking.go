package main

import (
	"fmt"
	"strings"
	//"unsafe"
)

func main(){

	var slice1[]int
	fmt.Printf("%#v\n",slice1)
	slice1=append(slice1, 1,2)
	fmt.Printf("Length: %d, Capacity: %d \n" ,len(slice1), cap(slice1))
	slice1=append(slice1, 3)
	fmt.Printf("Length: %d, Capacity: %d \n" ,len(slice1), cap(slice1))
	fmt.Println(slice1)
	slice1=append(slice1, 4,5)
	fmt.Printf("Length: %d, Capacity: %d \n" ,len(slice1), cap(slice1))
	slice1=append(slice1[0:3],60,70,80,90)
	fmt.Println(slice1)
	fmt.Printf("Length: %d, Capacity: %d \n" ,len(slice1), cap(slice1))
	slice1=append(slice1, 22,23)
	fmt.Printf("Length: %d, Capacity: %d \n" ,len(slice1), cap(slice1))
	fmt.Println(slice1)
	
	// // elements of this slice will be stored in Backing Array
	// slice1=slice1[5:7]
	// fmt.Println(slice1)

	fmt.Println(strings.Repeat("-",16))

	// slice2:=[]int{}
	// fmt.Printf("%#v\n",slice2)

	array1:=[5]int{1,2,3,4,5}
	fmt.Println(array1)
	//array1=append(array1[0:3], 6,7)
	slice3:=[]int{}
	slice3=append(slice3, array1[0:3]...)
	fmt.Println(slice3)

		// printing memory in bytes
		// array2:=[6]int{1,2,2,3,4,6}
		// slice5:=[]int{1,2,2,3,4,6}
		// fmt.Printf("Array's size in bytes: %d \n", unsafe.Sizeof(array2))
		// fmt.Printf(" Slice's size in bytes: %d \n", unsafe.Sizeof(slice5))

		n1:=[]int{10,20,30,40}
		n1=append(n1,100)
		fmt.Println(len(n1), cap(n1))




}