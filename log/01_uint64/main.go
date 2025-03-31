package main

import (
	"fmt"
	"math"
)

func main() {
	result := math.Log(800)

	fmt.Println("This is the natural log of 800 ", result)
	// constant e = math.E
	result2 := math.Pow(math.E, result)
	fmt.Println(math.E, "^", result, " = ", result2)
	fmt.Println("-----------------------------------------")

	result3 := math.Log10(800)
	fmt.Println("This is log10 of 800 ", result3)

	fmt.Println("10^", result3, " = ", math.Pow(10, result3))
}
