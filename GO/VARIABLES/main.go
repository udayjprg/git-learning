package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt ="then dont type the number and press enter"
func main(){

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substract = rand.Intn(8) + 2
	var answer = firstNumber * secondNumber - substract

	integer(firstNumber, secondNumber, substract, answer)

	}

func integer(firstNumber, secondNumber, substract, answer int){	

	rand.Seed(time.Now().UnixNano())

	

	reader :=bufio.NewReader(os.Stdin)

	fmt.Println(" Think of a number between 1 to 10 ", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result with the number you originally thought of", prompt)
	reader.ReadString('\n')
	
	fmt.Println("Now Substract", substract, prompt)
	reader.ReadString('\n')
 
	fmt.Println("The Answer is :",answer)
}


	

