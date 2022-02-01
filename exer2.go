package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the size of table")
		return
	}
	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return

	}
	// for the header
	fmt.Printf("%5s", "X")

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// for the vertical header
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		// for the cells

		for j := 0; j <= size; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

}
