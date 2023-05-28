package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	var test string
	fmt.Println(test)

	var x int = 1
	p := &x
	fmt.Println(p)
	*p = 2
	fmt.Println(&x)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
