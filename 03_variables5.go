package main

import "fmt"

func variables5() {
	var f int
	var g float64
	var h bool
	var l string

	fmt.Printf("%v %v %v %q \n", f, g, h, l)
}

// Zero values
// Variables declared without an explicit initial value are given their zero value.

// The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
