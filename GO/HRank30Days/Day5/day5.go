package main

import "fmt"

func main() {
	var n int
	//var x int
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		//x=n*i
		fmt.Printf("%d * %d = %d \n", n, i, n*i)
	}
}
