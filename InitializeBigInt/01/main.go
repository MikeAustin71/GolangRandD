package main

// https://golang.org/pkg/math/big/

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(int64(55))

	fmt.Println(a)
}
