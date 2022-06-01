package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var Arry = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Arry[i])
	}
	for j := 0; j < N; j++ {
		j++
		fmt.Printf(" %#v ", Arry[N-j])
		j--
	}
	fmt.Println(" ")
}
