package main

import "fmt"

var x bool

var y string = "cowboy"

func main()  {
	a := 10
	b := "golang"
	c := 4.17
	d := true

	x = false

	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, d)
	fmt.Printf("%v \t %T \n", x, x)
	fmt.Printf("%v \t %T \n", y, y)
}