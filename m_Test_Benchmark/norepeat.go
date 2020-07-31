package m_Test_Benchmark

//最长无重复字符串 中文实现
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
