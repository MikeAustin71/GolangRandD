package main

import (
	"fmt"
	"math/big"
)

func main() {

	total := big.NewInt(0)

	addSomeNums(total)

	fmt.Println("The total is ", total)
}

func addSomeNums(total *big.Int) {

	increment := big.NewInt(1)
	limit := big.NewInt(30000)

	for i := big.NewInt(0); i.Cmp(limit) < 0; i.Add(i, increment) {

		total.Add(total, increment)
	}
}
