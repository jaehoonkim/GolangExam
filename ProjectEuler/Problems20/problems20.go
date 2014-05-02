package Problems20

import (
	"fmt"
	"math/big"
	"strings"
	"strconv"
)

func factorial(n *big.Int) (result *big.Int) {
	if n.Cmp(big.NewInt(0)) == 0  {
		result = big.NewInt(1)
	} else {
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, big.NewInt(1))))
	}
	return
}

func Execute() {
	total := factorial(big.NewInt(100))
	//fmt.Println(total)
	str := total.String()
	count := strings.Count(str, "")
	
	var sum int64
	runes := []rune(str)
	for i := 0; i < count-1; i++ {
		s := string(runes[i])
		unit, _ := strconv.ParseInt(s, 10, 64)
		sum += unit
	}	
	fmt.Println(sum)
}