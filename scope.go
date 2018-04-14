package main

import "fmt"

var x = "OUTER"

func second() {
	fmt.Println("x is", x)
}

func first() {
	// var x = "INNER"
	second()
}

func main() {
	first()
}
