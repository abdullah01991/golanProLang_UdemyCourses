package main

import "fmt"

/*
Following table shows all the arithmetic operators supported by Go language.
Assume variable A holds 10 and variable B holds 20
*/
// global variable x1 & y1

var x1 = 100
var y1 = 25

func main() {
	// Local variable A & B
	A := 10
	B := 20
	fmt.Println(A - B)   // -10
	fmt.Println(B - A)   // 10
	fmt.Println(x1 - y1) //75
	fmt.Println(y1 - x1) // -75
}
