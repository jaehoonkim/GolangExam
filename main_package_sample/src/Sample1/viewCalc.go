package main 

import (
	"fmt"
	"calc"
)

func helloGopher() {

	fmt.Println("Hello Gopher~!")

}

func addNumber(x, y int) {

	fmt.Println(calc.Add(x, y))

}