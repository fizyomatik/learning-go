package main

import "fmt"

func main() {
	items := []string{
		"hello", "dog", "cat", "cow",
		"fish", "bird", "leo", "moon",
		"me", "you", "jo", "water",
		"me",
	}

	const pageSize = 4
	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}

		currentPage := items[from:to]
		head := fmt.Sprintf("Page #%d\n", (from/pageSize)+1)
		fmt.Println()
		fmt.Print(head, currentPage)
	}
	fmt.Println()

}
