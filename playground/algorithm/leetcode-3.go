package algorithm

import "strings"

//3. 无重复字符的最长子串(leetcode)

//滑动窗口，效果更好
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

// 借助map实现
func map_lengthOfLongestSubstring(s string) int {
	var lastOccurred = make(map[byte]int)

	maxLength := 0
	startP := 0
	for i, ch := range []byte(s) {
		lastPos, ok := lastOccurred[ch]
		if ok && lastPos >= startP {
			startP = lastPos + 1
		}
		if maxLength < i-startP+1 {
			maxLength = i - startP + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

//中文实现
func cn_lengthOfLongestSubstring(s string) int {
	var lastOccurred = make(map[rune]int)

	maxLength := 0
	startP := 0
	for i, ch := range []rune(s) {
		lastPos, ok := lastOccurred[ch]
		if ok && lastPos >= startP {
			startP = lastPos + 1
		}
		if maxLength < i-startP+1 {
			maxLength = i - startP + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
