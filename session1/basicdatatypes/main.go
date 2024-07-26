package main

import (
	"fmt"
	"math"
)

// Types
// int8, int16, int32, int64
// uint8, uint16, uint32, uint64, uintptr
// byte
// rune
// int
// uint
// float32, float64
// complex64, complex128
// bool
// string

// utf8
// constants

const e = 2.718281828
const Avagadro = 6.02214129e23

func main() {
	ints()
    floats()   
    // complex numbers
    // bools - true and false , operators ! (Not), && (And), || (Or)
    strings()
    // TODO below
    // utf8
    // bytes
    // constants
}

func ints() {
	// signed int - -2^n-1 to 2^n-1
	// unsigned int 0 to 2^n-1

	var x int8 = 127
	var y uint8 = 255

	// unsigned, atleast 32 bits, not an alias for int32, depends on processor architecture
	// for amdd64 its 64 bits
	var a int = 0

    z := int(x) + int(y)

	fmt.Printf("x = %d, y = %d, a = %d, z = %d \n", x, y, a, z)
}

func floats() {
    // float32, float64

    var x float32 = 3.141592653
    var y float64 = 3.141592653589793

    maxFloat32 := math.MaxFloat32
    minFloat32 := math.SmallestNonzeroFloat32

    fmt.Printf("x = %.3f, y = %.3f, maxFloat32 = %.3f, minFloat32 = %.3f \n", x, y, maxFloat32, minFloat32)
}

func strings() {
    // immutable sequence of bytes
    // intrpreted as UTF-8 encoded sequences of unicode code points

    s := "hello world"
    fmt.Println(len(s))
    fmt.Println(s[0], s[7])

    s = "hello" + " " + "world" + "!"
    fmt.Println(s)
}