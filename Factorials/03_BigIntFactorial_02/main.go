package main

import (
	"fmt"
	"math/big"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	n := big.NewInt(14)

	c := incrementor(n)

	cSum := puller(c)

	fmt.Println(n, "Factorial =", <-cSum)

}

func incrementor(n *big.Int) chan *big.Int {
	out := make(chan *big.Int)
	zero := big.NewInt(0)
	one := big.NewInt(1)

	go func() {
		for i := big.NewInt(0).Set(n); i.Cmp(zero) > 0; i.Sub(i, one) {
			out <- big.NewInt(0).Set(i)
		}
		close(out)

	}()

	return out
}

func puller(c chan *big.Int) chan *big.Int {
	out := make(chan *big.Int)

	go func() {
		sum := big.NewInt(1)
		for n := range c {
			sum = sum.Mul(sum, n)
		}
		out <- sum
		close(out)
	}()

	return out
}
