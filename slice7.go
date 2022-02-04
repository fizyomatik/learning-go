package main

import "fmt"

func main() {
	emo := []string{"emo", "bugun", "burda"}

	ebay := []string{"ebaya"}

	emo = append(emo, ebay...)

	fmt.Println(emo)
	if len(emo) == 4 {
		fmt.Println("moin")
	}
}
