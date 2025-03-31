package main

import (
	"fmt"
	"math/big"
)

func main()  {
	fmt.Println("**********************")

	numStr := "0.28579684557497233287"

	bf, isOk := big.NewFloat(0.0).SetString(numStr)

	if !isOk {
		panic("Big Float SetString failed to parse numStr")
	}

	s:= bf.Text('f', 20)

	fmt.Println("numStr: ", numStr)
	fmt.Println("output: ", s)

}
