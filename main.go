package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abcc", "dog cat d a"))
}
func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(pattern) != len(split) {
		return false
	}

	mp := make(map[byte]string)
	for i := 0; i < len(pattern); i++ {
		v, ok := mp[pattern[i]]
		if ok {
			if v != split[i] {
				return false
			}
		} else {
			for _, v1 := range mp {
				if v1 == split[i] {
					return false
				}
			}
			mp[pattern[i]] = split[i]
		}
	}
	return true
}
