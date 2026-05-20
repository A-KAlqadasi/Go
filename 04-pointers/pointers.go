package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println("Address of x:", p)
	fmt.Println("Value of x:", *p)
}