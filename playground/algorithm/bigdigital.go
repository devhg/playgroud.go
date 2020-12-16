package algorithm

import (
	"fmt"
	"strconv"
)

func printNumbers(n int) []string {
	base := make([]string, 10)
	for i := 0; i < 10; i++ {
		base[i] = strconv.Itoa(i)
	}

	for i := 1; i <= n-1; i++ {
		temp := make([]string, len(base))
		copy(temp, base)
		digit := i + 1
		for j := 1; j < 10; j++ {
			for _, s := range temp {
				zero := ""
				for x := 1; x < digit-len(s); x++ {
					zero += "0"
				}
				item := strconv.Itoa(j) + zero + s
				base = append(base, item)
			}
		}
	}
	fmt.Println(base)
	return base
}
