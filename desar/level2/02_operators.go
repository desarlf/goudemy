package main

import "fmt"

func main() {

	x := 42
	y := 24
	z := 32
	a := 33
	b := 43
	c := 42


	d := x == y
	e := y <= z
	f := z >= a
	g := a != b
	h := b < c
	i := c > x

	fmt.Println(d, e, f, g, h, i)

}