package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 3 || len(os.Args) <= 0 {
		fmt.Println("Please give a some arguments!")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || min > max {
		fmt.Println("Give numbers")
		return

	}
	var sum int

	for i := 0; i <= max; {
		i++
		if i%2 != 0 {
			continue
		}

		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print("+")
			i++

		}

	}
	fmt.Printf("=%d", sum)
	fmt.Println()
}
