package main

import (
	"fmt"
	"math/big"
)

/* 600851475143
We can use the Fundamental Theorem of Arithmetic which states:
Any integer greater than 1 is either a prime number, or can be written as a unique product of prime numbers (ignoring the order).
*/

func main() {
	n := big.NewInt(600851475143)
	largestFact := findLarestPrimeFactor(n)
	fmt.Println("The largest Prime Factor of", n)
	fmt.Println("is", largestFact)
}

func findLarestPrimeFactor(n *big.Int) *big.Int {

	newNum := big.NewInt(0).Set(n)
	limit := big.NewInt(1)
	mod := big.NewInt(1)
	zero := big.NewInt(0)
	largestFac := big.NewInt(0)
	increment := big.NewInt(1)

	for counter := big.NewInt(2); limit.Mul(counter, counter).Cmp(newNum) <= 0; {

		if mod.Mod(newNum, counter).Cmp(zero) == 0 {
			newNum = newNum.Div(newNum, counter)
			largestFac.Set(counter)
		} else {
			counter = counter.Add(counter, increment)
		}

	}

	if newNum.Cmp(largestFac) > 0 { // the remainder is a prime number
		largestFac.Set(newNum)
	}

	return largestFac
}

/*
http://www.mathblog.dk/project-euler-problem-3/
const long numm = 600851475143;
long newnumm = numm;
long largestFact = 0;

int counter = 2;
while (counter * counter <= newnumm) {
	if (newnumm % counter == 0) {
			newnumm = newnumm / counter;
			largestFact = counter;
		} else {
			counter++;
		}
}

if (newnumm > largestFact) { // the remainder is a prime number
	largestFact = newnumm;
}
*/
