package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()
	sint := "-123456654321"

	bSgnInt, isOk := big.NewInt(0).SetString(sint, 10)

	if !isOk {
		fmt.Println("Error from bigInt.SetSring(). Cannot convert %v", sint)
		return
	}

	fmt.Println("Original signed int string: ", sint)
	fmt.Println("                   bSgnInt: ", bSgnInt.String())

}


