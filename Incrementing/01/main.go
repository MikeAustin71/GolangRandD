package main

import (
	"fmt"
	"math/big"
)

// num++ doesn't work with big.Int

func main() {

	base := big.NewInt(0)
	r1 := incrementOp(base)
	fmt.Println("0 plus 20 using incrment operator", r1)
	fmt.Println("base:", base, "  result:", r1)
}

func incrementOp(num *big.Int) *big.Int {

	n2 := big.NewInt(0).Set(num)

	increment := big.NewInt(1)

	for i := 0; i < 20; i++ {

		n2.Add(n2, increment)
	}

	return n2
}
