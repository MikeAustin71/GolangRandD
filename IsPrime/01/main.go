package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("Test For Prime")
	fmt.Println("------------------------------------------------")

	var testNums []*big.Int

	testNums = append(testNums, big.NewInt(4))
	//testNums = append(testNums, big.NewInt(7))
	//testNums = append(testNums, big.NewInt(13))
	//testNums = append(testNums, big.NewInt(29))

	for i := 0; i < len(testNums); i++ {

		if isPrime(testNums[i]) {
			fmt.Println("  ", testNums[i], " IS A PRIME Number!")
		} else {
			fmt.Println("  ", testNums[i], " is NOT a Prime Number!")
		}
	}

}

func isPrime(n *big.Int) bool {

	temp := big.NewInt(0)
	increment := big.NewInt(1)
	zero := big.NewInt(0)

	for i := big.NewInt(2); i.Cmp(n) < 0; i.Add(i, increment) {

		r := big.NewInt(0)
		mod := big.NewInt(0)

		_, mod = temp.QuoRem(n, i, r)

		if mod.Cmp(zero) == 0 {
			return false
		}
	}

	return true
}
