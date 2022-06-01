package main

import "fmt"

func main() {
	var n int32
	fmt.Scan(&n)
	if n%2==1 || (n>5 && n<21) {
		fmt.Println("Weird")
	}else{
		fmt.Println("Not Weird")
	}
	}



	//**********************************
	// 	if n%2 == 0 {
	// 		if n>=2 && n <= 5 {
	// 	fmt.Println("Not Weird")
	// } else if  n >= 6 && n <= 20 {
	// 	fmt.Println(" Weird")
	// } else if n > 20 {
	// 	fmt.Println("Not Weird1")
	// }
	// }else{
	// 	fmt.Println("Weird")
	// }
	//**********************************

	


