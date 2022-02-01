package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)

func main() {
	rates := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Give me a number")
		return
	}

	fmt.Printf("%d USD is %g EUR\n", n, float64(n)*rates[EUR])
	fmt.Printf("%d USD is %g GBY\n", n, float64(n)*rates[GBP])
	fmt.Printf("%d USD is %g JPY\n", n, float64(n)*rates[JPY])

}
