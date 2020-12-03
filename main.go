package main

import "fmt"

func main() {
	str := "ecbacba"
	number := removeDuplicateLetters(str)
	fmt.Println(number)
}

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	exist := [26]bool{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	ret := make([]byte, 0)
	for _, ch := range s {
		if !exist[ch-'a'] {
			// 如果当前元素在栈中不存在，且他小于栈顶元素的字典序，且后面还有该元素
			for len(ret) > 0 && uint8(ch) < ret[len(ret)-1] && count[ret[len(ret)-1]-'a'] > 0 {
				exist[ret[len(ret)-1]-'a'] = false
				ret = ret[:len(ret)-1] // 出栈
			}
			ret = append(ret, uint8(ch))
			exist[ret[len(ret)-1]-'a'] = true
		}
		// 栈中已经存在，存量--
		count[ch-'a']--
	}
	return string(ret)
}

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	ans := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			ans++
			for j := i * 2; j < n; j += i {
				isPrime[j] = true
			}
		}
	}
	return ans
}
