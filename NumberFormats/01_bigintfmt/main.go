package main

import (
	"fmt"
	"math/big"
)

func main()  {


	fmt.Println("big int format test")

	num,_ := big.NewInt(0).SetString("0100", 10)
	fmt.Println("********* Test # 1 *************")
	fmt.Println("       Number: 0100")
	fmt.Println("String Result:", num.String())
	fmt.Println("********************************")
	fmt.Println("")
	thousand := big.NewInt(1000)
	num2 := big.NewInt(0).Div(num, thousand)

	fmt.Println("********* Test # 2 *************")
	fmt.Println("       Number: 0.1")
	fmt.Println("String Result:", num2.String())
	fmt.Println("********************************")
	fmt.Println("")


}

