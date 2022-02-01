package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%5s", "X")

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Please only one argument")
		return
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Give me a number")
		return
	}

	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

	}
	fmt.Println()
	for i := 0; i <= n; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= n; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()

	}
	fmt.Println()
}
