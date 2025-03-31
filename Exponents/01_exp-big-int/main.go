package main

/*
 */

import (
	"fmt"
	"math/big"
)

func main() {

	exp01()
}

func exp01() {

	/*
		x := big.NewInt(10)
		y := big.NewInt(-3)
		m := big.NewInt(0)
		Output:
		x=  10
		y=  -3
		m=  0
		z=  1
	*/

	/*
		x := big.NewInt(10)
		y := big.NewInt(3)
		m := nil
		x=  10
		y=  3
		m=  0
		z=  1000

	*/

	x := big.NewInt(10)
	y := big.NewInt(3)
	m := big.NewInt(0)

	z := big.NewInt(0).Exp(x, y, nil)

	fmt.Println("x= ", x)
	fmt.Println("y= ", y)
	fmt.Println("m= ", m)
	fmt.Println("z= ", z)

}
