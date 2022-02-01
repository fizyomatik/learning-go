package main

import "fmt"

func main() {
	hiz := 50
	force := 3.5

	hiz = int(float64(hiz) * force)
	fmt.Println(hiz)
}
