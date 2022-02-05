package main

import "fmt"

func main() {
	items := []string{
		"hello", "dog", "cat", "cow", "fish",
		"bird", "leo", "moon",
	}

	top3 := items[:3]
	fmt.Print("Items top3:\n ", top3)
	fmt.Println()

}
