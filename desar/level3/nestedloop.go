package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d, %d \n", i, j)
		}
	}

}