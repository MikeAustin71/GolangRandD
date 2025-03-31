package main

import (
	"fmt"
	"math/big"
)

func main() {
	five := big.NewInt(5)
	a := big.NewInt(35)

	b := big.NewInt(0).Set(a)
	b = b.Add(b, five)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
