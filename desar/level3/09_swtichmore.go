package main

import "fmt"

func main() {
	favSport := "surfing"

	switch favSport {
	case "skiing":
		fmt.Println("Go to the alps")
	case "swimming":
		fmt.Println("Go to the pool")
	}
}
