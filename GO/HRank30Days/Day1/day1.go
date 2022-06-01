package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var i uint64 = 4
	var d float64 = 4.0
	var s string ="hacker Rank"
	scanner:= bufio.NewScanner(os.Stdin)

	var x uint64
	var y float64

	fmt.Scan(&x, &y)
	
	for scanner.Scan(){
		fmt.Println(s + scanner.Text())
	}
	
	fmt.Println(i + x)
	fmt.Println(d + y)
	
}