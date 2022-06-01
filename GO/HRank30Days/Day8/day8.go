package 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//creating variable N
	//var N int
	// taking inputs for N
	//fmt.Scan(&N)
	//Map1 --> ***********************************************
	// // creating a map1 using make function
	// var map1 = make((map[string]int))
	// // assigning key, values to map1
	// map1["Uday"] = 2356898754
	// map1["Madhav"]=1245789865
	// map1["Cali"]=5285963641
	// // map1[4]="Sanjose"
	// //map1=append(map1,5)
	// fmt.Println("Map1: ",map1["Cali"])
	// fmt.Println("Map1: ", map1["Sanjose"])
	// name, _:= map1["Sanjose"]
	// 	//fmt.Println("\n key present or not: ", ok)
	// 	fmt.Println("Value: ", name)
	// Map2 --> ************************************************
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	map2 := make(map[string]string)
	
	for k := 0; k < n; k++ {
		scanner.Scan()
		map2[strings.Split(scanner.Text(), " ")[0]] = strings.Split(scanner.Text(), " ")[1]
	}
	for scanner.Scan() {
		if map2[scanner.Text()] == "" {
			fmt.Println("Not Found")
		} else {
			fmt.Printf("%s = %s\n", scanner.Text(), map2[scanner.Text()])
		}
	}
}

//	}
//fmt.Println(name)
//fmt.Println("Map2: ", map2)

// 	fmt.Println("Map2: ", map2)
