package Problems6

import (
	"fmt"
)

func Execute() {
	var sum, sum2 int
	for i := 0; i < 101; i++ {
		sum += i
		sum2 += (i * i)
	}
	fmt.Println((sum * sum) - sum2)
}