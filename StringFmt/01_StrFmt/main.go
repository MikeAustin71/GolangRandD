package main

import (
	"fmt"
	"math/big"
)

func main()  {

	fmt.Println()

	n1 := big.NewInt(123456123456)

	fmt.Println("n1 String()", n1.String())

	base10 := big.NewInt(10)
	ex := big.NewInt(6)
	scaleFactor := big.NewInt(0).Exp(base10, ex, nil)
	yMod := big.NewInt(0)

	s1, s1Mod := big.NewInt(0).DivMod(n1,scaleFactor, yMod)

	fmt.Println("    Quotient: ", s1.String())
	fmt.Println("     Modulus: ", s1Mod.String())
	fmt.Println("Scale Factor: ", scaleFactor.String())

	fmt.Println("=====================================")
	fmt.Println("Text Quotient: ", s1.Text(10))
	fmt.Println(" Text Modulus: ", s1Mod.Text(10))



}

