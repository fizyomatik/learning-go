package main

import (
	"fmt"
	"sort"
)

func main() {

	names := [...]string{
		"Einstein", "Shepard",
		"Tesla",
	}

	var books = []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	sort.Ints(nums[:])

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
