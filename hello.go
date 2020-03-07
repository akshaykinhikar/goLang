package main

import (
	"fmt"
	"time"
)

func main() {
	// print stmt
	fmt.Println("Hello, world.")

	// variable declaration, initialization
	var a int32

	a = 15
	fmt.Println("value of int a =", a)

	// dataType
	// boolean dataType
	var b bool

	b = true

	fmt.Println("value of bool b = ", b)

	// dataType
	// float
	var c float32

	c = 15.01

	fmt.Println("value of float c = ", c)

	// dataType - string

	var d string
	d = "lorem"
	fmt.Println("value of string d = ", d)

	// datatype - Integer
	var e int
	e = 123
	fmt.Println("value of int e = ", e)

	// dataType - const
	const f = 1234
	fmt.Println("value of const f = ", f)

	// for loop
	i := 10
	for i > 0 {
		fmt.Println("Value of i is ", i)
		time.Sleep(time.Second)
		i = i - 1
	}

	for i := 10; i > 0; i-- {
		fmt.Println("Value of i is ", i)
	}

}
