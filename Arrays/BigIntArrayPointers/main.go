package main

import (
	"fmt"
	"math/big"
)

func main() {
	upperLimit := int64(20)
	biArray := createBigIntArray(upperLimit)
	fmt.Println("This is the Big Int Array")
	fmt.Println(biArray)
}

func createBigIntArray(upperLimit int64) []*big.Int {

	bintArray := make([]*big.Int, upperLimit)

	for i := int64(1); i <= upperLimit; i++ {
		bintArray[i-1] = big.NewInt(i)
	}

	return bintArray
}
