package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me a name of month")
		return

	}
	switch m := os.Args[1]; m {
	case "jan", "feb", "dec":
		fmt.Println("Winter")
	case "ap", "may", "mar":
		fmt.Println("Spring")
	case "jun", "jul", "aug":
		fmt.Println("summer")
	case "sep", "oct", "nov":
		fmt.Println("autumn")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}
