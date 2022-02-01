package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a magnitude of the earthquake.")
		return
	}
	richter, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println("I could not get that, sorry")
		return
	}
	var desc string
	switch r := richter; {
	case r < 2:
		desc = "micro"
	case r >= 2 && r < 3:
		desc = "very minor"
	case r >= 3 && r < 4:
		desc = "minor"
	case r >= 4 && r < 10:
		desc = "great"
	default:
		desc = "massive"

	}
	fmt.Printf("%.2f is %s\n", richter, desc)
}
