package test

import "fmt"

func TestCalc() {
	base := 5000.0
	rate := 0.05
	for loop := 0; loop < 5; loop++ {
		base = base * (1.0 + rate)
		fmt.Println("loop", loop, "=", base)
	}
}
