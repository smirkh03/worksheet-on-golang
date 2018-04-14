package main

import "fmt"

/* 
  Example for scope checking, adapted from the 
  Programming Language Explorations (PLE) book
*/

var x = "OUTER"

func second() {
	fmt.Println("x is", x)
}

func first() {
	// var x = "INNER"
	// golang doesn't allow unusued variables
	// so we can't event declare this variable incase dynamic scoping
	// we allowed (and it is not!)
	second()
}

func main() {
	first()
}
