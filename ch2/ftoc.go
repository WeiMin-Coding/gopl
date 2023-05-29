package main

import (
	"fmt"
)

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

	var s string
	fmt.Println(s)

	//var i, j, k int
	//var b, f, s = true, 2.3, "four"
	//var f, err = os.Open("123")
	//freq := rand.Float64() * 3.0
	//t := 0.0
	//i := 100
	//var boiling float64 = 100
	//var names []string
	//var err error
	//var p Point
	//i, j := 0, 1
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
