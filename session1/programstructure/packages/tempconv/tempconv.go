package tempconv

import (
	"fmt"
	"log"
)

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

var logger *log.Logger
func init() {
	logger = log.Default()
	logger.SetPrefix("tempconv: ")
	logger.Print("Package initialized")
}

func CToF(c Celsius) Fahrenheit {
	logInfo(fmt.Sprintf("CToF(%g) called", c))
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	logInfo(fmt.Sprintf("FToC(%g) called", f))
	return Celsius((f - 32) * 5 / 9)
}

func logInfo(s string) {
	logger.Print(s)
}