package main

import (
	"fmt"
	"gophers-dojo/session1/programstructure/packages/tempconv"
)

func main() {
	fmt.Printf("Celsius(100): %g\n", tempconv.CToF(tempconv.BoilingC))
}