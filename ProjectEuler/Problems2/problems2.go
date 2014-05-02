package Problems2

import (
	"fmt"
)

func Execute() {
	var s1, s2, s3, sum int
	s1, s2, s3 = 1, 2, 0
	sum = s2
	for ; s3 < 4000000; {
		s3 = s1 + s2
		if (s3 % 2) == 0 {
			sum += s3
		}
		s1 = s2
		s2 = s3
	}
	
	fmt.Println(sum)
}

