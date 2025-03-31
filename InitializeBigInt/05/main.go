package main

// https://golang.org/pkg/math/big/

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(55)
	b := big.NewInt(0)

	fmt.Println(a)
	a = b
	fmt.Println(a)
}
