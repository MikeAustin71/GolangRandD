package main

import (
	"fmt"
	"math/big"
)

/* Problem -3 Largest Prime Factor
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

*/

/* 600851475143
We can use the Fundamental Theorem of Arithmetic which states:
Any integer greater than 1 is either a prime number, or can be written as a unique product of prime numbers (ignoring the order).

This solution finds the largest Prime Factor of the subject number.

*/

func main() {
	n := big.NewInt(600851475143)
	largestFact := findLargestPrimeFactor(n)
	fmt.Println("---------------------------------------------------")
	fmt.Println("The largest Prime Factor of", n, " is", largestFact)
	fmt.Println("---------------------------------------------------")
}

func findLargestPrimeFactor(n *big.Int) *big.Int {

	newNum := big.NewInt(0).Set(n)
	limit := big.NewInt(1)
	mod := big.NewInt(1)
	zero := big.NewInt(0)
	largestFac := big.NewInt(0)
	increment := big.NewInt(1)
	value2 := big.NewInt(2)

	for counter := big.NewInt(2); limit.Mul(counter, counter).Cmp(newNum) <= 0; {

		if mod.Mod(newNum, counter).Cmp(zero) == 0 {
			newNum = newNum.Div(newNum, counter)
			largestFac.Set(counter)
		} else {

			if counter.Cmp(value2) == 0 {
				counter = counter.Add(counter, increment)
				increment.Set(value2)
				continue
			}
			counter = counter.Add(counter, increment)
		}
	}

	if newNum.Cmp(largestFac) > 0 { // the remainder is a prime number
		largestFac.Set(newNum)
	}

	return largestFac
}

// The largest Prime Factor of 600851475143  is 6857

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
