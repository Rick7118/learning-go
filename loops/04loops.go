package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprintf("%fi", math.Sqrt(-x))
	}
	return fmt.Sprintf("%f", math.Sqrt(x))
}

func ifloops2() {
	ans := sqrt(-89.5)
	fmt.Println(ans)
}
