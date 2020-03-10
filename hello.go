package main

import (
	"fmt"
)

func greaterNo(i int, j int) int {
	if i < j {
		fmt.Println("greater no is ", j)
		return j
	} else {
		fmt.Println("greater no is ", i)
		return i
	}
}

func addByValue(a int, b int, c int) int {
	c = a + b
	// fmt.Println(c)
	return c
}

func addByReference(a int, b int, c *int) int {
	*c = a + b
	// fmt.Println(*c)
	return *c
}

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
		// time.Sleep(time.Second)
		i = i - 1
	}

	for i := 10; i > 0; i-- {
		fmt.Println("Value of i is ", i)
	}

	// if statements
	j := 10

	if j > 0 {
		fmt.Println("value of i is greater than 0")
	}

	for i := 0; i <= 10; i++ {
		switch i {
		case 10:
			fmt.Println("got 10")
		}

	}

	// function

	fmt.Println("function will print greater value", greaterNo(1, 2))

	// byValue vs byReference

	// normal flow

	// byValue(5, 2, 3)
	var c1 int
	// fmt.Println("", addByValue(5, 2, c1))
	fmt.Println("effect of using byValue on c1 --> ", c1)

	// add by reference

	// fmt.Println("", addByReference(5, 2, &c1))
	fmt.Println("effect of using by reference on c1--> ", c1)

	// Array

	arr := [5]int{0, 21, 33, 22, 11}

	fmt.Println(arr)

}
