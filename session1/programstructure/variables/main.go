package main

import "fmt"

// Pi is visible not only in this package, but also through the entire program by other packages
const Pi = 3.14

func main() {
    VariableDeclarations()
    VariableAssignments()
}

func VariableDeclarations() {
    var a int = 10
    
    // short declaration
    // preferred in local declartions because of brevity 
    b := 20

    // initialise multiple variables in single declaration
    var c, d int = 30, 40

    // multiple short declaration
    e, f := 50, 60
    // short declaration doesn't declare if already declared and just reassigns
    e, g := 80, 90

    // zero value
    var h int

    // built in method
    fmt.Printf("a = %d, b = %d, c = %d, d = %d, e = %d, f = %d, g = %d, h = %d\n", a, b, c, d, e, f, g, h)
}

func VariableAssignments() {
    var x, y int

    // assignment
    // assignment statment <variable> <operator> <expression>
    x = 10
    y = 20

    // tuple assignment
    // expressions in tuple must be same type
    // expressions are evaluated before variables are updated
    x, y = y, x

    // we can assign unwanted values to blank identifiers
    // _, ok = io.Copy(dst, src)
}
