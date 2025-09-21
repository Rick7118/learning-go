package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v > lim {
		return v
	}
	return lim
}

func ifloops3() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(2, 0.5, 1.5),
	)
}
