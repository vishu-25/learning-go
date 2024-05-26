package main

import "fmt"

func main() {
	a := 2
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "gophers"

	fmt.Println(b, c, d, f)

	// The below code is an example of zero value in Go
	/*
	 var g int
	 fmt.Println(g)
	*/
}
