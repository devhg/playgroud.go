package algorithm

import "strings"

//13. 罗马数字转整数

func romanToInt(s string) int {
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	split := strings.Split(s, "")
	var sum int = 0
	if len(split) > 1 {
		for i := len(split) - 1; i > 0; {
			r, l := m[split[i]], m[split[i-1]]
			if r > l {
				sum += r - l
				i -= 2
			} else {
				sum += r
				i--

			}
			if i == 0 {
				sum += m[split[i]]
				break
			}
		}
	} else {
		sum = m[split[0]]
	}
	return sum
}
