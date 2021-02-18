package main

import "fmt"

func main() {
	defer func() { fmt.Println("defer...0") }()

	fmt.Println("exec...0")

	defer func() { fmt.Println("defer...1") }()

	fmt.Println("exec...1")

	if true {
		return
	}

	defer func() { fmt.Println("defer...2") }()
}
