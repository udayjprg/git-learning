package main

import "fmt"

func main(){
	i:=0
	sum:=0
	for (i<100){
		sum = sum+i
		sum=i+sum
		i+=1
	}
	fmt.Println(sum)
}