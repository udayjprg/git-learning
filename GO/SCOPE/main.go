package main

import (
	//"fmt"
	"myapp/packageone"
)

//myVar is a package level variable so it can be used by any func within the package
var myVar="Hello"

func main(){

	// blockVar is a block level variable so it is used within the main funtion
	var blockVar = "Uday"
	//fmt.Println(one)

//newString :=packageone.PackageVar
packageone.PrintMe(myVar, blockVar)
//fmt.Println()
//fmt.Println(blockVar)

}

