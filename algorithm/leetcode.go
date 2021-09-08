package algorithm

import "fmt"

func Main() {
	var num = []int{1, 2, 3, 4}
	var resp []int
	resp = append(resp, num[0])

	length := len(num)
	for i := 1; i < length; i++ {
		resp = append(resp, resp[i-1]+num[i])
	}

	for i, i2 := range resp {
		fmt.Println(i, i2)
	}
}
