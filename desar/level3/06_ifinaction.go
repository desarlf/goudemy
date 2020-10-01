package main

import "fmt"

func main() {

	x := 10
	y := 20

	if x < y {
		fmt.Println("X is more than Y")

	} else if y < x {
		fmt.Println("Y is bigger than X")
	} else {
		fmt.Println("X is as big as Y")
	}

}