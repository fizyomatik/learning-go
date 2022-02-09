package main

import (
	"bytes"
	"fmt"
)

func main() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, png...)

	fmt.Print(bytes.Equal(png, header))

}
