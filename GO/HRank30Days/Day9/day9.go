package main

import "fmt"

//Recursion

func fib(n int) int {
	//fmt.Scan(&n)
	if n == 0 {
		return 0
	} else if n == 1 {

		fmt.Printf("Hi\n")
		return 1
	} else {
		return fib(n - 1)
	}
}

func main() {
	fmt.Println(fib(5))
}
