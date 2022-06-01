package main

import "fmt"

func main(){
fmt.Printf("%#v\n", reverse([]int{1,2,3}))
}
func reverse(numbers []int)[]int{
	for i:=0; i<len(numbers)-1; i++{
		j:=len(numbers)-i-1
		numbers[i], numbers[j]=numbers[j], numbers[i]
	}
	return numbers
}