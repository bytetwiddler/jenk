package main

import (
	"fmt"

	j "github.com/bytetwiddler/jenk/add"
)

func main() {
	fmt.Println("jenk")

	a, err := j.Add(1, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

}
