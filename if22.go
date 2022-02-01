package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month")
		return
	}

	m := os.Args[1]
	m = strings.ToLower(os.Args[1])
	if m == "june" || m == "july" || m == "august" {
		fmt.Println("Summer")
	} else if m == "jan" || m == "feb" || m == "dec" {
		fmt.Println("Winter")
	} else if m == "mar" || m == "ap" || m == "may" {
		fmt.Println("Spring")
	} else if m == "sep" || m == "oct" || m == "nov" {
		fmt.Println("Autumn")
	} else {
		fmt.Printf("%q is not a month.\n", os.Args[1])
	}
}
