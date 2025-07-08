package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// creates a complex number 2.5 + 3.1i
	x := complex(2.5, 3.1)
	// creates a complex number 10.2 + 2.0i
	y := complex(10.2, 2)
	// basic arithmetic operations on complex numbers
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	// use the built-in real function to extract the real portion of a complex number
	fmt.Println(real(x))
	// use the build in imag function to extract the imaginary portion of a complex number
	fmt.Println(imag(x))
	// use the Abs function in the math/cmplx package to take the absolute value of a complex number
	fmt.Println(cmplx.Abs(x))
}
