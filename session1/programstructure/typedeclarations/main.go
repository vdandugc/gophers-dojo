package main

import "fmt"

// defines new named type that has the same underlying type as an existing type
// provides a way to separate different and perhaps incompatible uses of the underlying type
// so that they can't be mixed unintentionally
type Celsius float64
type Fahrenheit float64

// func (c Celsius) String() string {
// 	return fmt.Sprintf("%gC", c)
// }

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	RoomTemperatureF Fahrenheit = 75
)

func main() {
		fmt.Printf("%g\n", CToF(BoilingC))
		fmt.Printf("%g\n", FToC(RoomTemperatureF))
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}