package main

import "fmt"

func main() {
	var max uint64 = 4000000
	var sn1 uint64 = 1
	var sn2 uint64 = 2
	var sum uint64 = sn2

	for {
			fb := sn1 + sn2
			sn1 = sn2
			sn2 = fb
			if fb > max {
				break
			}
			if fb % 2 == 0 {
				sum += fb
			}
	}

	fmt.Println(sum)
}
