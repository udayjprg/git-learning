package packageone

import "fmt"

// privateVar can not be imported as its scope is within the packageone.
var privateVar= "Madhav"

// PackageVar can be imported with in the project by importing packageone as it starts with Uppercase letter
var PackageVar= "Mothadaka"


//PrintMe vavilable inside and outside the package
func PrintMe(s1, s2 string){
	
	fmt.Println(s1, s2, PackageVar)
}