package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	BookTitles := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Tell me a book title")
		return
	}
	query := strings.ToLower(os.Args[0])

	fmt.Println("Search Results:")

	var found bool
	for _, v := range BookTitles {
		if strings.Contains(strings.ToLower(v), query) {
			fmt.Println("+", v)
			found = true
		}
	}
	if !found {
		fmt.Printf("We do'nt have the book: %q\n", query)
	}

}
