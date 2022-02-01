package main

import "fmt"

func main() {
	var sum int

	for i := 0; i <= 10; i++ {
		sum += i + 1
		fmt.Print(i)
		if i < 10 {
			fmt.Print("+")
		}
	}
	fmt.Printf("=%d\n", sum)
}
