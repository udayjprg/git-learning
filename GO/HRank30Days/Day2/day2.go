package main

import "fmt"

func main() {
	var mealCost float64
	var tax int
	var tip int
	//talCost(mealCost float64, taxPercent int, tipPercent int){
	fmt.Scan(&mealCost)
	fmt.Scan(&tax)
	fmt.Scan(&tip)
	
	var taxPercent float64 
	taxPercent=float64(mealCost/100)
	var x float64
	x=float64(tax)*taxPercent
	var tipPercent float64
	 tipPercent = (mealCost/100)
	var y float64
	y=float64(tip)*tipPercent
	totalMealCost:=mealCost + x + y
	fmt.Printf("TotalMealCost:%0.f \n",totalMealCost)
	fmt.Println("MealCost",mealCost)
	fmt.Println("TaxPercent",x)
	fmt.Println("TipPercent",y)

}


