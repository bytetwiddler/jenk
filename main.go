package main

import (
	"fmt"

	add "github.com/bytetwiddler/bbmath/add"
)

func main() {
	fmt.Println("bbmath")

	a, err := add.Add(1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

}
