package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	// var word string
	// fmt.Scan(&word)
	fmt.Println("Hello, World.")
	fmt.Println(word)
}
