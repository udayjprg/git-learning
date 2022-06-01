package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i:=0; i<T; i++{
		var word string
		fmt.Scan(&word)
		wordFun(word)
		fmt.Println("")
	}
}
	
func wordFun(word string){
	for j:=0; j<len(word); j+=2{
	if j%2 == 0 {
	fmt.Printf("%c", word[j])
		} 
	}
		fmt.Printf(" ")
		for j:=1; j<len(word); j+=2{
		 if( j%2==1){
		fmt.Printf("%c", word[j])
		}
	}
}

//*********************************************


	// 	fmt.Printf("  ")
	// 	for i := 0; i < len(s); i++ {
	// 		if i%2 == 1 {
	// 			fmt.Printf("%c", s[i])
	// 		}
	// }
	// fmt.Println("")

	// for i := 0; i < len(d); i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%c", d[i])
	// 	}
	// }
	// 	fmt.Printf("  ")
	// 	for i := 0; i < len(d); i++ {
	// 		if i%2 == 1 {
	// 			fmt.Printf("%c", d[i])
	// 		}
	// }
	// fmt.Println("")

