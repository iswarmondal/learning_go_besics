package main

import "fmt"

func main() {
	var a int = 1
	b := "great"

	x := &a

	type fahernite int
	type celsicus int

	var f fahernite = 0
	var c celsicus = 32

	fmt.Println(a, b, *x, f, c)

	c = celsicus((f - 32) * 5 / 9)
	fmt.Println(c)

	{
		var z = 0
		fmt.Println(z)
	}
}
